package types

import (
	"encoding/json"

	"github.com/surajrajput1024/go-atlassian-cloud/util"
)

type CurrentUserResponse struct {
	AccountID    string     `json:"accountId"`
	AccountType  string     `json:"accountType"`
	Active       bool       `json:"active"`
	AvatarUrls   AvatarUrls `json:"avatarUrls"`
	DisplayName  string     `json:"displayName"`
	EmailAddress string     `json:"emailAddress"`
	Groups       GroupItems `json:"groups"`
	Key          string     `json:"key"`
	Name         string     `json:"name"`
	Self         string     `json:"self"`
	TimeZone     string     `json:"timeZone"`
}

type GroupItems struct {
	Items []interface{} `json:"items"`
	Size  int           `json:"size"`
}

type StatusResponse struct {
	ID             string          `json:"-"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Self           string          `json:"self"`
	StatusCategory *StatusCategory `json:"statusCategory"`
}

func (r *StatusResponse) UnmarshalJSON(data []byte) error {
	type statusResponse StatusResponse
	var aux struct {
		statusResponse
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = StatusResponse(aux.statusResponse)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

type StatusCategory struct {
	ID        string `json:"-"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Self      string `json:"self"`
	ColorName string `json:"colorName"`
}

func (c *StatusCategory) UnmarshalJSON(data []byte) error {
	type statusCategory StatusCategory
	var aux struct {
		statusCategory
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*c = StatusCategory(aux.statusCategory)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	c.ID = id
	return nil
}

type PriorityResponse struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	Self        string `json:"self"`
	IconURL     string `json:"iconUrl"`
	StatusColor string `json:"statusColor"`
}

func (r *PriorityResponse) UnmarshalJSON(data []byte) error {
	type priorityResponse PriorityResponse
	var aux struct {
		priorityResponse
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = PriorityResponse(aux.priorityResponse)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

type FieldResponse struct {
	ID          string       `json:"-"`
	Name        string       `json:"name"`
	Custom      bool         `json:"custom"`
	Schema      *FieldSchema `json:"schema"`
	Description string       `json:"description"`
	Self        string       `json:"self"`
}

func (r *FieldResponse) UnmarshalJSON(data []byte) error {
	type fieldResponse FieldResponse
	var aux struct {
		fieldResponse
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = FieldResponse(aux.fieldResponse)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

type FieldSchema struct {
	Type     string `json:"type"`
	System   string `json:"system,omitempty"`
	Custom   string `json:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty"`
}

type PermissionSchemeResponse struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Self        string            `json:"self"`
	Expand      string            `json:"expand"`
	Permissions []PermissionGrant `json:"permissions,omitempty"`
}

type PermissionSchemeListResponse struct {
	PermissionSchemes []PermissionSchemeResponse `json:"permissionSchemes"`
}

type PermissionGrant struct {
	ID         int              `json:"id"`
	Permission string           `json:"permission"`
	Holder     PermissionHolder `json:"holder"`
	Self       string           `json:"self"`
}

type PermissionHolder struct {
	Type      string `json:"type"`
	Parameter string `json:"parameter"`
	Value     string `json:"value"`
	Expand    string `json:"expand,omitempty"`
}

type PermissionGrantsResponse struct {
	Expand      string            `json:"expand"`
	Permissions []PermissionGrant `json:"permissions"`
}

// ProjectRolesMapResponse is the response from GET project/{id}/role: role name -> role URL.
type ProjectRolesMapResponse map[string]string

// ProjectRoleResponse is the full project role with actors (GET project/{id}/role/{roleId}).
type ProjectRoleResponse struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Self        string             `json:"self"`
	Actors      []ProjectRoleActor `json:"actors"`
	Scope       *ProjectRoleScope  `json:"scope,omitempty"`
}

// ProjectRoleActor represents a user or group in a project role.
type ProjectRoleActor struct {
	ID          int64               `json:"id"`
	DisplayName string              `json:"displayName"`
	Name        string              `json:"name"`
	Type        string              `json:"type"`
	ActorUser   *ProjectRoleUser    `json:"actorUser,omitempty"`
	ActorGroup  *ProjectRoleGroup   `json:"actorGroup,omitempty"`
}

// ProjectRoleUser is the user part of an actor.
type ProjectRoleUser struct {
	AccountID string `json:"accountId"`
}

// ProjectRoleGroup is the group part of an actor.
type ProjectRoleGroup struct {
	GroupID     string `json:"groupId"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// ProjectRoleScope is the scope of a project role.
type ProjectRoleScope struct {
	Type    string         `json:"type"`
	Project *ProjectRef    `json:"project,omitempty"`
}

// ProjectRef is a minimal project reference.
type ProjectRef struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Name string `json:"name"`
}

// ProjectRoleDetailsItem is one entry from GET project/{id}/roledetails.
type ProjectRoleDetailsItem struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Self             string `json:"self"`
	Admin            bool   `json:"admin"`
	Default          bool   `json:"default"`
	RoleConfigurable bool   `json:"roleConfigurable"`
	TranslatedName   string `json:"translatedName"`
	Type             string `json:"type"`
}

// GroupResponse is the response from GET/POST group (group details and optional users).
type GroupResponse struct {
	GroupID string       `json:"groupId"`
	Name    string       `json:"name"`
	Self    string       `json:"self"`
	Expand  string       `json:"expand,omitempty"`
	Users   *GroupUsers  `json:"users,omitempty"`
}

// GroupUsers is the paginated users list in a group.
type GroupUsers struct {
	Size       int        `json:"size"`
	StartIndex int        `json:"start-index"`
	EndIndex   int        `json:"end-index"`
	MaxResults int        `json:"max-results"`
	Items      []UserRef  `json:"items,omitempty"`
}

// WorkflowSchemeProjectAssociationsResponse is the response from GET workflowscheme/project.
type WorkflowSchemeProjectAssociationsResponse struct {
	Values []WorkflowSchemeProjectAssociation `json:"values"`
}

// WorkflowSchemeProjectAssociation links a workflow scheme to project IDs.
type WorkflowSchemeProjectAssociation struct {
	ProjectIDs     []string         `json:"projectIds"`
	WorkflowScheme WorkflowSchemeRef `json:"workflowScheme"`
}

// WorkflowSchemeRef is a minimal workflow scheme reference.
type WorkflowSchemeRef struct {
	ID                int64             `json:"id"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Self              string            `json:"self"`
	DefaultWorkflow   string            `json:"defaultWorkflow,omitempty"`
	IssueTypeMappings map[string]string  `json:"issueTypeMappings,omitempty"`
}
