#!/bin/sh

#-ldflags="-w -s -buildid="
# zapctl config set signing_key `openssl rand -hex 12`

if [ "$1" != "run" ]
then
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/zapctl/zapctl.go
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/zapd/zapd.go
elif [ "$1" != "build" ]
then
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/zapctl/zapctl.go
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/zapd/zapd.go
else
    go build cmd/zapctl/zapctl.go
    go build cmd/zapd/zapd.go
fi


if [ "$1" = "run" ]
then
    sudo ./zapd server debug
elif [ "$1" = "pkg" ]
then
    rm -rf dist/zap
    #打包
    mkdir -p dist/zap
    cp -f zapd dist/zap/
    cp -f zapctl dist/zap/
    cp -Rf scripts dist/zap/
    cp -Rf conf dist/zap/
    cp -Rf data dist/zap/
    cd dist/
    tar czvf "zap-release-`date +"%Y%m%d"`.tar.gz" zap/
fi