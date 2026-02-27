package jira

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWorkflowSchemeProjectAssociations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, MockWorkflowSchemeProjectAssociationsResponse)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	out, err := j.GetWorkflowSchemeProjectAssociations([]string{"10000"})
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Values) != 1 || out.Values[0].WorkflowScheme.ID != 10032 {
		failOut(t, out)
	}
}

func TestAssignWorkflowSchemeToProject(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf(msgMethodFormat, r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	j := testJiraClientWithServer(t, srv)
	err := j.AssignWorkflowSchemeToProject("10000", "10032")
	if err != nil {
		t.Fatal(err)
	}
}
