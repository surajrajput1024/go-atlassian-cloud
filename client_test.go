package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient_InvalidConfig(t *testing.T) {
	_, err := NewClient(&Config{}, DefaultOptions())
	if err == nil {
		t.Error("expected error for invalid config")
	}
}

func TestNewClient_Ok(t *testing.T) {
	cfg := &Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, err := NewClient(cfg, DefaultOptions())
	if err != nil {
		t.Fatal(err)
	}
	if cl == nil {
		t.Error("client is nil")
	}
}

func TestClient_Do_SetsAuth(t *testing.T) {
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	cfg := &Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "secret"}
	cl, _ := NewClient(cfg, Options{MaxRetries: 0})
	req, _ := http.NewRequest(http.MethodGet, srv.URL, nil)
	resp, err := cl.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if gotAuth == "" || len(gotAuth) < 6 || gotAuth[:6] != "Basic " {
		t.Errorf("expected Basic auth header, got %q", gotAuth)
	}
}
