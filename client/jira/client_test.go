package jira

import (
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajsinghrajput/go-atlassian-cloud/client"
)

func TestNew(t *testing.T) {
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, err := atlassian.NewClient(cfg, atlassian.DefaultOptions())
	if err != nil {
		t.Fatal(err)
	}
	j := New(cl)
	if j == nil || j.c != cl {
		t.Error("New returned nil or wrong client")
	}
}

func TestClient_path(t *testing.T) {
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.DefaultOptions())
	j := New(cl)
	p := j.path("myself")
	if p == "" || p != cl.RestAPIURL()+"/myself" {
		t.Errorf("path(myself) = %q", p)
	}
}

func TestClient_getJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"accountId":"x","displayName":"y"}`))
	}))
	defer srv.Close()
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.Options{MaxRetries: 0})
	j := New(cl)
	var out struct {
		AccountID   string `json:"accountId"`
		DisplayName string `json:"displayName"`
	}
	if err := j.getJSON(srv.URL, &out); err != nil {
		t.Fatal(err)
	}
	if out.AccountID != "x" || out.DisplayName != "y" {
		t.Errorf("out = %+v", out)
	}
}
