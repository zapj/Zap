#!/usr/bin/bash

#-ldflags="-w -s -buildid="
# zapctl config set signing_key `openssl rand -hex 12`
VERSION=$(cat VERSION)
BUILD_DATE=$(date +'%Y%m%d_%H:%M:%S')
ARCH=$(uname -m)
OS=$(uname)
ZAP_BASEDIR=$(pwd)

if [[ $ARCH == x86_64* ]]; then
    ARCH="amd64"
elif  [[ $arch == arm* ]] || [[ $arch = aarch64 ]]; then
    ARCH="arm64"
fi

if [ "$OS" = "Linux" ];then
    OS="linux"
fi

create_pkg(){
    rm -rf dist
    #打包
    mkdir -p dist/zap
    cp -f zapd dist/zap/
    cp -f zapctl dist/zap/
    cp -Rf scripts dist/zap/
    cp -Rf conf dist/zap/
    # cp -Rf data dist/zap/
    mkdir dist/zap/data
    mkdir dist/zap/data/logs
    cp -Rf data/appstore dist/zap/data/appstore
    cd dist/
    tar czvf "zap-v${VERSION}-${OS}-${ARCH}.tar.gz" zap/
}



upload_pkg(){
    # zapfile put zap/version.txt $VERSION
    echo -n "linux_install.sh > "
    zapfile upload zap/dist/linux_install.sh zap/scripts/install.sh
    echo -n "latest.txt > "
    zapfile put zap/dist/latest.txt $VERSION
    echo -n "release file : zap-v${VERSION}-${OS}-${ARCH}.tar.gz > "
    zapfile upload zap/dist/ "zap-v${VERSION}-${OS}-${ARCH}.tar.gz"
    if [ $? -ne 0 ];then
        echo "v${VERSION} 发布失败"
        exit 1
    fi
    echo "v${VERSION} 发布成功"
}



LD_FLAGS_STRING="-w -s -X 'main.Version=v${VERSION}' -X 'main.BuildDate=${BUILD_DATE}'"

if [ "$1" = "pkg" ]
then
    if [ -n "$2" ]
    then
        VERSION=$2
        echo $VERSION > VERSION
    else
        VERSION=$(cat VERSION)
        VerArr=($(echo "$VERSION" | tr '.' ' '))
        VERSION="${VerArr[0]}.${VerArr[1]}.$(expr ${VerArr[2]} + 1)"
        echo $VERSION > VERSION
    fi
    rm -rf ./zapd
    rm -rf ./zapctl
    if [ -z "$SKIP_WEB" ];then 
        cd web
        npm run build
        cd ../
    fi
    LD_FLAGS_STRING="-w -s -X 'main.Version=v${VERSION}' -X 'main.BuildDate=${BUILD_DATE}'"
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="${LD_FLAGS_STRING}" -trimpath cmd/zapctl/zapctl.go
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="${LD_FLAGS_STRING}" -trimpath cmd/zapd/zapd.go
elif [ "$1" = "build" ]
then
    rm -rf ./zapd
    rm -rf ./zapctl
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="${LD_FLAGS_STRING}" -trimpath cmd/zapctl/zapctl.go
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="${LD_FLAGS_STRING}" -trimpath cmd/zapd/zapd.go
else
    echo "Current Version: ${VERSION}"
    go build -ldflags="${LD_FLAGS_STRING}" cmd/zapctl/zapctl.go
    go build -ldflags="${LD_FLAGS_STRING}" cmd/zapd/zapd.go
fi
echo "Build Success"

if [ "$1" = "run" ]
then
    sudo ZAP_MODE=DEV ./zapd master
elif [ "$1" = "build" ]
then
    create_pkg
elif [ "$1" = "pkg" ]
then
    
    create_pkg
    upload_pkg
    
fi


