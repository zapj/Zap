#!/bin/bash

source $ZAP_PATH/scripts/zap/bash_utils.sh

PHP_VERSION="${APP_VERSION}"
PHP_DOWNLOAD_URL="https://cn2.php.net/distributions/php-${PHP_VERSION}.tar.gz"
PHP_DOWNLOAD_NAME="php-${PHP_VERSION}.tar.gz"
PHP_DIRNAME="php-${APP_VERSION}"


PHP_INSTALL_PATH=${APPS_DIR}/php-${APP_VERSION}
cd $PKG_PATH
if [ -f "${PKG_PATH}/${PHP_DOWNLOAD_NAME}" ];then
echo "PHP File already exists, skipping download"
else
echo "Downloading PHP"
wget -c --progress=dot -e dotbytes=100k -O ${PHP_DOWNLOAD_NAME} ${PHP_DOWNLOAD_URL} 
fi

#>> ${LOG_FILE} 2>&1 

echo "unpacking PKGs"
tar -xvf ${PHP_DOWNLOAD_NAME} -C ${ZAP_DATA_PATH}/build
if [ "$?" != "0" ];then
echo "Error unpacking PHP"
exit 1
fi


cd ${ZAP_DATA_PATH}/build
echo "building PHP ${APP_VERSION}" 

cd php-${APP_VERSION}

./configure \
--prefix=${PHP_INSTALL_PATH} \
--with-config-file-path=${PHP_INSTALL_PATH}/etc \
--with-config-file-scan-dir=${PHP_INSTALL_PATH}/etc/php.d \
--disable-rpath \
--enable-sysvsem \
--enable-sysvshm \
--enable-pcntl \
--enable-fpm \
--with-openssl \
--with-zlib \
--with-zip \
--enable-soap \
--enable-sockets \
--with-curl \
--enable-gd \
--with-webp \
--with-jpeg \
--enable-mysqlnd \
--with-mysqli=mysqlnd \
--with-pdo-mysql=mysqlnd \
--with-pdo-pgsql \
--with-pgsql \
--enable-mbstring \
--enable-intl


echo "make install"
make -j `grep 'processor' /proc/cpuinfo | wc -l` && make install

if [ ! -d "${PHP_INSTALL_PATH}/etc" ];then
    mkdir -p ${PHP_INSTALL_PATH}/etc
fi

cp ${ZAP_DATA_PATH}/build/${PHP_DIRNAME}/php.ini-production ${PHP_INSTALL_PATH}/etc/php.ini
cp ${PHP_INSTALL_PATH}/etc/php-fpm.conf.default ${PHP_INSTALL_PATH}/etc/php-fpm.conf
cp ${PHP_INSTALL_PATH}/etc/php-fpm.d/www.conf.default ${PHP_INSTALL_PATH}/etc/php-fpm.d/www.conf
# update config

sed -i "s/;daemonize = yes/daemonize = yes/g" ${PHP_INSTALL_PATH}/etc/php-fpm.conf
sed -i "s/;pid =/pid =/g" ${PHP_INSTALL_PATH}/etc/php-fpm.conf
sed -i "s/listen = 127.0.0.1:9000/listen = /var/run/php-fpm/php-fpm-${PHP_VERSION}.sock/g" ${PHP_INSTALL_PATH}/etc/php-fpm.d/www.conf



echo "PHP ${APP_VERSION} installing successful"





