package jira

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajsinghrajput/go-atlassian-cloud"
	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func TestGetPriorities(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]types.PriorityResponse{{ID: "1", Name: "High"}})
	}))
	defer srv.Close()
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.Options{MaxRetries: 0})
	j := New(cl)
	var out []types.PriorityResponse
	if err := j.getJSON(srv.URL, &out); err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 || out[0].Name != "High" {
		t.Errorf("out = %+v", out)
	}
}
