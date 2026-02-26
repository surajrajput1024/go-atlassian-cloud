package types

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
	Permissions []PermissionGrantInput `json:"permissions,omitempty"`
}
