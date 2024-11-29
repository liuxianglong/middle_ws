#!/bin/sh

#golangci-lint -j 1 run ./... --timeout=10m
go env -w GOOS=linux
go build -p 1 -tags=pp -o bin/demo_http ./app/http/main.go
go build -p 1 -tags=pp -o bin/demo_job ./app/job/main.go


