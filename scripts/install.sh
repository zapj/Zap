#!/bin/sh

cp -Rf zapd.service /etc/systemd/system/
systemctl enable zapd.service
systemctl start zapd.service
systemctl status zapd.service

