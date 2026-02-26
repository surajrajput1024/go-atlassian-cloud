package types

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

type AvatarUrls struct {
	Size16 string `json:"16x16"`
	Size24 string `json:"24x24"`
	Size32 string `json:"32x32"`
	Size48 string `json:"48x48"`
}

type GroupItems struct {
	Items []interface{} `json:"items"`
	Size  int           `json:"size"`
}

type ProjectResponse struct {
	ID              string           `json:"id"`
	Key             string           `json:"key"`
	Name            string           `json:"name"`
	Self            string           `json:"self"`
	Description     string           `json:"description"`
	Style           string           `json:"style"`
	Simplified      bool             `json:"simplified"`
	ProjectCategory *ProjectCategory `json:"projectCategory"`
	AvatarUrls      *AvatarUrls      `json:"avatarUrls"`
	AssigneeType    string           `json:"assigneeType"`
	Lead            *UserRef         `json:"lead"`
	Email           string           `json:"email"`
}

type ProjectCategory struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Self        string `json:"self"`
}

type UserRef struct {
	AccountID   string `json:"accountId"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	Self        string `json:"self"`
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

type IssueTypeResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Self           string `json:"self"`
	Subtask        bool   `json:"subtask"`
	HierarchyLevel int    `json:"hierarchyLevel"`
	IconURL        string `json:"iconUrl"`
}

type StatusResponse struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Self           string         `json:"self"`
	StatusCategory *StatusCategory `json:"statusCategory"`
}

type StatusCategory struct {
	ID        string `json:"id"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Self      string `json:"self"`
	ColorName string `json:"colorName"`
}

type PriorityResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Self        string `json:"self"`
	IconURL     string `json:"iconUrl"`
	StatusColor string `json:"statusColor"`
}

type FieldResponse struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Custom      bool         `json:"custom"`
	Schema      *FieldSchema `json:"schema"`
	Description string       `json:"description"`
	Self        string       `json:"self"`
}

type FieldSchema struct {
	Type     string `json:"type"`
	System   string `json:"system,omitempty"`
	Custom   string `json:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty"`
}

type ProjectCategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Self        string `json:"self"`
}

type PermissionSchemeResponse struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Self        string             `json:"self"`
	Expand      string             `json:"expand"`
	Permissions []PermissionGrant  `json:"permissions,omitempty"`
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
