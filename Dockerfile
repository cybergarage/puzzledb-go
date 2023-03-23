USER root

COPY . /puzzledb
WORKDIR /puzzledb

FROM ubuntu:22.04

RUN make install
