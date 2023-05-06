FROM ubuntu:23.04

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN apt-get clean && apt-get update && apt-get upgrade && \
    apt-get install -y golang wget adduser && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN wget --directory-prefix=/tmp https://github.com/apple/foundationdb/releases/download/7.2.5/foundationdb-clients_7.2.5-1_amd64.deb &&  \
    apt install /tmp/foundationdb-clients_7.2.5-1_amd64.deb &&  \
    rm /tmp/*.deb

RUN go build -o /puzzledb-server github.com/cybergarage/puzzledb-go/cmd/puzzledb-server
RUN go build -o /puzzledb-cli github.com/cybergarage/puzzledb-go/cmd/puzzledb-cli

COPY ./puzzledb/conf/puzzledb.yaml /
COPY ./docker/entrypoint.sh /

ENV PUZZLEDB_TRACER_OPENTELEMETRY_ENABLED false
ENV PUZZLEDB_TRACER_OPENTRACING_ENABLED false

ENTRYPOINT ["/entrypoint.sh"]
