package jira

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func TestGetGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockGroupResponse)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.GetGroup("group-uuid", "")
	if err != nil {
		t.Fatal(err)
	}
	if out.GroupID != "group-uuid" || out.Name != "jira-users" {
		failOut(t, out)
	}
}

func TestCreateGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf(msgMethodFormat, r.Method)
		}
		writeJSON(w, http.StatusCreated, MockGroupResponseCreated)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.CreateGroup(&types.GroupCreateRequest{Name: "power-users"})
	if err != nil {
		t.Fatal(err)
	}
	if out.Name != "power-users" || out.GroupID != "new-group-uuid" {
		failOut(t, out)
	}
}

func TestDeleteGroup(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf(msgMethodFormat, r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	err := j.DeleteGroup("group-uuid", "", "", "")
	if err != nil {
		t.Fatal(err)
	}
}
