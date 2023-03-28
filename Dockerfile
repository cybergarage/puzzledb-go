FROM golang:1.20-alpine

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN go build -v -gcflags= github.com/cybergarage/puzzledb-go/bin/puzzledb-server

COPY ./conf/puzzledb.yaml /puzzledb.yaml
COPY ./docker/entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]