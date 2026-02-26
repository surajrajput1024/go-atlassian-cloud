package client

import (
	"regexp"
	"time"
)

const (
	JiraAPIVersionDefault = "3"
	RestAPIPathPrefix      = "/rest/api/"
	TenantInfoPath         = "/_edge/tenant_info"
)

const (
	ErrConfigNil          = "client config is nil"
	ErrDomainRequired     = "domain is required"
	ErrDomainInvalid      = "domain must be a valid Atlassian Cloud host (e.g. your-site.atlassian.net)"
	ErrEmailRequired      = "email is required"
	ErrAPITokenRequired   = "api_token is required"
	ErrBaseURLEmpty       = "invalid config: empty base URL"
	ErrUnsupportedVersion = "unsupported jira_api_version"
)

const domainPatternRegex = `^[a-zA-Z0-9][a-zA-Z0-9.-]*\.atlassian\.net$`

var (
	DomainRegex              = regexp.MustCompile(domainPatternRegex)
	SupportedJiraAPIVersions = map[string]struct{}{JiraAPIVersionDefault: {}}
)

const (
	DefaultHTTPTimeout     = 30 * time.Second
	DefaultMaxRetries      = 3
	DefaultRetryBackoffMin = 500 * time.Millisecond
	DefaultRetryBackoffMax = 5 * time.Second
)
