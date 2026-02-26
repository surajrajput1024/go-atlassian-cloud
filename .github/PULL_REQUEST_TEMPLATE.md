## Description

<!-- Clearly describe what this PR does. Include motivation, scope, and any breaking or notable behavior. -->

## Type of change

- [ ] Bug fix
- [ ] New feature or API
- [ ] Refactor (no behavior change)
- [ ] Docs or tooling only

## Checklist for reviewers

- [ ] **Description** — PR has enough context for a reviewer (what changed, why, and how to verify).
- [ ] **Tests** — New or changed behavior is covered by tests.
- [ ] **Test quality** — Tests are focused, stable, and use the existing patterns (e.g. `httptest` for HTTP).
- [ ] **Lint** — `go vet ./...` and `golangci-lint run ./...` pass locally.
- [ ] **No unintended changes** — Only files relevant to this change are included.
- [ ] **Errors** — Error returns are handled explicitly; no silent ignores unless intentional and documented.

## How to verify

<!-- Steps or commands a reviewer can use to verify the change (e.g. run a test, call an endpoint). -->

## Additional notes

<!-- Optional: links, follow-up work, or anything else reviewers should know. -->
