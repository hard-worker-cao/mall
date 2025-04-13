# 第一阶段：构建基础环境 
FROM ubuntu:22.04 AS builder 
 
# 设置时区与环境变量 
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive \
    GOPATH=/go \
    GOROOT=/usr/local/go 
 
# 安装系统依赖 
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    wget \
    build-essential \
    libssl-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/* 
 
# 安装Go 1.23.5 
RUN wget https://dl.google.com/go/go1.23.5.linux-amd64.tar.gz  -O /tmp/go.tar.gz  && \
    tar -C /usr/local -xzf /tmp/go.tar.gz  && \
    rm /tmp/go.tar.gz  
 
# 安装MySQL 8.0.40 
RUN wget https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.40-linux-glibc2.17-x86_64-minimal.tar.xz  -O /tmp/mysql.tar.xz  && \
    tar -xvf /tmp/mysql.tar.xz  -C /usr/local/ && \
    ln -s /usr/local/mysql-8.0.40-linux-glibc2.17-x86_64-minimal /usr/local/mysql && \
    rm /tmp/mysql.tar.xz  
 
# 安装Redis 7.4.2 
RUN wget https://download.redis.io/releases/redis-7.4.2.tar.gz  -O /tmp/redis.tar.gz  && \
    tar -xzf /tmp/redis.tar.gz  -C /tmp && \
    cd /tmp/redis-7.4.2 && \
    make BUILD_TLS=yes && \
    make install && \
    rm -rf /tmp/redis*
 
# 第二阶段：生成最终镜像 
FROM ubuntu:22.04 
 
# 复制第一阶段构建结果 
COPY --from=builder /usr/local/go /usr/local/go 
COPY --from=builder /usr/local/mysql /usr/local/mysql 
COPY --from=builder /usr/local/bin/redis-* /usr/local/bin/
 
# 设置环境路径 
ENV PATH="/usr/local/go/bin:/usr/local/mysql/bin:$PATH" \
    MYSQL_HOME="/usr/local/mysql" \
    REDIS_HOME="/usr/local/bin"
 
# 初始化脚本 
COPY init.sh  /init.sh  
RUN chmod +x /init.sh  
 
# 暴露端口 
EXPOSE 3306 6379 
 
# 启动服务 
CMD ["/init.sh"] 