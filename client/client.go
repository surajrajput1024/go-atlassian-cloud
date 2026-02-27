package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/surajrajput1024/go-atlassian-cloud/client/auth"
	httputil "github.com/surajrajput1024/go-atlassian-cloud/client/http"
	"github.com/surajrajput1024/go-atlassian-cloud/client/retry"
)

type Client struct {
	cfg    *Config
	opts   Options
	client *http.Client
}

func NewClient(cfg *Config, opts Options) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	o := opts
	if o.Timeout <= 0 {
		o.Timeout = DefaultHTTPTimeout
	}
	if o.MaxRetries < 0 {
		o.MaxRetries = 0
	}
	hc := &http.Client{Timeout: o.Timeout}
	if o.Transport != nil {
		hc.Transport = o.Transport
	}
	return &Client{cfg: cfg, opts: o, client: hc}, nil
}

func NewClientWithOptions(cfg *Config, optFuncs ...Option) (*Client, error) {
	return NewClient(cfg, applyOptions(optFuncs))
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.DoWithContext(req.Context(), req)
}

func (c *Client) DoWithContext(ctx context.Context, req *http.Request) (*http.Response, error) {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body.Close()
	}
	var lastErr error
	for attempt := 0; attempt <= c.opts.MaxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(retry.Backoff(attempt, c.opts.RetryBackoffMin, c.opts.RetryBackoffMax))
		}
		r := req.Clone(ctx)
		if len(bodyBytes) > 0 {
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
		auth.SetBasicAuth(r, c.cfg.Email, c.cfg.APIToken)
		resp, err := c.client.Do(r)
		if err != nil {
			lastErr = err
			continue
		}
		if !retry.IsRetryableStatusCode(resp.StatusCode) {
			return resp, nil
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		var apiErr *APIError
		apiErr, _ = newAPIError(resp)
		lastErr = apiErr
	}
	return nil, lastErr
}

func (c *Client) RestAPIURL() string {
	return c.cfg.RestAPIURL()
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.GetWithContext(context.Background(), path)
}

func (c *Client) GetWithContext(ctx context.Context, path string) (*http.Response, error) {
	u, err := httputil.ParseURL(c.cfg.BaseURL(), path)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c.DoWithContext(ctx, req)
}

func (c *Client) DoJSON(ctx context.Context, method, path string, body interface{}, out interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(raw)
	}
	u, err := httputil.ParseURL(c.cfg.BaseURL(), path)
	if err != nil {
		return fmt.Errorf("parse url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return err
	}
	auth.SetBasicAuth(req, c.cfg.Email, c.cfg.APIToken)
	resp, err := c.DoWithContext(ctx, req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		ae := &APIError{StatusCode: resp.StatusCode, Body: respBody}
		var parsed struct {
			ErrorMessages []string          `json:"errorMessages"`
			Errors        map[string]string `json:"errors"`
		}
		_ = json.Unmarshal(respBody, &parsed)
		ae.ErrorMessages = parsed.ErrorMessages
		ae.Errors = parsed.Errors
		return wrapAPIError(ae)
	}
	if out != nil && len(respBody) > 0 {
		return json.Unmarshal(respBody, out)
	}
	return nil
}
