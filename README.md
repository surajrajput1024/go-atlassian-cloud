# go-atlassian-cloud

Go client for Atlassian Cloud APIs (Jira, Confluence, etc.). Standalone module with no dependency on Terraform or any provider.

## Install

```bash
go get github.com/surajrajput1024/go-atlassian-cloud@v0.1.0
```

## Usage

Preferred (canonical paths):

```go
import (
	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/client/jira"
)

cfg := &atlassian.Config{
	Domain:   "your-site.atlassian.net",
	Email:    "you@example.com",
	APIToken: "your-api-token",
}
cl, err := atlassian.NewClient(cfg, atlassian.DefaultOptions())
// Or with functional options:
// cl, err = atlassian.NewClientWithOptions(cfg, atlassian.WithTimeout(10*time.Second), atlassian.WithRetries(2, 1*time.Second, 3*time.Second))
// ...
j := jira.New(cl)  // j implements same API; j.Projects, j.PermissionSchemes, etc. are available for direct service access
```

Backward-compatible (root re-exports): use `github.com/surajrajput1024/go-atlassian-cloud` and `.../jira` instead of `.../client` and `.../client/jira`.

```go
import (
	atlassian "github.com/surajrajput1024/go-atlassian-cloud"
	"github.com/surajrajput1024/go-atlassian-cloud/jira"
)

cfg := &atlassian.Config{
	Domain:   "your-site.atlassian.net",
	Email:    "you@example.com",
	APIToken: "your-api-token",
}
cl, err := atlassian.NewClient(cfg, atlassian.DefaultOptions())
if err != nil {
	log.Fatal(err)
}

j := jira.New(cl)
user, err := j.GetCurrentUser()
proj, err := j.GetProject("PROJ")
projects, err := j.GetProjects(jira.ProjectSearchParams{})
```

## Layout

```
client/
  auth/       Basic auth and request headers
  http/       URL parsing and path helpers (package httputil)
  retry/      Backoff and retryable status logic
  jira/       Jira API client: New(cl), projects, categories, issue types, statuses, priorities, fields,
              permission schemes and grants, project permission scheme attachment, project role actors,
              groups, workflow scheme project associations
  *.go        Config, Options, Client, ResolveCloudID, errors
internal/
examples/
docs/
types/        Request/response structs for Jira v3
constants/    Jira path segments
util/         Helpers (Int64String, IntString)
```

| Import | Purpose |
|--------|--------|
| `.../client` | Config, NewClient, NewClientWithOptions, DefaultOptions, WithTimeout, WithTransport, WithRetries, RESTDoer, DoJSON, GetWithContext, DoWithContext, ResolveCloudID, APIError, ErrNotFound, ErrUnauthorized, ErrForbidden, ErrBadRequest |
| `.../client/jira` | jira.New(do), Client.Projects, PermissionSchemes, ProjectPermissionScheme, ProjectRoles, Groups, WorkflowSchemeProjects; GetProject, GetProjects, GetCurrentUser, permission schemes/grants, project permission scheme, project role actors, groups, workflow scheme project |
| `.../types` | Request/response structs |
| `.../constants` | Path constants (re-exports from constants/jira) |
| `.../constants/jira` | Jira path constants by product |
| `.../util` | Int64String, IntString |

## Use in a Terraform provider

In your provider's `go.mod`:

```go
require github.com/surajrajput1024/go-atlassian-cloud v0.1.0
```

For local development (provider and client in sibling dirs):

```go
replace github.com/surajrajput1024/go-atlassian-cloud => ../go-atlassian-cloud
```

In provider code:

```go
import (
	atlassian "github.com/surajrajput1024/go-atlassian-cloud/client"
	"github.com/surajrajput1024/go-atlassian-cloud/client/jira"
)

cfg := &atlassian.Config{
	Domain:   providerConfig.Domain,
	Email:    providerConfig.Email,
	APIToken: providerConfig.APIToken,
}
cl, err := atlassian.NewClient(cfg, atlassian.DefaultOptions())
if err != nil {
	return err
}
j := jira.New(cl)
// Use j.GetProject(), j.GetProjects(), etc. in resources/data sources
```

## Versioning

Releases follow [semver](https://semver.org/). Merges to `main` produce an automatic patch-version tag and GitHub Release. Pin in your module with a version:

```bash
go get github.com/surajrajput1024/go-atlassian-cloud@v<version>
```

## Contributing

Contributions are via **pull requests** only; direct pushes to `main` are disabled. See [CONTRIBUTING.md](CONTRIBUTING.md) for how to contribute and [.github/README.md](.github/README.md) for PR guidelines and the pull request template.
