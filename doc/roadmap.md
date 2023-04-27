# Roadmap

PuzzleDB is currently in a technical preview release. Currently, PuzzleDB is in the process of developing and testing a distributed plugin for the next release. The development roadmap for PuzzleDB is shown below.

![doc/img/roadmap](img/roadmap.png)

## v2.0.0

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

## v1.0.0

-   New features

    -   ❏ Added TLS with mTLS Support

## v0.9.0

-   New features

    -   CLI Utilities

        -   ✓ Added puzzledb-cli

    -   Operator APIs

        -   Added gRPC server for operator APIs and CLI utilities.

    -   Configuration support

        -   ✓ Added support for configuration with environment variables.

        -   ✓ Added support for configuration with puzzledb.yaml.

-   New plug-ins

    -   Coordinator plug-ins

        -   ✓ Added memdb plug-in

        -   ❏ Added etcd plug-in

    -   Storage plug-ins

        -   Key-Value Store plug-ins

            -   ✓ Added FoundationDB plug-in

    -   Distributed Tracer plug-ins

        -   ✓ Added OpenTelemetry plug-in

        -   ✓ Added OpenTracing plug-in

## v0.8.0

-   Initial plug-ins

    -   Query plug-ins

        -   ✓ MySQL plug-in

        -   ✓ Redis plug-in

        -   ✓ MongoDB plug-in

    -   Storage plug-ins

        -   Document tore plug-in

            -   ✓ Key-Value storeplug-in

        -   Key-Value Store plug-ins

            -   ✓ memdb plug-in
