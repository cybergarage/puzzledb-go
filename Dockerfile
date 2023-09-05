FROM ubuntu:23.04

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN apt-get update && \
    apt-get install -y golang wget adduser libc6 libc6-dev && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN wget --directory-prefix=/tmp https://github.com/apple/foundationdb/releases/download/7.2.9/foundationdb-clients_7.2.9-1_amd64.deb &&  \
    apt install /tmp/foundationdb-clients_7.2.9-1_amd64.deb &&  \
    rm /tmp/*.deb

RUN wget --directory-prefix=/tmp https://github.com/apple/foundationdb/releases/download/7.2.9/foundationdb-server_7.2.9-1_amd64.deb &&  \
    apt install /tmp/foundationdb-server_7.2.9-1_amd64.deb &&  \
    rm /tmp/*.deb

RUN go build -o /puzzledb-server github.com/cybergarage/puzzledb-go/cmd/puzzledb-server
RUN go build -o /puzzledb-cli github.com/cybergarage/puzzledb-go/cmd/puzzledb-cli

COPY ./puzzledb/conf/puzzledb.yaml /
COPY ./docker/entrypoint.sh /

ENV PUZZLEDB_PLUGINS_STORE_KV_FDB_ENABLED false
ENV PUZZLEDB_PLUGINS_COORDINATOR_FDB_ENABLED false
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENABLED true
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENDPOINT http://host.docker.internal:14268/api/traces
ENV PUZZLEDB_PLUGINS_TRACER_OPENTRACING_ENABLED true
ENV PUZZLEDB_PLUGINS_TRACER_OPENTELEMETRY_ENDPOINT host.docker.internal:6831

ENTRYPOINT ["/entrypoint.sh"]
