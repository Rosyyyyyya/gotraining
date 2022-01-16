# 构建本地镜像
# 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
# 将镜像推送至 docker 官方镜像仓库
# 通过 docker 命令本地启动 httpserver
# 通过 nsenter 进入容器查看 IP 配置

FROM golang
LABEL maintainer="RosyRong"

WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN go build -o /runhttp
EXPOSE 80
CMD [ "/runhttp" ]

# how to use:
# docker build --tag runhttp .
# docker run --publish 80:80 runhttp
# curl http://localhost/header
# docker inspect 14d0cca41040 | grep -i pid
# nsenter -t #pid -n ip a