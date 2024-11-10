FROM golang:alpine3.20

COPY . /puzzledb

WORKDIR /puzzledb

RUN go build -o /puzzledb-server doc/cmd/server/main.go
RUN go build -o /puzzledb-cli doc/cmd/clie/main.go

COPY ./puzzledb/conf/puzzledb.yaml /
COPY ./docker/entrypoint.sh /

ENV PUZZLEDB_LOGGER_ENABLED true
ENV PUZZLEDB_LOGGER_LEVEL info
ENV PUZZLEDB_PPROF_ENABLED false
ENV PUZZLEDB_PLUGINS_STORE_KV_DEFAULT memdb
ENV PUZZLEDB_PLUGINS_STORE_KV_MEMDB_ENABLED true
ENV PUZZLEDB_PLUGINS_STORE_KV_FDB_ENABLED false
ENV PUZZLEDB_PLUGINS_COORDINATOR_DEFAULT memdb
ENV PUZZLEDB_PLUGINS_COORDINATOR_MEMDB_ENABLED true
ENV PUZZLEDB_PLUGINS_COORDINATOR_FDB_ENABLED false
ENV PUZZLEDB_PLUGINS_COORDINATOR_ETCD_ENABLED false
ENV PUZZLEDB_PLUGINS_TRACER_ENABLED false
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENABLED true
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENDPOINT http://host.docker.internal:14268/api/traces
ENV PUZZLEDB_PLUGINS_TRACER_OPENTRACING_ENABLED false
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENDPOINT host.docker.internal:6831

ENTRYPOINT ["/entrypoint.sh"]
