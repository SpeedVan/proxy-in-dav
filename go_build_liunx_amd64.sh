#!/bin/sh

target=$1

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $target ./main.go
