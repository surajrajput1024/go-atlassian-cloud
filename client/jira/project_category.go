package jira

import (
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetProjectCategories() ([]types.ProjectCategoryResponse, error) {
	var out []types.ProjectCategoryResponse
	if err := j.getJSON(j.path(constants.JiraPathProjectCategory), &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (j *Client) GetProjectCategory(id string) (*types.ProjectCategoryResponse, error) {
	var out types.ProjectCategoryResponse
	path := j.path(constants.JiraPathProjectCategory, url.PathEscape(id))
	if err := j.getJSON(path, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) CreateProjectCategory(req *types.ProjectCategoryCreateRequest) (*types.ProjectCategoryResponse, error) {
	var out types.ProjectCategoryResponse
	if err := j.post(j.path(constants.JiraPathProjectCategory), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) UpdateProjectCategory(id string, req *types.ProjectCategoryUpdateRequest) (*types.ProjectCategoryResponse, error) {
	var out types.ProjectCategoryResponse
	path := j.path(constants.JiraPathProjectCategory, url.PathEscape(id))
	if err := j.put(path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) DeleteProjectCategory(id string) error {
	path := j.path(constants.JiraPathProjectCategory, url.PathEscape(id))
	return j.delete(path)
}
