package jira

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetPermissionSchemes(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockPermissionSchemeListResponse)
	}))
	defer srv.Close()
	j := testJiraClient(t, atlassian.Options{MaxRetries: 0})
	var out types.PermissionSchemeListResponse
	if err := j.doJSON(context.Background(), "GET", srv.URL, nil, &out); err != nil {
		t.Fatal(err)
	}
	if len(out.PermissionSchemes) != 1 || out.PermissionSchemes[0].Name != "Default" {
		failOut(t, out)
	}
}

func TestGetPermissionSchemeGrants(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockPermissionGrantsResponse)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.GetPermissionSchemeGrants("10000")
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Permissions) != 1 || out.Permissions[0].Permission != "BROWSE_PROJECTS" {
		failOut(t, out)
	}
}

func TestCreatePermissionGrant(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf(msgMethodFormat, r.Method)
		}
		writeJSON(w, http.StatusOK, MockPermissionGrant)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	req := &types.PermissionGrantInput{
		Permission: "BROWSE_PROJECTS",
		Holder:     types.PermissionHolderInput{Type: "group", Parameter: "jira-users"},
	}
	out, err := j.CreatePermissionGrant("10000", req)
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != 100 || out.Permission != "BROWSE_PROJECTS" {
		failOut(t, out)
	}
}
