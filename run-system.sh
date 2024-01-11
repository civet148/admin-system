#!/bin/bash

# MySQL数据源(正式环境需修改成实际数据库配置)
DSN="root:123456@tcp(192.168.1.16:3306)/admin-system?charset=utf8"

# 集群管理系统HTTP服务监听地址
LISTEN_ADDR="0.0.0.0:8088"

# 域名配置
DOMAIN="https://admin.your-enterprise.com"

# 前端文件路径
FRONT_END=/opt/frontend/dist

./admin-system run --debug --domain "${DOMAIN}" --dsn "${DSN}" --static "${FRONT_END}" $LISTEN_ADDR

