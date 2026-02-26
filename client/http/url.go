package httputil

import (
	"net/url"
	"strings"
)

func JoinPath(base, path string) string {
	base = strings.TrimSpace(base)
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

func ParseURL(base, path string) (*url.URL, error) {
	joined := JoinPath(base, path)
	return url.Parse(joined)
}
