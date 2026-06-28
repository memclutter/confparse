# confparse

> Declarative command-line argument parser for Go — describe your configuration
> as a tagged struct and call `confparse.Parse`; fields become CLI flags, with
> per-field defaults and optional environment-variable fallbacks.

[![Release](https://img.shields.io/github/v/release/memclutter/confparse?sort=semver)](https://github.com/memclutter/confparse/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/memclutter/confparse.svg)](https://pkg.go.dev/github.com/memclutter/confparse)
[![Go Report Card](https://goreportcard.com/badge/github.com/memclutter/confparse)](https://goreportcard.com/report/github.com/memclutter/confparse)
[![Go version](https://img.shields.io/github/go-mod/go-version/memclutter/confparse)](go.mod)
[![CI](https://github.com/memclutter/confparse/actions/workflows/go.yml/badge.svg)](https://github.com/memclutter/confparse/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/memclutter/confparse/branch/main/graph/badge.svg?token=G0U5MRSSFZ)](https://codecov.io/gh/memclutter/confparse)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`confparse` turns a configuration struct into command-line flags without the
boilerplate of wiring the standard `flag` package by hand. You declare each
option as a struct field with tags, call `Parse`, and read the populated struct.

- **Declarative** — one struct describes the whole CLI surface.
- **Standard `flag` under the hood** — flags, defaults, and `-help` behave exactly
  as Go developers expect.
- **Environment fallbacks** — any field can take its default from an environment
  variable via a single tag.
- **Zero runtime dependencies** — only the Go standard library.

## Contents

- [Install](#install)
- [Usage](#usage)
- [Struct tags](#struct-tags)
- [Supported types](#supported-types)
- [Defaults and precedence](#defaults-and-precedence)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Install

```shell
go get github.com/memclutter/confparse
```

## Usage

Declare command-line arguments with `struct` tags and pass a pointer to `Parse`:

```go
type Config struct {
	Addr    string        `name:"addr" value:":8000" usage:"Listen and serve address"`
	Timeout time.Duration `name:"timeout" value:"200ms" usage:"Request timeout"`
}

cfg := &Config{}
if err := confparse.Parse(cfg); err != nil {
	log.Fatalf("parse configuration: %s", err)
}
```

`Parse` reads the tags off each field, registers a flag bound to that field, and
then parses `os.Args`. It returns an error only when a field's default `value`
cannot be parsed into the field's type.

## Struct tags

| Tag      | Meaning                                                          |
|----------|------------------------------------------------------------------|
| `name`   | Flag name — `name:"addr"` registers `-addr`.                     |
| `value`  | Default value, as a string, parsed into the field's type.       |
| `usage`  | Help text shown in the standard `flag` usage output.            |
| `envVar` | Environment variable to source the default from (see below).    |

A field with no recognised tags, or of an unsupported type, is simply skipped.

## Supported types

```text
string   int   int64   uint   uint64   bool   time.Duration
```

- `string` — used as-is (the default field type).
- `int`, `int64`, `uint`, `uint64` — parsed with `strconv`; e.g. `1`, `300`, `-23`
  (signed types only).
- `bool` — `true` / `false`.
- `time.Duration` — Go duration syntax, e.g. `10s`, `500ms`, `20us`.

## Defaults and precedence

For each field the effective default is resolved before the flag is parsed, then
the CLI flag (if present) wins:

```text
CLI flag  >  environment variable (envVar)  >  value default  >  zero value
```

Set `envVar` to read a default from the environment. When that variable is set
and non-empty, its value replaces the `value` default; an empty or unset variable
is ignored. A passed `-name` flag always overrides whatever default was chosen.

## Example

A small web server configured entirely through `confparse`:

```go
package main

import (
	"log"
	"net/http"

	"github.com/memclutter/confparse"
)

type Config struct {
	Addr   string `name:"addr" value:":8000" usage:"Listen and serve address"`
	ApiKey string `name:"apiKey" envVar:"API_KEY" usage:"API key"`
}

var appConfig = &Config{}

func main() {
	if err := confparse.Parse(appConfig); err != nil {
		log.Fatalf("Error parse configuration: %s", err)
	}

	log.Printf("API Key: %s", appConfig.ApiKey)
	log.Printf("Listen and serve on %s", appConfig.Addr)
	if err := http.ListenAndServe(appConfig.Addr, nil); err != nil {
		log.Fatalf("Listen and serve error: %s", err)
	}
}
```

```shell
go run . -addr :9000          # -addr overrides the default
API_KEY=secret go run .       # apiKey comes from the environment
```

## Contributing

Contributions are welcome — see [CONTRIBUTING.md](CONTRIBUTING.md) for setup,
coding conventions, and the commit/PR process. Changes are recorded in
[CHANGELOG.md](CHANGELOG.md).

## License

Released under the [MIT License](LICENSE).
</content>
