package constants

import "testing"

func TestJiraPathConstants_NonEmpty(t *testing.T) {
	if JiraPathMyself == "" || JiraPathProject == "" || JiraPathPermissionScheme == "" {
		t.Error("Jira path constants must be non-empty")
	}
}
