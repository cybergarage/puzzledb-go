USER root

COPY . /puzzledb
WORKDIR /puzzledb

FROM ubuntu:22.04

RUN make build

go build -o /puzzledb-server github.com/cybergarage/puzzledb-go/bin/puzzledb-server

COPY ./docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]