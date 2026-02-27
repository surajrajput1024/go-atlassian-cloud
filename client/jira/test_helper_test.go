package jira

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
)

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"

	testDomain   = "site.atlassian.net"
	testEmail    = "u@e.com"
	testAPIToken = "tok"

	msgOutFormat = "out = %+v"
)

func failOut(t *testing.T, out interface{}) {
	t.Helper()
	t.Errorf(msgOutFormat, out)
}

// writeJSON sets Content-Type, status code, and encodes v as JSON. For test handlers.
func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(code)
	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}

func testConfig() *atlassian.Config {
	return &atlassian.Config{Domain: testDomain, Email: testEmail, APIToken: testAPIToken}
}

func testJiraClient(t *testing.T, opts atlassian.Options) *Client {
	t.Helper()
	if opts.MaxRetries < 0 {
		opts.MaxRetries = 0
	}
	cl, err := atlassian.NewClient(testConfig(), opts)
	if err != nil {
		t.Fatal(err)
	}
	return New(cl)
}

// redirectToServerTransport sends all requests to the target server (for tests).
type redirectToServerTransport struct {
	base *url.URL
	rt   http.RoundTripper
}

func (r *redirectToServerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := req.Clone(req.Context())
	u := *req2.URL
	u.Scheme = r.base.Scheme
	u.Host = r.base.Host
	req2.URL = &u
	return r.rt.RoundTrip(req2)
}

// testJiraClientWithServer returns a Jira client that redirects requests to srv. Use for testing API methods that build URLs from config.
func testJiraClientWithServer(t *testing.T, srv *httptest.Server) *Client {
	t.Helper()
	base, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	tr := &redirectToServerTransport{base: base, rt: http.DefaultTransport}
	return testJiraClient(t, atlassian.Options{MaxRetries: 0, Transport: tr})
}
