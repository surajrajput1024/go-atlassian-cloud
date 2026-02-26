# Contributing to go-atlassian-cloud

Contributions are welcome. All changes must go through a pull request; direct pushes to `main` are not allowed.

## How to contribute

1. Fork the repository and clone your fork.
2. Create a branch from `main`: `git checkout -b your-feature-or-fix`.
3. Make your changes. Keep commits focused and messages clear.
4. Run locally: `go test ./...`, `go vet ./...`. Ensure `golangci-lint run ./...` passes if you use it.
5. Push your branch and open a **pull request** against `main`. Describe what changed and why.
6. Wait for CI to pass and for review/merge. Do not push directly to `main` in the upstream repo.

## Code expectations

- Follow existing style and package layout (client at root, APIs under `jira/`, types in `types/`, constants in `constants/`, util in `util/`).
- No comments in code unless they explain non-obvious behavior.
- Add or update tests for new or changed behavior.
- Validate at boundaries; keep errors explicit and do not swallow them.

## Versioning and releases

Releases are driven by merges to `main`. Each merge triggers an automatic patch-version tag and a GitHub Release. Do not create version tags manually unless you are doing a special release.
