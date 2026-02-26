package jira

import (
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/constants"
)

func TestGetCurrentUser(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"accountId":"acc1","displayName":"Test User","emailAddress":"u@e.com"}`))
	}))
	defer srv.Close()
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.Options{MaxRetries: 0})
	j := New(cl)
	_ = constants.JiraPathMyself
	path := j.path(constants.JiraPathMyself)
	if path == "" {
		t.Fatal("path empty")
	}
	var out struct {
		AccountID    string `json:"accountId"`
		DisplayName  string `json:"displayName"`
		EmailAddress string `json:"emailAddress"`
	}
	if err := j.getJSON(srv.URL, &out); err != nil {
		t.Fatal(err)
	}
	if out.AccountID != "acc1" {
		t.Errorf("accountId = %q", out.AccountID)
	}
}
