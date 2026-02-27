package jira

import (
	"context"
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
)

func (j *Client) GetProjectCategories() ([]types.ProjectCategoryResponse, error) {
	var out []types.ProjectCategoryResponse
	if err := j.doJSON(context.Background(), "GET", j.path(constants.JiraPathProjectCategory), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (j *Client) GetProjectCategory(id string) (*types.ProjectCategoryResponse, error) {
	var out types.ProjectCategoryResponse
	path := j.path(constants.JiraPathProjectCategory, url.PathEscape(id))
	if err := j.doJSON(context.Background(), "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) CreateProjectCategory(req *types.ProjectCategoryCreateRequest) (*types.ProjectCategoryResponse, error) {
	var out types.ProjectCategoryResponse
	if err := j.doJSON(context.Background(), "POST", j.path(constants.JiraPathProjectCategory), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) UpdateProjectCategory(id string, req *types.ProjectCategoryUpdateRequest) (*types.ProjectCategoryResponse, error) {
	var out types.ProjectCategoryResponse
	path := j.path(constants.JiraPathProjectCategory, url.PathEscape(id))
	if err := j.doJSON(context.Background(), "PUT", path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) DeleteProjectCategory(id string) error {
	path := j.path(constants.JiraPathProjectCategory, url.PathEscape(id))
	return j.delete(context.Background(), path)
}
