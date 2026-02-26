package client

import (
	"errors"
	"net/url"
	"strings"

	httputil "github.com/surajsinghrajput/go-atlassian-cloud/client/http"
)

type Config struct {
	Domain         string
	CloudID        string
	Email          string
	APIToken       string
	JiraAPIVersion string
}

func (c *Config) Validate() error {
	if c == nil {
		return errors.New(ErrConfigNil)
	}
	if empty(c.Domain) {
		return errors.New(ErrDomainRequired)
	}
	if !DomainRegex.MatchString(trim(c.Domain)) {
		return errors.New(ErrDomainInvalid)
	}
	if empty(c.Email) {
		return errors.New(ErrEmailRequired)
	}
	if empty(c.APIToken) {
		return errors.New(ErrAPITokenRequired)
	}
	ver := c.JiraVersion()
	if _, ok := SupportedJiraAPIVersions[ver]; !ok {
		return errors.New(ErrUnsupportedVersion + ": " + ver)
	}
	return nil
}

func (c *Config) JiraVersion() string {
	return orDefault(c.JiraAPIVersion, JiraAPIVersionDefault)
}

func (c *Config) BaseURL() string {
	domain := trim(c.Domain)
	if domain == "" {
		return ""
	}
	return "https://" + domain
}

func (c *Config) RestAPIURL() string {
	return c.BaseURL() + RestAPIPathPrefix + c.JiraVersion()
}

func (c *Config) RestAPI3URL() string {
	return c.BaseURL() + RestAPIPathPrefix + JiraAPIVersionDefault
}

func (c *Config) TenantInfoURL() string {
	return c.BaseURL() + TenantInfoPath
}

func (c *Config) ParseURL(path string) (*url.URL, error) {
	base := c.BaseURL()
	if base == "" {
		return nil, errors.New(ErrBaseURLEmpty)
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return url.Parse(path)
	}
	return httputil.ParseURL(base, path)
}
