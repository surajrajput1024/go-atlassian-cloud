package jira

import (
	"context"
	"net/http"
	"strings"

	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
)

type Client struct {
	do atlassian.RESTDoer

	Projects                *ProjectService
	PermissionSchemes        *PermissionSchemeService
	ProjectPermissionScheme *ProjectPermissionSchemeService
	ProjectRoles            *ProjectRoleService
	Groups                  *GroupService
	WorkflowSchemeProjects  *WorkflowSchemeProjectService
}

type ProjectService struct{ c *Client }
type PermissionSchemeService struct{ c *Client }
type ProjectPermissionSchemeService struct{ c *Client }
type ProjectRoleService struct{ c *Client }
type GroupService struct{ c *Client }
type WorkflowSchemeProjectService struct{ c *Client }

func New(do atlassian.RESTDoer) *Client {
	c := &Client{do: do}
	c.Projects = &ProjectService{c: c}
	c.PermissionSchemes = &PermissionSchemeService{c: c}
	c.ProjectPermissionScheme = &ProjectPermissionSchemeService{c: c}
	c.ProjectRoles = &ProjectRoleService{c: c}
	c.Groups = &GroupService{c: c}
	c.WorkflowSchemeProjects = &WorkflowSchemeProjectService{c: c}
	return c
}

func (c *Client) path(segments ...string) string {
	base := c.do.RestAPIURL()
	for _, s := range segments {
		base = strings.TrimSuffix(base, "/") + "/" + strings.TrimPrefix(s, "/")
	}
	return base
}

func (c *Client) doJSON(ctx context.Context, method, path string, body interface{}, out interface{}) error {
	return c.do.DoJSON(ctx, method, path, body, out)
}

func (c *Client) delete(ctx context.Context, path string) error {
	return c.do.DoJSON(ctx, http.MethodDelete, path, nil, nil)
}
