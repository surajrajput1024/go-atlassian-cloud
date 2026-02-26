package atlassian

import "github.com/surajrajput1024/go-atlassian-cloud/client"

type Config = client.Config
type Options = client.Options
type Client = client.Client
type APIError = client.APIError

var NewClient = client.NewClient
var DefaultOptions = client.DefaultOptions
var IsRetryableStatusCode = client.IsRetryableStatusCode
