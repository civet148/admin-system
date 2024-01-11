#!/bin/bash

#镜像地址和版本(本地)
IMAGE_URL='admin-system:latest'

#删除原来的容器
docker rm -f admin-system

# MySQL数据源(正式环境需修改成实际数据库配置)
DSN='root:123456@tcp(192.168.1.16:3306)/admin-system?charset=utf8'

# 管理系统HTTP服务监听地址
LISTEN_ADDR="0.0.0.0:8088"

# 域名配置
DOMAIN=https://admin.your-enterprise.com


docker run --net=host -p 8080:8088 --restart always --name admin-system -d $IMAGE_URL \
       admin-system run --debug --domain "${DOMAIN}" --dsn "${DSN}" --static "${FRONT_END}" $LISTEN_ADDR


