package jira

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetIssueTypes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, []types.IssueTypeResponse{{ID: "1", Name: "Task"}})
	}))
	defer srv.Close()
	j := testJiraClient(t, atlassian.Options{MaxRetries: 0})
	var out []types.IssueTypeResponse
	if err := j.doJSON(context.Background(), "GET", srv.URL, nil, &out); err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 || out[0].Name != "Task" {
		failOut(t, out)
	}
}
