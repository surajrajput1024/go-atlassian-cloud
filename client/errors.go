package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/surajrajput1024/go-atlassian-cloud/client/retry"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrBadRequest   = errors.New("bad request")
)

type APIError struct {
	StatusCode    int
	Body          []byte
	ErrorMessages []string
	Errors        map[string]string
}

func (e *APIError) Unwrap() error {
	switch e.StatusCode {
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusBadRequest:
		return ErrBadRequest
	default:
		return nil
	}
}

func (e *APIError) Error() string {
	if len(e.ErrorMessages) > 0 {
		return fmt.Sprintf("api error %d: %s", e.StatusCode, e.ErrorMessages[0])
	}
	if len(e.Errors) > 0 {
		for k, v := range e.Errors {
			return fmt.Sprintf("api error %d: %s: %s", e.StatusCode, k, v)
		}
	}
	return fmt.Sprintf("api error %d: %s", e.StatusCode, string(e.Body))
}

func wrapAPIError(ae *APIError) error {
	if ae == nil {
		return nil
	}
	if w := ae.Unwrap(); w != nil {
		return fmt.Errorf("%w: %s", w, ae.Error())
	}
	return ae
}

type atlassianErrorBody struct {
	ErrorMessages []string          `json:"errorMessages"`
	Errors        map[string]string `json:"errors"`
	Status        *int              `json:"status"`
}

func newAPIError(resp *http.Response) (*APIError, error) {
	if resp == nil {
		return &APIError{StatusCode: 0}, nil
	}
	if resp.Body == nil {
		return &APIError{StatusCode: resp.StatusCode}, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error response body: %w", err)
	}
	out := &APIError{StatusCode: resp.StatusCode, Body: body}
	var parsed atlassianErrorBody
	if err := json.Unmarshal(body, &parsed); err == nil {
		out.ErrorMessages = parsed.ErrorMessages
		out.Errors = parsed.Errors
	}
	return out, nil
}

func IsRetryableStatusCode(code int) bool {
	return retry.IsRetryableStatusCode(code)
}
