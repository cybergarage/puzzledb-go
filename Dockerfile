FROM golang:1.20-alpine

USER root

RUN apk add bash

COPY . /puzzledb
WORKDIR /puzzledb

RUN go build -v -gcflags= github.com/cybergarage/puzzledb-go/bin/puzzledb-server

COPY ./conf/puzzledb.yaml /
COPY ./docker/entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]