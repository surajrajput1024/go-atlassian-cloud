package jira

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajsinghrajput/go-atlassian-cloud/client"
	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func TestGetFields(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]types.FieldResponse{{ID: "summary", Name: "Summary"}})
	}))
	defer srv.Close()
	cfg := &atlassian.Config{Domain: "site.atlassian.net", Email: "u@e.com", APIToken: "tok"}
	cl, _ := atlassian.NewClient(cfg, atlassian.Options{MaxRetries: 0})
	j := New(cl)
	var out []types.FieldResponse
	if err := j.getJSON(srv.URL, &out); err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 || out[0].Name != "Summary" {
		t.Errorf("out = %+v", out)
	}
}
