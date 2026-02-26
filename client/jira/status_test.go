package jira

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetStatuses(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]types.StatusResponse{{ID: "1", Name: "To Do"}})
	}))
	defer srv.Close()
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.Options{MaxRetries: 0})
	j := New(cl)
	var out []types.StatusResponse
	if err := j.getJSON(srv.URL, &out); err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 || out[0].Name != "To Do" {
		t.Errorf("out = %+v", out)
	}
}
