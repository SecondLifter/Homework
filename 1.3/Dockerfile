FROM golang:1.16-alpine AS builder

ENV GO111MODULE on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN set -ex \
    && env GOOS=linux GOARCH=amd64 go  build -o httpdserver main.go

FROM alpine

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV LANGUAGE en_US:en

EXPOSE 8080
RUN set -ex \
    && apk add bash tzdata curl ca-certificates \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && rm -rf /var/cache/apk/*

COPY --from=builder /app/httpdserver /httpdserver

ENTRYPOINT ["/httpdserver"]

CMD ["--help"]
