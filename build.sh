#!/bin/sh

#-ldflags="-w -s -buildid="

if [ "$1" != "run" ]
then
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/zapj/zap/cmd/zapctl/cmd.BuildDate=`date +"%Y%m%d"` \
                   -X github.com/zapj/zap/cmd/zapctl/cmd.BuildVersion=`git symbolic-ref -q --short HEAD`-`git rev-parse --short HEAD` \
                   -X github.com/zapj/zap/cmd/zapctl/cmd.Version=`cat <./VERSION` \
                   -w -s -buildid="  cmd/zapctl/zapctl.go
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.BuildDate=`date +"%Y%m%d"` \
                   -X main.BuildVersion=`git symbolic-ref -q --short HEAD`-`git rev-parse --short HEAD` \
                   -X main.Version=`cat <./VERSION` \
                   -w -s -buildid="  cmd/zapd/zapd.go
else
    go build -ldflags "-X github.com/zapj/zap/cmd/zapctl/cmd.BuildDate=`date +"%Y%m%d"` \
                    -X github.com/zapj/zap/cmd/zapctl/cmd.BuildVersion=`git symbolic-ref -q --short HEAD`-`git rev-parse --short HEAD` \
                    -X github.com/zapj/zap/cmd/zapctl/cmd.Version=`cat <./VERSION` \
                    -w -s -buildid="  cmd/zapctl/zapctl.go

    go build -ldflags "-X main.BuildDate=`date +"%Y%m%d"` \
                    -X main.BuildVersion=`git symbolic-ref -q --short HEAD`-`git rev-parse --short HEAD` \
                    -X main.Version=`cat <./VERSION` \
                    -w -s -buildid="  cmd/zapd/zapd.go
fi


if [ "$1" = "run" ]
then
    sudo ./zapd server
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