package client

import (
	"net/http"
	"time"
)

type Options struct {
	Timeout         time.Duration
	MaxRetries      int
	RetryBackoffMin time.Duration
	RetryBackoffMax time.Duration
	Transport       http.RoundTripper
}

func DefaultOptions() Options {
	return Options{
		Timeout:         DefaultHTTPTimeout,
		MaxRetries:      DefaultMaxRetries,
		RetryBackoffMin: DefaultRetryBackoffMin,
		RetryBackoffMax: DefaultRetryBackoffMax,
	}
}
