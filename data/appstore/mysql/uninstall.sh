#!/bin/bash


if [ -d "$APP_PATH" ];then
    echo "Removing $APP_NAME..."
    rm -rf $APP_PATH
    echo "Removing $APP_NAME done."
fi