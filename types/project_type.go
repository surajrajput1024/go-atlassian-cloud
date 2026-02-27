package types

import (
	"encoding/json"

	"github.com/surajrajput1024/go-atlassian-cloud/util"
)

type AvatarUrls struct {
	Size16 string `json:"16x16"`
	Size24 string `json:"24x24"`
	Size32 string `json:"32x32"`
	Size48 string `json:"48x48"`
}

type UserRef struct {
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	Self        string `json:"self"`
}

type ProjectCategory struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Self        string `json:"self"`
}

func (c *ProjectCategory) UnmarshalJSON(data []byte) error {
	type projectCategory ProjectCategory
	var aux struct {
		projectCategory
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*c = ProjectCategory(aux.projectCategory)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	c.ID = id
	return nil
}

type ProjectResponse struct {
	ID              string           `json:"-"`
	Key             string           `json:"key"`
	Name            string           `json:"name"`
	Self            string           `json:"self"`
	Expand          string           `json:"expand"`
	Description     string           `json:"description"`
	Style           string           `json:"style"`
	Simplified      bool             `json:"simplified"`
	ProjectTypeKey  string           `json:"projectTypeKey"`
	IsPrivate       bool             `json:"isPrivate"`
	Properties      map[string]any   `json:"properties"`
	EntityID        string           `json:"entityId"`
	UUID            string           `json:"uuid"`
	ProjectCategory *ProjectCategory `json:"projectCategory"`
	AvatarUrls      *AvatarUrls      `json:"avatarUrls"`
	AssigneeType    string           `json:"assigneeType"`
	Lead            *UserRef         `json:"lead"`
	Email           string           `json:"email"`
}

func (r *ProjectResponse) UnmarshalJSON(data []byte) error {
	type projectResponse ProjectResponse
	var aux struct {
		projectResponse
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = ProjectResponse(aux.projectResponse)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

type ProjectSearchResponse struct {
	StartAt    int               `json:"startAt"`
	MaxResults int               `json:"maxResults"`
	Total      int               `json:"total"`
	IsLast     bool              `json:"isLast"`
	NextPage   string            `json:"nextPage"`
	Self       string            `json:"self"`
	Values     []ProjectResponse `json:"values"`
}

type ProjectCreateResponse struct {
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

func (r *ProjectCreateResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		ID   json.RawMessage `json:"id"`
		Key  string          `json:"key"`
		Self string          `json:"self"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	r.Key = raw.Key
	r.Self = raw.Self
	id, err := util.ParseStringOrNumber(raw.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

type ProjectCategoryResponse struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Self        string `json:"self"`
}

func (r *ProjectCategoryResponse) UnmarshalJSON(data []byte) error {
	type projectCategoryResponse ProjectCategoryResponse
	var aux struct {
		projectCategoryResponse
		ID json.RawMessage `json:"id"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = ProjectCategoryResponse(aux.projectCategoryResponse)
	id, err := util.ParseStringOrNumber(aux.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

type ProjectCreateRequest struct {
	Key                      string `json:"key"`
	Name                     string `json:"name"`
	ProjectTypeKey           string `json:"projectTypeKey"`
	ProjectTemplateKey       string `json:"projectTemplateKey,omitempty"`
	Description              string `json:"description,omitempty"`
	LeadAccountID            string `json:"leadAccountId,omitempty"`
	URL                      string `json:"url,omitempty"`
	AssigneeType             string `json:"assigneeType,omitempty"`
	AvatarID                 int64  `json:"avatarId,omitempty"`
	CategoryID               int64  `json:"categoryId,omitempty"`
	IssueSecurityScheme      int64  `json:"issueSecurityScheme,omitempty"`
	PermissionScheme         int64  `json:"permissionScheme,omitempty"`
	NotificationScheme       int64  `json:"notificationScheme,omitempty"`
	IssueTypeScheme          int64  `json:"issueTypeScheme,omitempty"`
	IssueTypeScreenScheme    int64  `json:"issueTypeScreenScheme,omitempty"`
	FieldConfigurationScheme int64  `json:"fieldConfigurationScheme,omitempty"`
}

type ProjectUpdateRequest struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	URL           string `json:"url,omitempty"`
	LeadAccountID string `json:"leadAccountId,omitempty"`
	AssigneeType  string `json:"assigneeType,omitempty"`
	CategoryID    int64  `json:"categoryId,omitempty"`
}

type ProjectCategoryCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type ProjectCategoryUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
