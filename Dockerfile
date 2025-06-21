FROM ubuntu:24.04 AS foundationdb

ARG BUILDOS
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN apt-get update && \
    apt-get install -y curl wget adduser build-essential

RUN ./scripts/fdb_install.sh -a "$TARGETARCH" -o "$TARGETOS"

RUN LATEST_GO_VER=$(wget -qO- 'https://go.dev/VERSION?m=text' | head -n 1) && \
    LATEST_GO_PKG=${LATEST_GO_VER}.${TARGETOS}-${TARGETARCH}.tar.gz && \ 
    wget https://go.dev/dl/${LATEST_GO_PKG} -O ${LATEST_GO_PKG} && \
    tar -C /usr/local -xzf ${LATEST_GO_PKG} && \
    rm ${LATEST_GO_PKG}
ENV PATH="/usr/local/go/bin:${PATH}"

RUN CGO_ENABLED=1 go build -o /puzzledb-server github.com/cybergarage/puzzledb-go/cmd/puzzledb-server
RUN CGO_ENABLED=1 go build -o /puzzledb-cli github.com/cybergarage/puzzledb-go/cmd/puzzledb-cli

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