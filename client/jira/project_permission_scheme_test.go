package jira

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetProjectPermissionScheme(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockPermissionSchemeResponse)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.GetProjectPermissionScheme("PROJ")
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != 10000 || out.Name != "Example scheme" {
		failOut(t, out)
	}
}

func TestAssignPermissionSchemeToProject(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf(msgMethodFormat, r.Method)
		}
		writeJSON(w, http.StatusOK, types.PermissionSchemeResponse{ID: 10000, Name: "Example scheme"})
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.AssignPermissionSchemeToProject("PROJ", 10000)
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != 10000 {
		failOut(t, out)
	}
}
