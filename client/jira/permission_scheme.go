package jira

import (
	"net/url"

	"github.com/surajsinghrajput/go-atlassian-cloud/constants"
	"github.com/surajsinghrajput/go-atlassian-cloud/types"
)

func (j *Client) GetPermissionSchemes() (*types.PermissionSchemeListResponse, error) {
	var out types.PermissionSchemeListResponse
	if err := j.getJSON(j.path(constants.JiraPathPermissionScheme), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) GetPermissionScheme(id string) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	path := j.path(constants.JiraPathPermissionScheme, url.PathEscape(id))
	if err := j.getJSON(path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) CreatePermissionScheme(req *types.PermissionSchemeCreateRequest) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	if err := j.post(j.path(constants.JiraPathPermissionScheme), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) UpdatePermissionScheme(id string, req *types.PermissionSchemeUpdateRequest) (*types.PermissionSchemeResponse, error) {
	var out types.PermissionSchemeResponse
	path := j.path(constants.JiraPathPermissionScheme, url.PathEscape(id))
	if err := j.put(path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) DeletePermissionScheme(id string) error {
	path := j.path(constants.JiraPathPermissionScheme, url.PathEscape(id))
	return j.delete(path)
}
