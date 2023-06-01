![doc/img/logo](doc/img/logo.png)

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/puzzledb-go) [![Go](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/puzzledb-go.svg)](https://pkg.go.dev/github.com/cybergarage/puzzledb-go) [![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/puzzledb-go) [![codecov](https://codecov.io/gh/cybergarage/puzzledb-go/branch/main/graph/badge.svg?token=C3Q82XPE44)](https://codecov.io/gh/cybergarage/puzzledb-go)

PuzzleDB aims to be a high-performance, distributed, cloud-native, multi-API, multi-model database. This Technology Preview version is developed in the Go language.

# What is PuzzleDB

PuzzleDB is a multi-data model database that handles key-value model, relational model, and document model. In addition, PuzzleDB is a multi-API database and is compatible with existing database protocols such as MySQL, Redis, and MongoDB.

![concept](doc/img/concept.png)

PuzzleDB supports existing query protocols such as MongoDB, Redia, and MySQL. Thus, developers can start using PuzzleDB as a scalable, high-performance distributed database with existing database client drivers without any learning curve.

The name PuzzleDB comes from the ability to combine multiple modules such as coordinators, storages, and existing database protocol handlers to form a database.

## Features

PuzzleDB has the following features:

-   Flexibility: PuzzleDB allows extensibility through its plugin architecture and pluggable modules for queries, data models, storage, and more.

-   Scalability: PuzzleDB seamlessly transitions from an in-memory standalone storage plugin module to a scalable, shared-nothing, horizontally distributed database using an ordered distributed key-value store plugin module.

-   Facility: PuzzleDB supports major database model and protocol plugin modules, such as Redis, MongoDB, and MySQL, simplifying application migration.

-   Safety: PuzzleDB offers ACID-compliant plugin modules, enabling the development of intuitive and secure applications.

-   Efficiency: PuzzleDB manages various database data models, including key-value, document, and relational, by consolidating them into a single core model.

# Get Started

See the following guide to learn about how to get started.

-   [Quick Start](doc/quick-start.md)

# How does PuzzleDB work?

For architecture of PuzzleDB, see the following concept documents:

-   [Design Concepts](doc/concept.md)

    -   [Layer Concept](doc/layer-concept.md)

        -   [Plug-In Concept](doc/plugin-concept.md)

    -   [Coordinator Concept](doc/coordinator-concept.md)

    -   [Storage Concept](doc/storage-concept.md)

    -   [Data Model](doc/data-model.md)

    -   [Consistency Model](doc/consistency-model.md)

# Supported Protocols

This technology preview version partially supports major database models and protocols. Please refer to the following documents for details on support status and limitations.

-   Compatibility

    -   [Redis](doc/redis.md)

    -   [MongoDB](doc/mongodb.md)

    -   [MySQL](doc/mysql.md)

    -   PostgreSQL (Planning)

# Roadmap

PuzzleDB is currently in a technical preview release. Currently, PuzzleDB is in the process of developing and testing a distributed plugin for the next release. The development roadmap for PuzzleDB is shown below.

![doc/img/roadmap](doc/img/roadmap.png)

For more information about the roadmap, please refer to [Roadmap](doc/roadmap.adoc).

# User Guides

-   Get Started

    -   [Quick Start](doc/quick-start.md)

        -   [puzzledb-server](doc/cmd/server/puzzledb-server.md)

    -   [Configuring PullzeDB](doc/configuring.md)

-   Operation

    -   [CLI (puzzledb-cli)](doc/cmd/cli/puzzledb-cli.md)

    -   [Distributed Tracing](doc/tracing.md)

-   Benchmarking

    -   [puzzledb-bench](https://github.com/cybergarage/puzzledb-go/puzzledb-bench)

-   Distribution

    -   [Docker Hub (cybergarage/puzzledb)](https://hub.docker.com/repository/docker/cybergarage/puzzledb/general)

# Developer Guides

-   References

    -   [Go Reference](https://pkg.go.dev/github.com/cybergarage/puzzledb-go)

-   Contributing (Planning)

    -   [Coding Guidelines](doc/coding_guideline.md)

-   Building and Testing

    -   [Build on macOS](doc/build-on-macos.md)

    -   [Build on Ubuntu](doc/build-on-macos.md)

-   Extending PuzzleDB

    -   [Plug-In Concept](doc/plugin-concept.md)

    -   [Plug-in Services](doc/plugin-types.md)

        -   [Building User Plug-ins](doc/plugin-tutorial.md)

-   Specification

    -   [Coordinator Specification](doc/spec/coordinator-spec.md)

        -   [Coordinator Messaging Specification](doc/spec/coordinator-msg-spec.adoc)

        -   [Coordinator Key-Value Store Specification](doc/spec/coordinator-spec.md)

    -   Store Specification

        -   [Store Key-Value Specification](doc/spec/store-kv-spec.md)

    -   Transversed Specifications

        -   [Key-Value Store Specification](doc/spec/kv-store-spec.md)
