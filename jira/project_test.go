package jira

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajsinghrajput/go-atlassian-cloud"
	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func TestGetProject(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types.ProjectResponse{ID: "10000", Key: "PROJ", Name: "Test"})
	}))
	defer srv.Close()
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.Options{MaxRetries: 0})
	j := New(cl)
	var out types.ProjectResponse
	if err := j.getJSON(srv.URL, &out); err != nil {
		t.Fatal(err)
	}
	if out.Key != "PROJ" || out.Name != "Test" {
		t.Errorf("out = %+v", out)
	}
}

func TestGetProjects(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(types.ProjectSearchResponse{Total: 1, Values: []types.ProjectResponse{{Key: "P1"}}})
	}))
	defer srv.Close()
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.Options{MaxRetries: 0})
	j := New(cl)
	var out types.ProjectSearchResponse
	if err := j.getJSON(srv.URL, &out); err != nil {
		t.Fatal(err)
	}
	if out.Total != 1 || len(out.Values) != 1 || out.Values[0].Key != "P1" {
		t.Errorf("out = %+v", out)
	}
}

func TestProjectSearchParams(t *testing.T) {
	params := ProjectSearchParams{StartAt: 10, MaxResults: 5}
	if params.StartAt != 10 || params.MaxResults != 5 {
		t.Errorf("params = %+v", params)
	}
}
