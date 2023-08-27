FROM ubuntu:23.04

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN apt-get update && \
    apt-get install -y golang wget adduser && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN wget --directory-prefix=/tmp https://github.com/apple/foundationdb/releases/download/7.3.15/foundationdb-clients_7.3.15-1_amd64.deb &&  \
    apt install /tmp/foundationdb-clients_7.3.15-1_amd64.deb &&  \
    rm /tmp/*.deb

RUN go build -o /puzzledb-server github.com/cybergarage/puzzledb-go/cmd/puzzledb-server
RUN go build -o /puzzledb-cli github.com/cybergarage/puzzledb-go/cmd/puzzledb-cli

COPY ./puzzledb/conf/puzzledb.yaml /
COPY ./docker/entrypoint.sh /

ENV PUZZLEDB_TRACER_OPENTELEMETRY_ENABLED false
ENV PUZZLEDB_TRACER_OPENTELEMETRY_ENDPOINT http://host.docker.internal:14268/api/traces
ENV PUZZLEDB_TRACER_OPENTRACING_ENABLED false
ENV PUZZLEDB_TRACER_OPENTELEMETRY_ENDPOINT host.docker.internal:6831

ENTRYPOINT ["/entrypoint.sh"]
