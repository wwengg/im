FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn


WORKDIR /go/src/server
COPY . .

RUN cat ./im.docker.yaml

RUN go mod tidy

RUN go env && go build -o im ./main.go


FROM alpine:latest
LABEL MAINTAINER="wwwwangg@163.com"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache  gettext tzdata   && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >  /etc/timezone && \
    date && \
    apk del tzdata

WORKDIR /go/src/server
COPY --from=0 /go/src/server/im ./
COPY --from=0 /go/src/server/im.docker.yaml ./im.yaml