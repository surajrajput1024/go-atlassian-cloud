package client

import (
	"context"
	"net/http"
)

// RESTDoer is the minimal interface needed by product clients (e.g. Jira) to perform REST calls.
// *Client implements RESTDoer. Tests can provide a fake implementation.
type RESTDoer interface {
	GetWithContext(ctx context.Context, path string) (*http.Response, error)
	DoWithContext(ctx context.Context, req *http.Request) (*http.Response, error)
	RestAPIURL() string
	DoJSON(ctx context.Context, method, path string, body interface{}, out interface{}) error
}
