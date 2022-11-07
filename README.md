# Confparse

[![Language Golang](https://img.shields.io/badge/language-golang-blue.svg)](https://img.shields.io/badge/language-golang-blue.svg)
[![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/memclutter/confparse)
[![codecov](https://codecov.io/gh/memclutter/confparse/branch/main/graph/badge.svg?token=G0U5MRSSFZ)](https://codecov.io/gh/memclutter/confparse)
[![Go](https://github.com/memclutter/confparse/actions/workflows/go.yml/badge.svg)](https://github.com/memclutter/confparse/actions/workflows/go.yml)

Declarative command line argument parser for golang projects. 

## Install

Install go module

```shell
go get github.com/memclutter/confparse
```

## Usage

Use `struct` tags to declare command-line arguments. 

```go
// ...

type Config struct {
	Argument1 string `name:"arg1" usage:"Argument 1 help text"`
	Timeout time.Duration `name:"timeout" value:"200ms" usage:"Timeout argument"`
}

// ...
```

## Supported Types

Different types of arguments are supported:

- `string` by default any arguments is a string
- `int` like `1`, `2`, `300`, `-23` etc
- `time.Duration` for time interval argument, like `10s`, `500ms`, `20us` etc
- `bool` for boolean argument

## Environment Extension

Use special struct tag `envVar` if you application read configuration from environment variables. 
Set environment variable name in `envVar` and confparse read value from there.

## Example

The following is an example of defining a configuration for a simple web server

```go
package main

import (
	"log"
	"net/http"

	"github.com/memclutter/confparse"
)

type Config struct {
	Addr string `name:"addr" value:":8000" usage:"Listen and serve address"`
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

