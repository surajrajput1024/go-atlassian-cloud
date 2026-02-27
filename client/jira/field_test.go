package jira

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, []types.FieldResponse{{ID: "summary", Name: "Summary"}})
	}))
	defer srv.Close()
	j := testJiraClient(t, atlassian.Options{MaxRetries: 0})
	var out []types.FieldResponse
	if err := j.doJSON(context.Background(), "GET", srv.URL, nil, &out); err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 || out[0].Name != "Summary" {
		failOut(t, out)
	}
}
