package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	e := &APIError{StatusCode: 404, ErrorMessages: []string{"not found"}}
	if e.Error() == "" {
		t.Error("Error() should not be empty")
	}
}

func TestIsRetryableStatusCode(t *testing.T) {
	if !IsRetryableStatusCode(500) || !IsRetryableStatusCode(429) {
		t.Error("5xx and 429 should be retryable")
	}
	if IsRetryableStatusCode(200) || IsRetryableStatusCode(404) {
		t.Error("2xx and 4xx (except 429) should not be retryable")
	}
}

func TestNewAPIError_WithBody(t *testing.T) {
	rec := httptest.NewRecorder()
	rec.WriteHeader(400)
	rec.Body.WriteString(`{"errorMessages":["bad request"]}`)
	resp := rec.Result()
	ae, err := newAPIError(resp)
	if err != nil {
		t.Fatal(err)
	}
	if ae.StatusCode != 400 || len(ae.ErrorMessages) != 1 {
		t.Errorf("StatusCode=%d ErrorMessages=%v", ae.StatusCode, ae.ErrorMessages)
	}
}

func TestWrapAPIError_IsAndAs(t *testing.T) {
	ae := &APIError{StatusCode: http.StatusNotFound, ErrorMessages: []string{"not found"}}
	err := wrapAPIError(ae)
	if err == nil {
		t.Fatal("wrapAPIError must not return nil for non-nil APIError")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Error("errors.Is(err, ErrNotFound) should be true")
	}
	var got *APIError
	if !errors.As(err, &got) || got != ae {
		t.Errorf("errors.As(err, &apiErr) should recover *APIError; got %v", got)
	}
}
