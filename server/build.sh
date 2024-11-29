#!/bin/bash

##############################################################
# 用途：Stellar 构建脚本
# 作者：@DK
##############################################################

# 提交 ID 作为版本号保存
commit_id=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Go 版本
go_version=$(go version | awk '{print $3}' | sed 's/go//g' || echo "unknown")

# 构建，并设置版本号，需要使用绝对路径，否则无法替换值
# go run -ldflags "-X stellar/common.SystemVersion=${commit_id} -X stellar/common.SystemGoVersion=${go_version}" main.go start
go build -ldflags "-X stellar/common.SystemVersion=${commit_id} -X stellar/common.SystemGoVersion=${go_version}" -o stellar main.go

# 镜像运行
# docker run -p 13306:3306 --name mysql-dev -e MYSQL_ROOT_PASSWORD=123456 -d mysql
# docker run -p 16379:6379 --name redis-dev -e REDIS_PASSWORD=123456 -d redis
