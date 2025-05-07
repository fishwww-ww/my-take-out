## 编译阶段：引用最小编译环境
#FROM golang:1.21.0 AS builder
#
## 镜像默认工作目录
#WORKDIR /build
#
#
## 防止多次拉取依赖
#ADD go.mod .
#ADD go.sum .
## 配置镜像golang的默认配置,方便拉取依赖
#RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN go mod download
#
## 拷贝当前目录所有文件到工作目录
#COPY . .
#
## 设置编译环境并进行编译
#RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64  go build -o /app/gin-server .
#
## 构建阶段：使用 alpine 最小构建
#FROM alpine
#
## 设置镜像工作目录
#WORKDIR /app
#
## 在builder阶段复制可执行的go二进制文件app/go-exporter 到/app/go_exporter中
#COPY --from=builder /app/gin-server /app/gin-server
## 创建日志文件夹
#RUN mkdir /app/logger
#
## 时区设置
#ENV TZ="Asia/Shanghai"
#
## 开放端口
#EXPOSE 8080
#
## 启动服务器
#CMD ["/app/gin-server"]

# 编译阶段
FROM golang:1.21.0 AS builder

WORKDIR /build

# 设置Go环境变量（替代 go env -w）
ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 先只复制依赖文件（利用缓存层）
COPY go.mod go.sum ./
RUN go mod download

# 复制源码并编译
COPY . .
RUN go build -o /app/gin-server .

# 运行阶段
FROM alpine:3.18

WORKDIR /app

# 安装时区数据
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata

# 创建非root用户
RUN adduser -D -g '' appuser && \
    chown -R appuser:appuser /app

# 从构建阶段复制二进制文件
COPY --from=builder --chown=appuser:appuser /app/gin-server /app/gin-server

# 创建日志目录（权限给非root用户）
RUN mkdir /app/logger && \
    chown appuser:appuser /app/logger

# 环境变量
ENV TZ="Asia/Shanghai"

# 切换到非root用户
USER appuser

EXPOSE 8080
CMD ["/app/gin-server"]