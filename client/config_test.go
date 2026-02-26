package client

import "testing"

func TestConfig_Validate_Nil(t *testing.T) {
	var c *Config
	if err := c.Validate(); err == nil {
		t.Error("expected error for nil config")
	}
}

func TestConfig_Validate_Ok(t *testing.T) {
	c := &Config{Domain: "site.atlassian.net", Email: "a@b.com", APIToken: "tok"}
	if err := c.Validate(); err != nil {
		t.Errorf("expected no error: %v", err)
	}
}

func TestConfig_BaseURL(t *testing.T) {
	c := &Config{Domain: "site.atlassian.net"}
	if got := c.BaseURL(); got != "https://site.atlassian.net" {
		t.Errorf("BaseURL() = %q", got)
	}
}

func TestConfig_RestAPIURL(t *testing.T) {
	c := &Config{Domain: "site.atlassian.net"}
	if got := c.RestAPIURL(); got != "https://site.atlassian.net/rest/api/3" {
		t.Errorf("RestAPIURL() = %q", got)
	}
}
