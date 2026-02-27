package jira

import (
	"context"
	"net/url"

	"github.com/surajrajput1024/go-atlassian-cloud/constants"
	"github.com/surajrajput1024/go-atlassian-cloud/types"
	"github.com/surajrajput1024/go-atlassian-cloud/util"
)

type ProjectSearchParams struct {
	StartAt    int
	MaxResults int
	OrderBy    string
	Query      string
	TypeKey    string
	CategoryID int64
}

func (j *Client) GetProject(projectIDOrKey string) (*types.ProjectResponse, error) {
	return j.Projects.GetProject(projectIDOrKey)
}

func (j *Client) GetProjectWithContext(ctx context.Context, projectIDOrKey string) (*types.ProjectResponse, error) {
	return j.Projects.GetProjectWithContext(ctx, projectIDOrKey)
}

func (s *ProjectService) GetProject(projectIDOrKey string) (*types.ProjectResponse, error) {
	return s.GetProjectWithContext(context.Background(), projectIDOrKey)
}

func (s *ProjectService) GetProjectWithContext(ctx context.Context, projectIDOrKey string) (*types.ProjectResponse, error) {
	var out types.ProjectResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey))
	if err := s.c.doJSON(ctx, "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) GetProjects(params ProjectSearchParams) (*types.ProjectSearchResponse, error) {
	return j.Projects.GetProjects(params)
}

func (j *Client) GetProjectsWithContext(ctx context.Context, params ProjectSearchParams) (*types.ProjectSearchResponse, error) {
	return j.Projects.GetProjectsWithContext(ctx, params)
}

func (s *ProjectService) GetProjects(params ProjectSearchParams) (*types.ProjectSearchResponse, error) {
	return s.GetProjectsWithContext(context.Background(), params)
}

func (s *ProjectService) GetProjectsWithContext(ctx context.Context, params ProjectSearchParams) (*types.ProjectSearchResponse, error) {
	path := s.c.path(constants.JiraPathProjectSearch)
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
	if err := s.c.doJSON(ctx, "GET", path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) CreateProject(req *types.ProjectCreateRequest) (*types.ProjectCreateResponse, error) {
	return j.Projects.CreateProject(req)
}

func (j *Client) CreateProjectWithContext(ctx context.Context, req *types.ProjectCreateRequest) (*types.ProjectCreateResponse, error) {
	return j.Projects.CreateProjectWithContext(ctx, req)
}

func (s *ProjectService) CreateProject(req *types.ProjectCreateRequest) (*types.ProjectCreateResponse, error) {
	return s.CreateProjectWithContext(context.Background(), req)
}

func (s *ProjectService) CreateProjectWithContext(ctx context.Context, req *types.ProjectCreateRequest) (*types.ProjectCreateResponse, error) {
	var out types.ProjectCreateResponse
	if err := s.c.doJSON(ctx, "POST", s.c.path(constants.JiraPathProject), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) UpdateProject(projectIDOrKey string, req *types.ProjectUpdateRequest) (*types.ProjectResponse, error) {
	return j.Projects.UpdateProject(projectIDOrKey, req)
}

func (j *Client) UpdateProjectWithContext(ctx context.Context, projectIDOrKey string, req *types.ProjectUpdateRequest) (*types.ProjectResponse, error) {
	return j.Projects.UpdateProjectWithContext(ctx, projectIDOrKey, req)
}

func (s *ProjectService) UpdateProject(projectIDOrKey string, req *types.ProjectUpdateRequest) (*types.ProjectResponse, error) {
	return s.UpdateProjectWithContext(context.Background(), projectIDOrKey, req)
}

func (s *ProjectService) UpdateProjectWithContext(ctx context.Context, projectIDOrKey string, req *types.ProjectUpdateRequest) (*types.ProjectResponse, error) {
	var out types.ProjectResponse
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey))
	if err := s.c.doJSON(ctx, "PUT", path, req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (j *Client) DeleteProject(projectIDOrKey string) error {
	return j.Projects.DeleteProject(projectIDOrKey)
}

func (j *Client) DeleteProjectWithContext(ctx context.Context, projectIDOrKey string) error {
	return j.Projects.DeleteProjectWithContext(ctx, projectIDOrKey)
}

func (s *ProjectService) DeleteProject(projectIDOrKey string) error {
	return s.DeleteProjectWithContext(context.Background(), projectIDOrKey)
}

func (s *ProjectService) DeleteProjectWithContext(ctx context.Context, projectIDOrKey string) error {
	path := s.c.path(constants.JiraPathProject, url.PathEscape(projectIDOrKey))
	return s.c.delete(ctx, path)
}
