package auth

import (
	"encoding/base64"
	"net/http"
)

func SetBasicAuth(req *http.Request, email, apiToken string) {
	raw := email + ":" + apiToken
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(raw)))
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
}
