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

type Option func(*Options)

func WithTimeout(d time.Duration) Option {
	return func(o *Options) { o.Timeout = d }
}

func WithTransport(rt http.RoundTripper) Option {
	return func(o *Options) { o.Transport = rt }
}

func WithRetries(n int, backoffMin, backoffMax time.Duration) Option {
	return func(o *Options) {
		o.MaxRetries = n
		o.RetryBackoffMin = backoffMin
		o.RetryBackoffMax = backoffMax
	}
}

func UseOptions(opts Options) Option {
	return func(o *Options) { *o = opts }
}

func applyOptions(opts []Option) Options {
	o := DefaultOptions()
	for _, f := range opts {
		f(&o)
	}
	return o
}
