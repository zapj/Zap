#!/bin/bash
set -ex

URL="https://nginx.org/download/nginx-${APP_VERSION}.tar.gz"
NGINX_FILENAME=nginx-${APP_VERSION}.tar.gz
INSTAL_PATH=${APPS_DIR}/nginx-${APP_VERSION}
if [ -f "${PKG_PATH}/nginx-${APP_VERSION}.tar.gz" ];then
echo "File already exists, skipping download"
else
echo "Downloading nginx"
wget -P ${PKG_PATH} -O ${NGINX_FILENAME} ${URL}
fi
echo "Building for linux/amd64"
cd ${PKG_PATH}
tar -xvf ${NGINX_FILENAME}
cd nginx-${APP_VERSION}
./configure \
--prefix=${INSTAL_PATH} \
--user=nginx \
--group=nginx \
--with-http_ssl_module \
--with-http_v2_module \
--with-http_gzip_static_module \
--with-http_stub_status_module \
--with-http_realip_module \
--with-http_auth_request_module \
--with-http_addition_module \
--with-http_image_filter_module \
--with-http_sub_module \
--with-http_secure_link_module \
--with-http_random_index_module \
--with-http_gunzip_module \
--with-http_slice_module \
--with-http_spdy_module \
--with-http_realip_module \
--with-http_perl_module \
--with-http_xslt_module \
--with-http_mp4_module \
--with-http_geoip_module \
--with-http_slice_module \
--with-http_dav_module \
--with-http_xslt_module \
--with-http_image_filter_module \
--with-http_flv_module \
--with-http_mp4_module \
--with-http_random_index_module \
--with-http_secure_link_module \
--with-http_sub_module \
--with-http_xslt_module \
--with-http_v2_module \
--with-http_gunzip_module \
--with-http_slice_module \


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