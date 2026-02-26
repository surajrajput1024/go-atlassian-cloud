package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/surajsinghrajput/go-atlassian-cloud/client/auth"
	httputil "github.com/surajsinghrajput/go-atlassian-cloud/client/http"
	"github.com/surajsinghrajput/go-atlassian-cloud/client/retry"
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
	if opts.Timeout <= 0 {
		opts.Timeout = DefaultHTTPTimeout
	}
	if opts.MaxRetries < 0 {
		opts.MaxRetries = 0
	}
	hc := &http.Client{Timeout: opts.Timeout}
	if opts.Transport != nil {
		hc.Transport = opts.Transport
	}
	c := &Client{
		cfg:    cfg,
		opts:   opts,
		client: hc,
	}
	return c, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
		req.Body.Close()
	}
	auth.SetBasicAuth(req, c.cfg.Email, c.cfg.APIToken)
	var lastErr error
	for attempt := 0; attempt <= c.opts.MaxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(retry.Backoff(attempt, c.opts.RetryBackoffMin, c.opts.RetryBackoffMax))
			if len(bodyBytes) > 0 {
				req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}
		} else if len(bodyBytes) > 0 {
			req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
		resp, err := c.client.Do(req)
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
	u, err := httputil.ParseURL(c.cfg.BaseURL(), path)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
