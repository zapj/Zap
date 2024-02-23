#!/bin/bash

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
if [ "$?" != "0" ];then
echo "Error download nginx"
exit 1
fi

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
INSTENCE_NAME=${NGINX_DIRNAME}
if [ ! -d "${APPS_DIR}/nginx" ];then
    ln -s ${INSTALL_PATH} ${APPS_DIR}/nginx
    INSTENCE_NAME=nginx
    echo "nginx symlink created"
fi

echo "Setting up nginx service"
if [ ! -f "/etc/systemd/system/nginx.service" ];then
    cp ${ZAP_PATH}/scripts/systemd/nginx.service /etc/systemd/system/nginx.service
fi





echo "nginx installed to ${INSTALL_PATH}"

echo "Enabling nginx service"
systemd enable nginx.service
echo "Starting nginx service"
systemd start nginx.service
sleep 1
nginx_status=$(systemctl is-active nginx.service)

echo "Setting up nginx config"

CHANGE_APPS_CONFIG="install_dir=/usr/local/apps/nginx-${APP_VERSION},\
expose=\"http=80\&https=443\",\
status=active,\
app_status=${nginx_status},\
instance=${INSTENCE_NAME},\
pid_file=${INSTALL_PATH}/logs/nginx.pid,\
config_file=${INSTALL_PATH}/conf/nginx.conf"

${ZAPCTL} table apps -d ${CHANGE_APPS_CONFIG} -w "id=${APP_ID}"
echo "nginx config updated"
if [ "$nginx_status" != "active" ];then
echo "Error start nginx.service faild"
fi
echo "nginx install successful"
