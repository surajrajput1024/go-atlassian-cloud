package jira

import (
	"context"
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetPermissionSchemes() (*types.PermissionSchemeListResponse, error) {
	return j.PermissionSchemes.GetPermissionSchemes()
}

func (j *Client) GetPermissionScheme(id string) (*types.PermissionSchemeResponse, error) {
	return j.PermissionSchemes.GetPermissionScheme(id)
}

func (j *Client) CreatePermissionScheme(req *types.PermissionSchemeCreateRequest) (*types.PermissionSchemeResponse, error) {
	return j.PermissionSchemes.CreatePermissionScheme(req)
}

func (j *Client) UpdatePermissionScheme(id string, req *types.PermissionSchemeUpdateRequest) (*types.PermissionSchemeResponse, error) {
	return j.PermissionSchemes.UpdatePermissionScheme(id, req)
}

func (j *Client) DeletePermissionScheme(id string) error {
	return j.PermissionSchemes.DeletePermissionScheme(id)
}

func (j *Client) GetPermissionSchemeGrants(schemeID string) (*types.PermissionGrantsResponse, error) {
	return j.PermissionSchemes.GetPermissionSchemeGrants(schemeID)
}

func (j *Client) CreatePermissionGrant(schemeID string, req *types.PermissionGrantInput) (*types.PermissionGrant, error) {
	return j.PermissionSchemes.CreatePermissionGrant(schemeID, req)
}

func (j *Client) GetPermissionGrant(schemeID, grantID string) (*types.PermissionGrant, error) {
	return j.PermissionSchemes.GetPermissionGrant(schemeID, grantID)
}

func (j *Client) DeletePermissionGrant(schemeID, grantID string) error {
	return j.PermissionSchemes.DeletePermissionGrant(schemeID, grantID)
}

func (s *PermissionSchemeService) GetPermissionSchemes() (*types.PermissionSchemeListResponse, error) {
	var out types.PermissionSchemeListResponse
	if err := s.c.doJSON(context.Background(), "GET", s.c.path(constants.JiraPathPermissionScheme), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PermissionSchemeService) GetPermissionScheme(id string) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	path := s.c.path(constants.JiraPathPermissionScheme, url.PathEscape(id))
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PermissionSchemeService) CreatePermissionScheme(req *types.PermissionSchemeCreateRequest) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	if err := s.c.doJSON(context.Background(), "POST", s.c.path(constants.JiraPathPermissionScheme), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PermissionSchemeService) UpdatePermissionScheme(id string, req *types.PermissionSchemeUpdateRequest) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	path := s.c.path(constants.JiraPathPermissionScheme, url.PathEscape(id))
	if err := s.c.doJSON(context.Background(), "PUT", path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PermissionSchemeService) DeletePermissionScheme(id string) error {
	path := s.c.path(constants.JiraPathPermissionScheme, url.PathEscape(id))
	return s.c.delete(context.Background(), path)
}

func (s *PermissionSchemeService) GetPermissionSchemeGrants(schemeID string) (*types.PermissionGrantsResponse, error) {
	var out types.PermissionGrantsResponse
	path := s.c.path(constants.JiraPathPermissionScheme, url.PathEscape(schemeID), constants.JiraPathPermission)
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PermissionSchemeService) CreatePermissionGrant(schemeID string, req *types.PermissionGrantInput) (*types.PermissionGrant, error) {
	var out types.PermissionGrant
	path := s.c.path(constants.JiraPathPermissionScheme, url.PathEscape(schemeID), constants.JiraPathPermission)
	if err := s.c.doJSON(context.Background(), "POST", path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PermissionSchemeService) GetPermissionGrant(schemeID, grantID string) (*types.PermissionGrant, error) {
	var out types.PermissionGrant
	path := s.c.path(constants.JiraPathPermissionScheme, url.PathEscape(schemeID), constants.JiraPathPermission, url.PathEscape(grantID))
	if err := s.c.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PermissionSchemeService) DeletePermissionGrant(schemeID, grantID string) error {
	path := s.c.path(constants.JiraPathPermissionScheme, url.PathEscape(schemeID), constants.JiraPathPermission, url.PathEscape(grantID))
	return s.c.delete(context.Background(), path)
}
