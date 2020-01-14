# alpine 微型 linux 操作系统
FROM alpine:latest

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

RUN apk upgrade && apk add --no-cache ca-certificates

RUN mkdir /log

WORKDIR / 

COPY deploy /

COPY config.ini /

EXPOSE 8080

ENTRYPOINT ["./deploy serve"] 