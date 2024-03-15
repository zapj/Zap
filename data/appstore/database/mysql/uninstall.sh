#!/bin/bash

if [ -d "$APP_PATH/data" ];then
    mkdir -p /root/zap_bak/mysql
    cp -Rf $APP_PATH/data /root/zap_bak/mysql
    cp -Rf /etc/mysql /root/zap_bak/mysql/etc_mysql.cnf
fi

if [ -L "/usr/local/mysql" ];then
    rm -rf /usr/local/mysql
fi

rm -rf /usr/local/bin/mysql
rm -rf /usr/local/bin/mysqldump
rm -rf /usr/local/bin/myisamchk
rm -rf /usr/local/bin/mysqld_safe
rm -rf /usr/local/bin/mysqlcheck

if [ -d "$APP_PATH" ];then
    echo "Removing $APP_NAME..."
    rm -rf $APP_PATH
    echo "Removing $APP_NAME done."
fi