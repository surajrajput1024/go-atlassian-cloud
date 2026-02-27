package jira

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetProject(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockProjectResponse)
	}))
	defer srv.Close()
	j := testJiraClient(t, atlassian.Options{MaxRetries: 0})
	var out types.ProjectResponse
	if err := j.doJSON(context.Background(), "GET", srv.URL, nil, &out); err != nil {
		t.Fatal(err)
	}
	if out.Key != "PROJ" || out.Name != "Test" {
		failOut(t, out)
	}
}

func TestGetProjects(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockProjectSearchResponse)
	}))
	defer srv.Close()
	j := testJiraClient(t, atlassian.Options{MaxRetries: 0})
	var out types.ProjectSearchResponse
	if err := j.doJSON(context.Background(), "GET", srv.URL, nil, &out); err != nil {
		t.Fatal(err)
	}
	if out.Total != 1 || len(out.Values) != 1 || out.Values[0].Key != "P1" {
		failOut(t, out)
	}
}

func TestProjectSearchParams(t *testing.T) {
	params := ProjectSearchParams{StartAt: 10, MaxResults: 5}
	if params.StartAt != 10 || params.MaxResults != 5 {
		t.Errorf(msgParamsFormat, params)
	}
}
