# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Root package `github.com/surajrajput1024/go-atlassian-cloud` now re-exports `ErrNotFound`, `ErrUnauthorized`, `ErrForbidden`, `ErrBadRequest` so callers can use `errors.Is(err, atlassian.ErrNotFound)` without importing `client`.

### Changed

- API error wrapping now keeps `*APIError` in the error chain: both `errors.Is(err, client.ErrNotFound)` and `errors.As(err, &apiErr)` work on errors returned from `DoJSON` and other client methods.

### Fixed

- `APIError.Error()` no longer relies on non-deterministic map iteration when formatting `Errors`; first key is chosen after sorting for stable output.

### Security

- (List security-related changes here.)

---

## [0.1.8] and earlier

- Jira client: projects, permission schemes and grants, project permission scheme, project roles and actors, groups, workflow scheme project associations.
- Config validation, retries with backoff, typed API errors (`APIError`, `ErrNotFound`, `ErrUnauthorized`, etc.).
- Context-aware API methods where applicable (`GetProjectWithContext`, `CreateProjectWithContext`, etc.).

[Unreleased]: https://github.com/surajrajput1024/go-atlassian-cloud/compare/v0.1.8...HEAD
[0.1.8]: https://github.com/surajrajput1024/go-atlassian-cloud/releases/tag/v0.1.8
