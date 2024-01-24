#!/usr/bin/bash

#-ldflags="-w -s -buildid="
# zapctl config set signing_key `openssl rand -hex 12`
VERSION="0.0.0"
if [ -n "$2" ]
then
    VERSION=$2
else
    VERSION=`cat VERSION`
    VerArr=($(echo "$VERSION" | tr '.' ' '))
    VERSION="${VerArr[0]}.${VerArr[1]}.`expr ${VerArr[2]} + 1`"
    echo $VERSION > VERSION
fi


if [ "$1" != "run" ]
then
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/zapj/zap/core.Version=v${VERSION}'" cmd/zapctl/zapctl.go
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/zapj/zap/core.Version=v${VERSION}'" cmd/zapd/zapd.go
elif [ "$1" != "build" ]
then
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/zapj/zap/core.Version=v${VERSION}'" cmd/zapctl/zapctl.go
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'github.com/zapj/zap/core.Version=v${VERSION}'" cmd/zapd/zapd.go
else
    go build -ldflags="-X 'github.com/zapj/zap/core.Version=v${VERSION}'" cmd/zapctl/zapctl.go
    go build -ldflags="-X 'github.com/zapj/zap/core.Version=v${VERSION}'" cmd/zapd/zapd.go
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
    # cp -Rf data dist/zap/
    cd dist/
    tar czvf "zap-release-v${VERSION}.tar.gz" zap/
    zapfile put zap/version.txt $VERSION
    zapfile put zap/latest.txt $VERSION
    zapfile upload zap/ "zap-release-v${VERSION}.tar.gz"
fi
