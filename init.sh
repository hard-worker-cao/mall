#!/bin/bash
# 启动MySQL服务
/usr/local/mysql/bin/mysqld --initialize-insecure
/usr/local/mysql/bin/mysqld_safe &

# 启动Redis服务
redis-server --daemonize yes

# 启动Nginx服务
nginx

# 保持容器运行
tail -f /dev/null