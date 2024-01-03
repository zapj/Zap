#!/bin/sh

go build assist/zap_assist.go
go build server/zap_server.go
go build cmd/zap_cli.go