package client

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func TestResolveCloudID_UseConfigWhenSet(t *testing.T) {
	cfg := &Config{
		Domain:   "site.atlassian.net",
		Email:    "u@e.com",
		APIToken: "tok",
		CloudID:  "existing-cloud-id",
	}
	cl, err := NewClient(cfg, DefaultOptions())
	if err != nil {
		t.Fatal(err)
	}
	got, err := cl.ResolveCloudID()
	if err != nil {
		t.Fatal(err)
	}
	if got != "existing-cloud-id" {
		t.Errorf("ResolveCloudID() = %q, want existing-cloud-id", got)
	}
}

func TestResolveCloudID_CallsTenantInfo(t *testing.T) {
	wantCloudID := "test-cloud-id-123"
	transport := &mockRoundTripper{
		status: http.StatusOK,
		body:   mustJSON(types.TenantInfoResponse{CloudID: wantCloudID}),
	}
	cfg := &Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, err := NewClient(cfg, Options{MaxRetries: 0, Transport: transport})
	if err != nil {
		t.Fatal(err)
	}
	got, err := cl.ResolveCloudID()
	if err != nil {
		t.Fatal(err)
	}
	if got != wantCloudID {
		t.Errorf("ResolveCloudID() = %q, want %q", got, wantCloudID)
	}
}

type mockRoundTripper struct {
	status int
	body   []byte
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	h := make(http.Header)
	h.Set("Content-Type", "application/json")
	return &http.Response{
		StatusCode: m.status,
		Header:     h,
		Body:       &readCloser{body: m.body},
		Request:    req,
	}, nil
}

type readCloser struct {
	body []byte
	pos  int
}

func (r *readCloser) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.body) {
		return 0, io.EOF
	}
	n = copy(p, r.body[r.pos:])
	r.pos += n
	return n, nil
}

func (r *readCloser) Close() error { return nil }

func mustJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
