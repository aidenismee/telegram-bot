#!/usr/bin/env bash

set -e
now=$(date +'%Y-%m-%dT%T')
version=$(git rev-parse --short HEAD)
package="github.com/nekizz/telegram-bot/pkg/server"

go build -a -ldflags "-s -w -X $package.version=$version -X $package.buildTime=$now" -o main cmd/api/main.go