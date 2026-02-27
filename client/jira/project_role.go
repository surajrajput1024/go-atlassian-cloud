package jira

import (
	"context"
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetProjectRoles(projectIDOrKey string) (types.ProjectRolesMapResponse, error) {
	return j.ProjectRoles.GetProjectRoles(projectIDOrKey)
}

func (j *Client) GetProjectRoleDetails(projectIDOrKey string) ([]types.ProjectRoleDetailsItem, error) {
	return j.ProjectRoles.GetProjectRoleDetails(projectIDOrKey)
}

func (j *Client) GetProjectRole(projectIDOrKey, roleID string) (*types.ProjectRoleResponse, error) {
	return j.ProjectRoles.GetProjectRole(projectIDOrKey, roleID)
}

func (j *Client) AddProjectRoleActors(projectIDOrKey, roleID string, req *types.ProjectRoleAddActorsRequest) (*types.ProjectRoleResponse, error) {
	return j.ProjectRoles.AddProjectRoleActors(projectIDOrKey, roleID, req)
}

func (j *Client) SetProjectRoleActors(projectIDOrKey, roleID string, req *types.ProjectRoleSetActorsRequest) (*types.ProjectRoleResponse, error) {
	return j.ProjectRoles.SetProjectRoleActors(projectIDOrKey, roleID, req)
}

func (j *Client) DeleteProjectRoleActors(projectIDOrKey, roleID string, user, group, groupID string) error {
	return j.ProjectRoles.DeleteProjectRoleActors(projectIDOrKey, roleID, user, group, groupID)
}

func (s *ProjectRoleService) GetProjectRoles(projectIDOrKey string) (types.ProjectRolesMapResponse, error) {
	var out types.ProjectRolesMapResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey), constants.JiraPathRole)
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *ProjectRoleService) GetProjectRoleDetails(projectIDOrKey string) ([]types.ProjectRoleDetailsItem, error) {
	var out []types.ProjectRoleDetailsItem
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey), constants.JiraPathRoleDetails)
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *ProjectRoleService) GetProjectRole(projectIDOrKey, roleID string) (*types.ProjectRoleResponse, error) {
	var out types.ProjectRoleResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey), constants.JiraPathRole, url.PathEscape(roleID))
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *ProjectRoleService) AddProjectRoleActors(projectIDOrKey, roleID string, req *types.ProjectRoleAddActorsRequest) (*types.ProjectRoleResponse, error) {
	var out types.ProjectRoleResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey), constants.JiraPathRole, url.PathEscape(roleID))
	if err := s.c.doJSON(context.Background(), "POST", path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *ProjectRoleService) SetProjectRoleActors(projectIDOrKey, roleID string, req *types.ProjectRoleSetActorsRequest) (*types.ProjectRoleResponse, error) {
	var out types.ProjectRoleResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey), constants.JiraPathRole, url.PathEscape(roleID))
	if err := s.c.doJSON(context.Background(), "PUT", path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *ProjectRoleService) DeleteProjectRoleActors(projectIDOrKey, roleID string, user, group, groupID string) error {
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey), constants.JiraPathRole, url.PathEscape(roleID))
	q := url.Values{}
	if user != "" {
		q.Set("user", user)
	}
	if group != "" {
		q.Set("group", group)
	}
	if groupID != "" {
		q.Set("groupId", groupID)
	}
	if len(q) > 0 {
		path = path + "?" + q.Encode()
	}
	return s.c.delete(context.Background(), path)
}
