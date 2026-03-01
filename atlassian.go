package atlassian

import "github.com/surajrajput1024/go-atlassian-cloud/client"

type Config = client.Config
type Options = client.Options
type Client = client.Client
type APIError = client.APIError

var (
	NewClient             = client.NewClient
	DefaultOptions        = client.DefaultOptions
	IsRetryableStatusCode = client.IsRetryableStatusCode
	// Sentinel errors for errors.Is(err, atlassian.ErrNotFound) etc.
	ErrNotFound     = client.ErrNotFound
	ErrUnauthorized = client.ErrUnauthorized
	ErrForbidden    = client.ErrForbidden
	ErrBadRequest   = client.ErrBadRequest
)
