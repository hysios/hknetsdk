FROM golang:1.14.3-stretch

ENV LD_LIBRARY_PATH=/go/src/lib
ENV GOPROXY=https://goproxy.cn

WORKDIR /go/src/
