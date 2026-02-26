package client

import (
	"net/url"
	"strings"
)

func trim(s string) string {
	return strings.TrimSpace(s)
}

func empty(s string) bool {
	return trim(s) == ""
}

func orDefault(s, defaultVal string) string {
	if empty(s) {
		return defaultVal
	}
	return trim(s)
}

func joinPath(base, path string) string {
	base = trim(base)
	if base == "" {
		return path
	}
	path = strings.TrimSpace(path)
	if path == "" {
		return base
	}
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return path
	}
	if strings.HasPrefix(path, "/") {
		return base + path
	}
	return base + "/" + path
}

func parseURL(base, path string) (*url.URL, error) {
	joined := joinPath(base, path)
	return url.Parse(joined)
}
