# go-atlassian-cloud

Go client for Atlassian Cloud APIs (Jira, Confluence, etc.). Standalone module with no dependency on Terraform or any provider.

## Install

```bash
go get github.com/surajsinghrajput/go-atlassian-cloud@v0.1.0
```

## Usage

```go
import (
	atlassian "github.com/surajsinghrajput/go-atlassian-cloud"
	"github.com/surajsinghrajput/go-atlassian-cloud/jira"
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

## Module layout

| Path | Purpose |
|------|--------|
| Root | `Config`, `NewClient`, `DefaultOptions`, `ResolveCloudID`, `RestAPIURL`, `Get` |
| `jira/` | Jira API client: `New(cl)`, projects, project categories, issue types, statuses, priorities, fields, permission schemes |
| `types/` | Request/response structs for Jira v3 |
| `constants/` | Path segments and shared constants |
| `util/` | Helpers (e.g. `Int64String`, `IntString`) |

## Use in a Terraform provider

In your provider's `go.mod`:

```go
require github.com/surajsinghrajput/go-atlassian-cloud v0.1.0
```

For local development (provider and client in sibling dirs):

```go
replace github.com/surajsinghrajput/go-atlassian-cloud => ../go-atlassian-cloud
```

In provider code, build the client from your config and pass it to Jira:

```go
import (
	atlassian "github.com/surajsinghrajput/go-atlassian-cloud"
	"github.com/surajsinghrajput/go-atlassian-cloud/jira"
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
go get github.com/surajsinghrajput/go-atlassian-cloud@v<version>
```

## Contributing

Contributions are via **pull requests** only; direct pushes to `main` are disabled. See [CONTRIBUTING.md](CONTRIBUTING.md) for how to contribute and [.github/README.md](.github/README.md) for PR guidelines and the pull request template.
