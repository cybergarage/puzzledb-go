# Roadmap

PuzzleDB is currently in a technical preview release. Currently, PuzzleDB is in the process of developing and testing a distributed plugin for the next release. The development roadmap for PuzzleDB is shown below.

![doc/img/roadmap](img/roadmap.png)

## v2.0.0

-   New features

    -   ❏ GUI Console

-   New plug-ins

    -   Enterprise plug-ins

        -   ❏ Added QoS plug-ins

    -   Security plug-ins

        -   ❏ Added Authenticator plug-ins﻿

        -   ❏ Added Audit plug-ins

        -   ❏ Added Encrypt plug-ins

## v1.x.x

-   New plug-ins

    -   Query plug-ins

        -   ❏ Added PostgreSQL plug-in

    -   Storage plug-ins

        -   ❏ Added TiKV plug-in

    -   Metrics plug-ins

        -   ❏ Added Graphite plug-in

    -   Distributed Tracer plug-ins

        -   ❏ Added OpenTracing plug-in

-   Kubernetes features

    -   ❏ Operator

## v1.0.0 (2023-06-xx)

-   New features

    -   ❏ Added TLS with mTLS Support

-   New plug-ins

    -   Document store plug-in

        -   Key-Value Store plug-ins

            -   ❏ Key-Value cache plug-in

    -   Coordinator plug-ins

        -   ❏ Added etcd plug-in

    -   Storage plug-ins

        -   Key-Value Store plug-ins

            -   ✓ Added FoundationDB plug-in

            -   ❏ Added cache store plug-in

-   Update plug-ins

    -   Query plug-ins

        -   MySQL plug-in

            -   Supported queries

                -   ❏ ALTER TABLE

                -   ❏ CREATE INDEX

## v0.9.0 (2023-05-07)

-   New features

    -   CLI Utilities

        -   ✓ Added [puzzledb-cli](cmd/cli/puzzledb-cli.md)

    -   Operator APIs

        -   ✓ Added gRPC services for operator APIs and CLI utilities.

        -   ✓ Added Prometheus metrics expoter

    -   Configuration support

        -   ✓ Added support for configuration with environment variables.

        -   ✓ Added support for configuration with puzzledb.yaml.

-   New plug-ins

    -   Coordinator plug-ins

        -   ✓ Added memdb plug-in

    -   Distributed tracer plug-ins

        -   ✓ Added OpenTelemetry plug-in

-   Update plug-ins

    -   Coder plug-ins

        -   Key coder plug-ins

            -   Tuple plug-in

                -   Fix encoder not to panic on Ubuntu 20.04

    -   Query plug-ins

        -   ✓ MySQL plug-in

            -   Supported queries

                -   ✓ DROP DATABASE

                -   ✓ DROP TABLE

## v0.8.0 (2023-04-10)

-   Initial public release

-   Initial release plug-ins

    -   Query plug-ins

        -   ✓ MySQL plug-in

        -   ✓ Redis plug-in

        -   ✓ MongoDB plug-in

    -   Storage plug-ins

        -   Document store plug-in

            -   ✓ Key-Value store plug-in

        -   Key-Value Store plug-ins

            -   ✓ memdb plug-in

    -   Coder plug-ins

        -   Document coder plug-ins

            -   ✓ CBOR coder plug-in

        -   Key coder plug-ins

            -   ✓ Tuple plug-in
