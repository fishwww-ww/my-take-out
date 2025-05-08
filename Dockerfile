# 编译阶段：引用最小编译环境
FROM docker.xuanyuan.me/golang:1.21.0 AS builder

# 镜像默认工作目录
WORKDIR /build


# 防止多次拉取依赖
ADD go.mod .
ADD go.sum .
# 配置镜像golang的默认配置,方便拉取依赖
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 拷贝当前目录所有文件到工作目录
COPY . .

# 设置编译环境并进行编译
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64  go build -o /app/gin-server .

# 构建阶段：使用 alpine 最小构建
FROM docker.xuanyuan.me/alpine

# 设置镜像工作目录
WORKDIR /app

# 在builder阶段复制可执行的go二进制文件app/go-exporter 到/app/go_exporter中
COPY --from=builder /app/gin-server /app/gin-server
# 创建日志文件夹
RUN mkdir /app/logger

# 时区设置
ENV TZ="Asia/Shanghai"

# 开放端口
EXPOSE 8080

# 启动服务器
CMD ["/app/gin-server"]

# 构建阶段：使用中国镜像源加速
#FROM golang:1.21.0 AS builder
#
#WORKDIR /build
#
## 设置中国区Go代理 + 编译环境
#ENV GOPROXY=https://goproxy.cn,direct \
#    GOSUMDB=sum.golang.google.cn \
#    GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
## 使用阿里云Docker镜像加速（构建阶段）
#RUN echo "https://registry.cn-hangzhou.aliyuncs.com" > /etc/containerd/config.toml && \
#    mkdir -p /etc/docker && \
#    echo '{"registry-mirrors": ["https://registry.cn-hangzhou.aliyuncs.com"]}' > /etc/docker/daemon.json
#
## 先复制依赖文件（利用缓存层）
#COPY go.mod go.sum ./
#RUN go mod download
#
## 复制源码并编译
#COPY . .
#RUN go build -o /app/gin-server .
#
## 运行阶段：使用中国镜像源
#FROM registry.cn-hangzhou.aliyuncs.com/library/alpine:3.18
#
#WORKDIR /app
#
## 配置中国时区（阿里云镜像）
#RUN apk add --no-cache tzdata && \
#    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
#    echo "Asia/Shanghai" > /etc/timezone && \
#    apk del tzdata
#
## 安全设置：非root用户
#RUN adduser -D -g '' appuser && \
#    chown -R appuser:appuser /app
#
## 复制二进制文件
#COPY --from=builder --chown=appuser:appuser /app/gin-server /app/gin-server
#
## 日志目录
#RUN mkdir /app/logger && \
#    chown appuser:appuser /app/logger
#
## 环境变量
#ENV TZ="Asia/Shanghai"
#
#USER appuser
#EXPOSE 8080
#CMD ["/app/gin-server"]