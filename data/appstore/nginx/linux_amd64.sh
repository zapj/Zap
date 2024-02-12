#!/bin/bash
URL="https://nginx.org/download/nginx-${APP_VERSION}.tar.gz"

if [ -f "${PKG_PATH}/nginx-${APP_VERSION}.tar.gz" ];then
echo "File already exists, skipping download"
else
echo "Downloading nginx"
wget -P ${PKG_PATH} ${URL}
fi
echo "Building for linux/amd64"
echo $URL
echo $APP_ID
echo $APP_VERSION
echo $APP_NAME
echo $APP_TITLE
echo $BUILD_PATH
echo $LOG_FILE
echo $APP_PATH
echo $APPS_DIR
echo $SCRIPT_PATH
echo $PKG_PATH
echo $DATA_PATH