package client

import "strings"

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
