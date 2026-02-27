package util

import (
	"encoding/json"
)

// ParseStringOrNumber unmarshals a JSON value that may be a string or number into a string.
// Jira API sometimes returns ids as numbers; this avoids unmarshal errors.
func ParseStringOrNumber(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	if data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return "", err
		}
		return s, nil
	}
	var n json.Number
	if err := json.Unmarshal(data, &n); err != nil {
		return "", err
	}
	return n.String(), nil
}
