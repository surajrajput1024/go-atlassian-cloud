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
	Permissions []PermissionGrantInput `json:"permissions,omitempty"`
}
