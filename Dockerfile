USER root

COPY . /puzzledb
WORKDIR /puzzledb

FROM ubuntu:22.04

RUN make build

COPY ./docker/entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]