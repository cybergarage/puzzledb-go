FROM ubuntu:24.04 AS foundationdb

ARG BUILDOS
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN apt-get update && \
    apt-get install -y curl wget 

RUN ./foundationdb.sh -a "$TARGETARCH" -o "$TARGETOS"

FROM foundationdb AS golang

RUN apt-get install -y golang adduser && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

FROM golang AS puzzledb-build

RUN go get github.com/apple/foundationdb/bindings/go@v0.0.0-20250616221319-fffa38913379
RUN go mod tidy
RUN go build -o /puzzledb-server github.com/cybergarage/puzzledb-go/cmd/puzzledb-server
RUN go build -o /puzzledb-cli github.com/cybergarage/puzzledb-go/cmd/puzzledb-cli

COPY ./puzzledb/conf/puzzledb.yaml /
COPY ./docker/entrypoint.sh /

ENV PUZZLEDB_LOGGER_ENABLED true
ENV PUZZLEDB_LOGGER_LEVEL info
ENV PUZZLEDB_PPROF_ENABLED false
ENV PUZZLEDB_PLUGINS_STORE_KV_DEFAULT fdb
ENV PUZZLEDB_PLUGINS_STORE_KV_MEMDB_ENABLED false
ENV PUZZLEDB_PLUGINS_STORE_KV_FDB_ENABLED true
ENV PUZZLEDB_PLUGINS_COORDINATOR_DEFAULT fdb
ENV PUZZLEDB_PLUGINS_COORDINATOR_MEMDB_ENABLED false
ENV PUZZLEDB_PLUGINS_COORDINATOR_FDB_ENABLED true
ENV PUZZLEDB_PLUGINS_COORDINATOR_ETCD_ENABLED false
ENV PUZZLEDB_PLUGINS_TRACER_ENABLED false
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENABLED true
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENDPOINT http://host.docker.internal:14268/api/traces
ENV PUZZLEDB_PLUGINS_TRACER_OPENTRACING_ENABLED false
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENDPOINT host.docker.internal:6831

ENTRYPOINT ["/entrypoint.sh"]