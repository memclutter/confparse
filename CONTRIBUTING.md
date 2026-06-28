# Contributing

Thanks for your interest in improving `confparse`. This document describes how to
set up the project, the conventions the codebase follows, and how to get a change
merged.

## Prerequisites

- Go 1.18 or newer (the module targets `go 1.18`).
- That is all — `confparse` has no runtime dependencies beyond the standard
  library, and the test suite needs no external services.

## Getting started

```bash
git clone https://github.com/memclutter/confparse.git
cd confparse
go mod download
go test ./...
```

`confparse` is a library, not a binary: there is nothing to run on its own. The
quickest way to exercise a change is to add or adjust a case in `parse_test.go`.

## Development workflow

- Branch off `main` with a short-lived topic branch; keep one logical change per
  pull request.
- Run the checks below before pushing; CI must be green before review.

```bash
go build ./...                              # compiles
go test ./... -race -coverprofile=cover.out # runs the suite with the race detector
gofmt -l .                                  # must print nothing (formatting)
```

CI (`.github/workflows/go.yml`) runs `go test ./... -race` across a matrix of Go
versions and operating systems on every push and pull request, and uploads
coverage to Codecov. The suite is kept at 100% coverage — keep it there.

## Code style

- Format with `gofmt` / `goimports`; do not hand-format.
- Keep the public surface minimal: the package exposes a single function,
  `Parse(container interface{}) error`. Treat it and the existing struct-tag
  names (`name`, `value`, `usage`, `envVar`) as a stability contract.
- Wrap errors with context rather than returning bare errors where it adds
  information.
- Prefer table-driven tests; every supported type and error path has a row in
  `parse_test.go` — add one for anything you change.

## Adding a supported type

Supported field types are a contract. When adding one:

1. Add a `case *T:` to the type switch in `declareFlag` (`parse.go`), converting
   the `value` string into `T` and registering the matching `flag.TVar`.
2. Add a conversion helper if the type needs one, mirroring the existing
   `toInt` / `toUint` / `toTimeDuration` helpers (return the conversion error).
3. Add success and error-path rows to `parse_test.go`, and document the type in
   `README.md`.

## Commit messages

This project uses [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(parse): support float64 fields
fix(parse): return error on invalid uint default
docs(readme): document default precedence
```

Use `feat` / `fix` / `docs` / `refactor` / `test` / `chore` as appropriate; mark
breaking changes with a `!` or a `BREAKING CHANGE:` footer. Adding a type or tag
extends the public contract — call it out in the message.

## Pull requests

- Describe what changes and why; link any related issue.
- Keep the diff focused and the history readable.
- Update `README.md` and `CHANGELOG.md` (under `## [Unreleased]`) when your change
  affects behaviour users can see.

## Releases

Releases follow [Semantic Versioning](https://semver.org/). While the major
version is `0`, the public API may change in any release. Maintainers cut releases
by tagging `vMAJOR.MINOR.PATCH` and publishing notes built from the changelog.
</content>
