package jira

import (
	"context"
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetProjectPermissionScheme(projectKeyOrID string) (*types.PermissionSchemeResponse, error) {
	return j.ProjectPermissionScheme.Get(projectKeyOrID)
}

func (j *Client) AssignPermissionSchemeToProject(projectKeyOrID string, schemeID int64) (*types.PermissionSchemeResponse, error) {
	return j.ProjectPermissionScheme.Assign(projectKeyOrID, schemeID)
}

func (s *ProjectPermissionSchemeService) Get(projectKeyOrID string) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectKeyOrID), constants.JiraPathProjectPermissionScheme)
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *ProjectPermissionSchemeService) Assign(projectKeyOrID string, schemeID int64) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectKeyOrID), constants.JiraPathProjectPermissionScheme)
	req := &types.ProjectPermissionSchemeAssignRequest{ID: schemeID}
	if err := s.c.doJSON(context.Background(), "PUT", path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
