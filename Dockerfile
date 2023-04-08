FROM ubuntu:22.10

USER root

COPY . /puzzledb
WORKDIR /puzzledb

RUN apt-get update && apt-get install -y \
    golang wget \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN wget --directory-prefix=/tmp https://github.com/apple/foundationdb/releases/download/7.2.5/foundationdb-clients_7.2.5-1_amd64.deb &&  \
    apt install /tmp/foundationdb-clients_7.2.5-1_amd64.deb &&  \
    rm /tmp/*.deb

RUN go build -o /puzzledb-server github.com/cybergarage/puzzledb-go/bin/puzzledb-server

COPY ./conf/puzzledb.yaml /
COPY ./docker/entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]