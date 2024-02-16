#!/bin/bash
set -e
set -o errexit
set -o pipefail

ZLIB_VERSION="1.3.1"
ZLIB_PKG_URL=https://www.zlib.net/zlib-${ZLIB_VERSION}.tar.gz
ZLIB_PKG_NAME=zlib-${ZLIB_VERSION}.tar.gz
ZLIB_DIRNAME=zlib-${ZLIB_VERSION}


NGINX_PKG_URL="https://nginx.org/download/nginx-${APP_VERSION}.tar.gz"
NGINX_PKG_NAME=nginx-${APP_VERSION}.tar.gz
NGINX_DIRNAME=nginx-${APP_VERSION}


PCRE2_VERSION="10.37"
PCRE2_PKG_URL=https://sourceforge.net/projects/pcre/files/pcre2/${PCRE2_VERSION}/pcre2-${PCRE2_VERSION}.tar.gz/download
PCRE2_PKG_NAME=pcre2-${PCRE2_VERSION}.tar.gz
PCRE2_DIRNAME=pcre2-${PCRE2_VERSION}

echo $PKG_PATH

INSTALL_PATH=${APPS_DIR}/nginx-${APP_VERSION}
if [ -f "${PKG_PATH}/${NGINX_PKG_NAME}" ];then
echo "Nginx File already exists, skipping download"
else
echo "Downloading nginx"
wget -nv -P ${PKG_PATH} -O ${NGINX_PKG_NAME} ${NGINX_PKG_URL}
fi

if [ -f "${PKG_PATH}/${ZLIB_PKG_NAME}" ];then
echo "zlib File already exists, skipping download"
else
echo "Downloading zlib"
wget -nv -P ${PKG_PATH} -O ${ZLIB_PKG_NAME} ${ZLIB_PKG_URL}
fi

if [ -f "${PKG_PATH}/${PCRE2_PKG_NAME}" ];then
echo "pcre2 File already exists, skipping download"
else
echo "Downloading pcre2"
wget -nv -P ${PKG_PATH} -O ${PCRE2_PKG_NAME} ${PCRE2_PKG_URL}
fi



echo "unpacking PKGs"
cd ${PKG_PATH}
tar -xvf ${ZLIB_PKG_NAME}
tar -xvf ${NGINX_PKG_NAME}
tar -xvf ${PCRE2_PKG_NAME}
if [ "$?" != "0" ];then
echo "Error unpacking pcre2"
exit 1
fi


echo "building nginx"

cd $NGINX_DIRNAME
./configure \
--prefix=${INSTALL_PATH} \
--with-http_ssl_module \
--with-http_v2_module \
--with-pcre=${PKG_PATH}/${PCRE2_DIRNAME} \
--with-zlib=${PKG_PATH}/${ZLIB_DIRNAME}

make
echo "nginx build success"
make install

if [ "$?" != "0" ];then
echo "Error building nginx"
exit 1
fi

echo "Setting up nginx symlink"
if [ ! -d "${APPS_DIR}/nginx" ];then
    ln -s ${INSTALL_PATH} ${APPS_DIR}/nginx
fi

echo "Setting up nginx service"
if [ ! -f "/etc/systemd/system/nginx.service" ];then
    cp ${ZAP_PATH}/scripts/systemd/nginx.service /etc/systemd/system/nginx.service
fi


echo "Setting up nginx config"
CHANGE_APPS_CONFIG='{
    "install_dir": "${INSTALL_PATH}"
}'
${ZAPCTL} update apps -d ${CHANGE_APPS_CONFIG} -w "id=${APP_ID}"

echo "nginx installed to ${INSTALL_PATH}"

echo "Enabling nginx service"
systemd enable nginx.service
echo "Starting nginx service"
systemd start nginx.service
