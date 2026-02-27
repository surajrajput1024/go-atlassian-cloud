package jira

import (
	"context"
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetWorkflowSchemeProjectAssociations(projectIDs []string) (*types.WorkflowSchemeProjectAssociationsResponse, error) {
	return j.WorkflowSchemeProjects.GetWorkflowSchemeProjectAssociations(projectIDs)
}

func (j *Client) AssignWorkflowSchemeToProject(projectID, workflowSchemeID string) error {
	return j.WorkflowSchemeProjects.AssignWorkflowSchemeToProject(projectID, workflowSchemeID)
}

func (s *WorkflowSchemeProjectService) GetWorkflowSchemeProjectAssociations(projectIDs []string) (*types.WorkflowSchemeProjectAssociationsResponse, error) {
	q := url.Values{}
	for _, id := range projectIDs {
		q.Add("projectId", id)
	}
	path := s.c.path(constants.JiraPathWorkflowSchemeProject)
	if len(q) > 0 {
		path = path + "?" + q.Encode()
	}
	var out types.WorkflowSchemeProjectAssociationsResponse
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *WorkflowSchemeProjectService) AssignWorkflowSchemeToProject(projectID, workflowSchemeID string) error {
	path := s.c.path(constants.JiraPathWorkflowSchemeProject)
	req := &types.WorkflowSchemeProjectAssignRequest{
		ProjectID:        projectID,
		WorkflowSchemeID: workflowSchemeID,
	}
	return s.c.doJSON(context.Background(), "PUT", path, req, nil)
}
