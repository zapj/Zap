#!/bin/sh



go build -ldflags "-X github.com/zapj/zap/cmd/zapctl/cmd.BuildDate=`date +"%Y%m%d"` \
                   -X github.com/zapj/zap/cmd/zapctl/cmd.BuildVersion=`git symbolic-ref -q --short HEAD`-`git rev-parse --short HEAD` \
                   -X github.com/zapj/zap/cmd/zapctl/cmd.Version=`cat <./VERSION`"  cmd/zapctl/zapctl.go

go build -ldflags "-X main.BuildDate=`date +"%Y%m%d"` \
                   -X main.BuildVersion=`git symbolic-ref -q --short HEAD`-`git rev-parse --short HEAD` \
                   -X main.Version=`cat <./VERSION`"  cmd/deamon/zap_deamon.go

if [ "$1" = "run" ];then
    sudo ./zap_deamon server
fi

