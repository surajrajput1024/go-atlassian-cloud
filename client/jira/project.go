package jira

import (
	"net/url"

	"github.com/surajsinghrajput/go-atlassian-cloud/constants"
	"github.com/surajsinghrajput/go-atlassian-cloud/types"
	"github.com/surajsinghrajput/go-atlassian-cloud/util"
)

func (j *Client) GetProject(projectIDOrKey string) (*types.ProjectResponse, error) {
	var out types.ProjectResponse
	path := j.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey))
	if err := j.getJSON(path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type ProjectSearchParams struct {
	StartAt    int
	MaxResults int
	OrderBy    string
	Query      string
	TypeKey    string
	CategoryID int64
}

func (j *Client) GetProjects(params ProjectSearchParams) (*types.ProjectSearchResponse, error) {
	path := j.path(constants.JiraPathProjectSearch)
	q := url.Values{}
	if params.StartAt > 0 {
		q.Set("startAt", util.IntString(params.StartAt))
	}
	if params.MaxResults > 0 {
		q.Set("maxResults", util.IntString(params.MaxResults))
	}
	if params.OrderBy != "" {
		q.Set("orderBy", params.OrderBy)
	}
	if params.Query != "" {
		q.Set("query", params.Query)
	}
	if params.TypeKey != "" {
		q.Set("typeKey", params.TypeKey)
	}
	if params.CategoryID != 0 {
		q.Set("categoryId", util.Int64String(params.CategoryID))
	}
	if len(q) > 0 {
		path = path + "?" + q.Encode()
	}
	var out types.ProjectSearchResponse
	if err := j.getJSON(path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) CreateProject(req *types.ProjectCreateRequest) (*types.ProjectCreateResponse, error) {
	var out types.ProjectCreateResponse
	if err := j.post(j.path(constants.JiraPathProject), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) UpdateProject(projectIDOrKey string, req *types.ProjectUpdateRequest) (*types.ProjectResponse, error) {
	var out types.ProjectResponse
	path := j.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey))
	if err := j.put(path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) DeleteProject(projectIDOrKey string) error {
	path := j.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey))
	return j.delete(path)
}
