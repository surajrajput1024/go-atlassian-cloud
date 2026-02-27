package jira

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
)

func TestNew(t *testing.T) {
	cl, err := atlassian.NewClient(testConfig(), atlassian.DefaultOptions())
	if err != nil {
		t.Fatal(err)
	}
	j := New(cl)
	if j == nil || j.do != cl {
		t.Error("New returned nil or wrong client")
	}
}

func TestClientPath(t *testing.T) {
	cl, _ := atlassian.NewClient(testConfig(), atlassian.DefaultOptions())
	j := New(cl)
	p := j.path("myself")
	if p == "" || p != cl.RestAPIURL()+"/myself" {
		t.Errorf(msgPathFormat, p)
	}
}

func TestClientGetJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"accountId": "x", "displayName": "y"})
	}))
	defer srv.Close()
	j := testJiraClient(t, atlassian.Options{MaxRetries: 0})
	var out struct {
		AccountID   string `json:"accountId"`
		DisplayName string `json:"displayName"`
	}
	if err := j.doJSON(context.Background(), "GET", srv.URL, nil, &out); err != nil {
		t.Fatal(err)
	}
	if out.AccountID != "x" || out.DisplayName != "y" {
		failOut(t, out)
	}
}
