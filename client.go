package client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
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
	c.setAuth(req)
	var lastErr error
	for attempt := 0; attempt <= c.opts.MaxRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(c.backoff(attempt))
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
		if !IsRetryableStatusCode(resp.StatusCode) {
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

func (c *Client) setAuth(req *http.Request) {
	raw := c.cfg.Email + ":" + c.cfg.APIToken
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(raw)))
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
}

func (c *Client) backoff(attempt int) time.Duration {
	min := c.opts.RetryBackoffMin
	max := c.opts.RetryBackoffMax
	if min <= 0 {
		min = DefaultRetryBackoffMin
	}
	if max <= 0 {
		max = DefaultRetryBackoffMax
	}
	d := min * time.Duration(1<<uint(attempt))
	if d > max {
		return max
	}
	return d
}

func (c *Client) RestAPIURL() string {
	return c.cfg.RestAPIURL()
}

func (c *Client) Get(path string) (*http.Response, error) {
	u, err := parseURL(c.cfg.BaseURL(), path)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
