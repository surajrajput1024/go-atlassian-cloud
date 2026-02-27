package jira

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetProjectRoles(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockProjectRolesMapResponse)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.GetProjectRoles("PROJ")
	if err != nil {
		t.Fatal(err)
	}
	if out["Developers"] == "" {
		failOut(t, out)
	}
}

func TestGetProjectRole(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockProjectRoleResponse)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.GetProjectRole("PROJ", "10000")
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != 10000 || out.Name != "Developers" {
		failOut(t, out)
	}
}

func TestAddProjectRoleActors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf(msgMethodFormat, r.Method)
		}
		writeJSON(w, http.StatusOK, MockProjectRoleResponseWithActors)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	req := &types.ProjectRoleAddActorsRequest{GroupID: []string{"group-id-1"}}
	out, err := j.AddProjectRoleActors("PROJ", "10000", req)
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != 10000 || len(out.Actors) != 1 {
		failOut(t, out)
	}
}

func TestDeleteProjectRoleActors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf(msgMethodFormat, r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	err := j.DeleteProjectRoleActors("PROJ", "10000", "", "jira-users", "")
	if err != nil {
		t.Fatal(err)
	}
}
