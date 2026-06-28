# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
While the major version is `0`, the public API may change in any release.

## [Unreleased]

## [0.0.4] - 2026-06-28

### Added

- `CONTRIBUTING.md` describing setup, the test/lint workflow, code style, and the
  commit/PR/release process.
- This `CHANGELOG.md`.

### Fixed

- CI is green again. The GitHub Actions workflow tested Go 1.14–1.17, which have
  no `darwin/arm64` builds and so failed to install on the now-Apple-Silicon
  `macOS-latest` runners, cancelling the whole matrix. The workflow now tests the
  module's real floor (`1.18`) plus the two latest Go releases, sets
  `fail-fast: false`, and upgrades to `actions/checkout@v4`,
  `actions/setup-go@v5`, and `codecov/codecov-action@v5`.

### Changed

- Relicensed the project from the Apache License 2.0 to the MIT License.
- Reworked `README.md`: status badges, a tag reference table, the supported-type
  list, an explicit default-precedence rule (CLI flag > `envVar` > `value` > zero
  value), and a runnable example.

## [0.0.3] - 2022-11-08

### Added

- `uint` and `uint64` field types.

### Changed

- Converted the project to a Go module (`module github.com/memclutter/confparse`,
  `go 1.18`).
- Replaced Travis CI with a GitHub Actions workflow running `go test ./... -race`
  across a Go version matrix (1.14–1.17) on Linux, macOS, and Windows, and
  uploading coverage to Codecov.
- Reworked the test suite into table-driven cases at 100% coverage, including the
  default-value error paths for every supported type.

## [0.0.2] - 2018-10-08

### Added

- `int64` field type.

## [0.0.1] - 2018-09-25

First release of `confparse` — a declarative command-line argument parser for Go.

### Added

- `Parse(container interface{}) error`, which reflects over a pointer-to-struct
  and registers each field as a standard-library `flag`.
- Struct tags `name` (flag name), `value` (string default), `usage` (help text),
  and `envVar` (environment-variable fallback for the default).
- Supported field types: `string`, `int`, `bool`, and `time.Duration`.
- Environment extension: when a field's `envVar` variable is set and non-empty,
  its value becomes the field's default.

[Unreleased]: https://github.com/memclutter/confparse/compare/v0.0.4...HEAD
[0.0.4]: https://github.com/memclutter/confparse/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/memclutter/confparse/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/memclutter/confparse/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/memclutter/confparse/releases/tag/v0.0.1
</content>
