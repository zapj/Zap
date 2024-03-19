#!/bin/bash

if [ -d "/usr/local/mysql" ];then
    echo "mysql already installed"
    exit 1
fi

source $ZAP_PATH/scripts/zap/bash_utils.sh

if [ groups mysql >/dev/null 2>&1 ];then
    echo "mysql group exists"
else
    echo "create mysql group"
    groupadd mysql
    useradd -r -g mysql -s /bin/false mysql
fi

if [ ! -d "/var/log/mariadb" ];then
    mkdir -p /var/log/mariadb
fi
chown mysql:mysql -R /var/log/mariadb

cd $PKG_PATH
if [ ! -f "mariadb-${APP_VERSION}-linux-systemd-x86_64.tar.gz" ];then
    wget -O mariadb-${APP_VERSION}-linux-systemd-x86_64.tar.gz https://mirrors.aliyun.com/mariadb/mariadb-${APP_VERSION}/bintar-linux-systemd-x86_64/mariadb-${APP_VERSION}-linux-systemd-x86_64.tar.gz
fi


if [ ! -f "mariadb-${APP_VERSION}-linux-systemd-x86_64.tar.gz" ];then
    echo "Error download mariadb-${APP_VERSION}-linux-systemd-x86_64.tar.gz"
    exit 1
fi

tar xf mariadb-${APP_VERSION}-linux-systemd-x86_64.tar.gz -C "$APPS_DIR"

cd $APPS_DIR

INSTALL_DIR="${APPS_DIR}/mariadb-${APP_VERSION}"

mv mariadb-${APP_VERSION}-linux-systemd-x86_64 mariadb-${APP_VERSION}

cd mariadb-${APP_VERSION}
if [ ! -d "/usr/local/mysql" ];then
    ln -sf ${INSTALL_DIR} /usr/local/mysql
fi
cd /usr/local/mysql

cp -Rf $ZAP_PATH/scripts/zap/conf/mariadb.cnf /etc/mysql/my.cnf

echo "init mariadb script"
scripts/mariadb-install-db --user=mysql
if [ $? -ne 0 ];then
    echo "Error init mariadb"
    exit 1
fi

MYSQL_ROOT_PASSWORD=$(openssl rand -hex 8)
if [ -z "${MYSQL_ROOT_PASSWORD}" ];then
    MYSQL_ROOT_PASSWORD=$(< /dev/urandom tr -dc A-Za-z0-9_ | head -c9)
fi

echo "set mysql root password"
bin/mysql -u root -e "ALTER USER 'root'@'localhost' IDENTIFIED BY '${MYSQL_ROOT_PASSWORD}';"
echo "create mysql zapadm user"
bin/mysql -u root -e "GRANT ALL PRIVILEGES ON *.* TO 'zapadm'@'localhost' IDENTIFIED BY '${MYSQL_ROOT_PASSWORD}';"
bin/mysql -u root -e "FLUSH PRIVILEGES;"

wzap_conf mariadb_u "root"
wzap_conf mariadb_p "${MYSQL_ROOT_PASSWORD}"

#write mysql root
${ZAPCTL} config set mariadb_u "zapadm"
${ZAPCTL} config set mariadb_p "${MYSQL_ROOT_PASSWORD}"

cp -Rf support-files/mysql.server /etc/init.d/mariadb
chmod +x /etc/init.d/mariadb
cp -Rf support-files/systemd/mariadb.service /usr/lib/systemd/system/mariadb.service

systemctl daemon-reload

if command -v systemctl >/dev/null 2>&1;then
    systemctl enable mariadb.service
    systemctl start mariadb.service
elif command -v chkconfig >/dev/null 2>&1;then
    chkconfig --add mariadb
    chkconfig mariadb on
fi

COLS_DATA="install_dir=${INSTALL_DIR},\
expose=unix=/tmp/mysql.sock\ntcp=127.0.0.1\:3306,\
status=active,\
app_status=stoped,\
instance=mariadb${APP_VERSION},\
pid_file=${INSTALL_DIR}/data/mysql.pid,\
config_file=/etc/mysql/my.cnf"


echo "update zap apps"
${ZAPCTL} table apps -d "${COLS_DATA}" -w "id=${APP_ID}"

echo "mariadb install successful"