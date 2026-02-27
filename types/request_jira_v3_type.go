package types

type PermissionHolderInput struct {
	Type      string `json:"type"`
	Parameter string `json:"parameter,omitempty"`
	Value     string `json:"value,omitempty"`
}

type PermissionGrantInput struct {
	Permission string               `json:"permission"`
	Holder     PermissionHolderInput `json:"holder"`
}

type PermissionSchemeCreateRequest struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description,omitempty"`
	Permissions []PermissionGrantInput `json:"permissions,omitempty"`
}

type PermissionSchemeUpdateRequest struct {
	Name        string                  `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	Permissions []PermissionGrantInput  `json:"permissions,omitempty"`
}

// ProjectPermissionSchemeAssignRequest is the body for assigning a permission scheme to a project.
type ProjectPermissionSchemeAssignRequest struct {
	ID int64 `json:"id"`
}

// ProjectRoleAddActorsRequest is the body for adding actors (users/groups) to a project role.
// At least one of Group, GroupID, or User must be non-empty. Group is group name; GroupID is preferred (stable).
type ProjectRoleAddActorsRequest struct {
	Group   []string `json:"group,omitempty"`
	GroupID []string `json:"groupId,omitempty"`
	User    []string `json:"user,omitempty"`
}

// ProjectRoleSetActorsRequest replaces all actors for a project role. Keys are actor type identifiers
// (e.g. "atlassian-group-role-actor-id", "atlassian-user-role-actor"); values are arrays of group ID or account ID.
type ProjectRoleSetActorsRequest struct {
	CategorisedActors map[string][]string `json:"categorisedActors"`
}

// GroupCreateRequest is the body for creating a Jira group.
type GroupCreateRequest struct {
	Name string `json:"name"`
}

// WorkflowSchemeProjectAssignRequest is the body for assigning a workflow scheme to a project.
type WorkflowSchemeProjectAssignRequest struct {
	ProjectID        string `json:"projectId"`
	WorkflowSchemeID string `json:"workflowSchemeId"`
}
