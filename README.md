# Confparse

[![Language Golang](https://img.shields.io/badge/language-golang-blue.svg)](https://img.shields.io/badge/language-golang-blue.svg)
[![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)](https://github.com/memclutter/confparse)
[![Build Status](https://travis-ci.com/memclutter/confparse.svg?branch=master)](https://travis-ci.com/memclutter/confparse)

Parser of configuration files for golang projects.

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
}

var appConfig = &Config{}

func main() {
	if err := confparse.Parse(appConfig); err != nil {
		log.Fatalf("Error parse configuration: %s", err)
	}

	log.Printf("Listen and serve on %s", appConfig.Addr)
	if err := http.ListenAndServe(appConfig.Addr, nil); err != nil {
		log.Fatalf("Listen and serve error: %s", err)
	}
}

```

