FROM ubuntu:24.04 AS foundationdb

ARG BUILDOS
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETOS

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN apt-get update && \
    apt-get install -y curl wget adduser g++ build-essential

RUN ./foundationdb.sh -a "$TARGETARCH" -o "$TARGETOS"

# Install latest Go for the target OS and architecture
RUN LATEST_GO_VERSION=$(wget -qO- 'https://go.dev/VERSION?m=text' | head -n 1) && \
    wget https://go.dev/dl/${LATEST_GO_VERSION}.${TARGETOS}-${TARGETARCH}.tar.gz -O /tmp/go.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

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