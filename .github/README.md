# Repository guidelines

## Pull requests

All changes must go through a **pull request** against `main`. Direct pushes to `main` are disabled.

When opening a PR, use the [pull request template](.github/PULL_REQUEST_TEMPLATE.md). It reminds you and reviewers to:

- Provide a clear **description** (what, why, how to verify).
- Confirm **tests** were added or updated for new/changed behavior.
- Ensure changes are **well tested** and follow existing test patterns.
- Run **lint** (`go vet`, `golangci-lint`) before requesting review.
- Handle **errors** explicitly and avoid unintended changes.

Reviewers will use the same checklist before approving. CI must pass (test + lint) before merge.

## Branch and release flow

- **Branch:** Create a feature/fix branch from `main`, make changes, then open a PR.
- **Merge:** After review and CI pass, the PR is merged into `main`.
- **Release:** Each merge to `main` triggers an automatic patch-version tag and GitHub Release. No manual tagging required.

## More

- [CONTRIBUTING.md](../CONTRIBUTING.md) — How to contribute, code expectations, and repo setup (visibility, branch protection).
- [README.md](../README.md) — Project overview, install, usage, and module layout.
