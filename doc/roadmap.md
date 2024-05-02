# Roadmap

PuzzleDB is currently in a technical preview release. Currently, PuzzleDB is in the process of developing and testing a distributed plugin for the next release. The development roadmap for PuzzleDB is shown below.

![doc/img/roadmap](img/roadmap.png)

## v2.0.0

-   New features

    -   ❏ GUI Console

    -   ❏ Added TLS with mTLS Support

-   New plug-ins

    -   Enterprise plug-ins

        -   ❏ Added QoS plug-ins

    -   Security plug-ins

        -   ❏ Added Audit plug-ins

        -   ❏ Added Encrypt plug-ins

## v1.4.x

-   New features

    -   Kubernetes features

        -   ❏ Operator

    -   New plug-ins

        -   Storage plug-ins

            -   ❏ Added JunoDB plug-in

            -   ❏ Added TiKV plug-in

        -   Coordinator plug-ins

            -   ❏ Added etcd plug-in

        -   Metrics plug-ins

            -   ❏ Added Graphite plug-in

-   Improvements

    -   Plug-in improvements

        -   Query plug-ins

            -   MySQL plug-in

                -   Supported queries

                    -   ❏ ALTER TABLE

                    -   ❏ CREATE INDEX

                    -   ❏ LIMIT and ORDER BY in SELECT queries

            -   MongoDB plug-in

                -   Supported queries

                    -   ❏ createIndex (Only single field and Comound indexes)

                -   Disabled features

                    -   ❏ Auto Indexing

## v1.3.0 (2023-12-30)

-   New features

    -   Security plug-ins

        -   ✓ Added authenticator plug-ins

            -   ✓ Added clear text password plug-ins

                -   ✓ Supported PostgreSQL plug-in

                -   ❏ Supported MySQL plug-in

                -   ✓ Supported Redis plug-in

                -   ❏ Supported MongoDB plug-in

-   Improvements

    -   Query plug-ins

        -   Redis plug-in

            -   Supported queries

                -   ✓ HASH commands

    -   Updated storage format

## v1.2.0 (2023-11-15)

-   New features

    -   Distributed plug-ins

        -   ✓ Cache Storage plug-ins

            -   Enabled cache store plug-in (ristretto) as default

-   Updates

    -   Redis plug-in

        -   Support new commands

            -   ✓ DEL and EXISTS

-   Improvements

    -   Updated to set service metrics to prometheus

        -   Query plug-ins

            -   PostgreSQL, MySQL, Redis and Mongo

        -   Storage plug-ins

            -   Cache store (ristretto)

-   Fixed

    -   PostgreSQL plug-in

        -   ✓ Fixed transaction hangup using copy commands

        -   ✓ Fixed to run pgbench on Ubuntu platforms

## v1.1.0 (2023-10-20)

-   New features

    -   Query plug-ins

        -   MySQL plug-in

            -   Support transaction control statements

                -   ✓ BEGIN, COMMIT and ROLLBACK

        -   PostgreSQL plug-in

            -   Support transaction control statements

                -   ✓ BEGIN, COMMIT and ROLLBACK

    -   ✓ Enable pprof

-   Improvements

    -   PostgreSQL plug-in

        -   Supported basic aggregate functions

            -   ✓ COUNT, SUM, AVG, MIN and MAX

        -   Supported basic math functions

            -   ✓ ABS, CEIL and FLOOR

        -   Supported new data types

            -   ✓ TIMESTAMP

        -   Improved schema validation for INSERT, SELECT, UPDATE, and DELETE queries

        -   ✓ Enabled PICT based scenario tests of go-sqltest

        -   Supported new statements

            -   ✓ TRUNCATE, VACCUM and COPY

        -   Supported pgbench workload

    -   MySQL plug-in

        -   Supported new data types

            -   ✓ DATETIME and TIMESTAMP

        -   Improved schema validation for INSERT, SELECT, UPDATE, and DELETE queries

        -   Enabled PICT based scenario tests of go-sqltest

## v1.0.0 (2023-08-30)

-   New plug-ins

    -   Query plug-ins

        -   ✓ Added PostgreSQL plug-in

    -   Coordinator plug-ins

        -   ✓ Added FoundationDB plug-in

    -   Storage plug-ins

        -   Key-Value Store plug-ins

            -   ✓ Added FoundationDB plug-in

            -   ✓ Added cache store plug-in

-   Improvements

    -   CLI Utilities

        -   ✓ Added status commands to [puzzledb-cli](cmd/cli/puzzledb-cli.md)

    -   Storage plug-ins

        -   Key-Value Store plug-ins

            -   ✓ Update store interface to Support limit and order options in Range queries

            -   memdb plug-in

                -   ✓ Support limit and order options in Range queries

            -   FoundationDB plug-in

                -   ✓ Support limit and order options in Range queries

        -   Document store plug-in

            -   ✓ Support limit and order options in Range queries based on key-value Store plug-ins

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

        -   ✓ Added OpenTracing plug-in

-   Plug-in improvements

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
