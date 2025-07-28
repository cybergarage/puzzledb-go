![doc/img/logo](doc/img/logo.png)

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/puzzledb-go) [![Go](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml) [![go-Reference](https://pkg.go.dev/badge/github.com/cybergarage/puzzledb-go.svg)](https://pkg.go.dev/github.com/cybergarage/puzzledb-go) [![go-Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/puzzledb-go) [![codecov](https://codecov.io/gh/cybergarage/puzzledb-go/branch/main/graph/badge.svg?token=C3Q82XPE44)](https://codecov.io/gh/cybergarage/puzzledb-go)

PuzzleDB aspires to be a high-performance, distributed, cloud-native, multi-API, multi-model database. This Technology Preview version has been developed in the Go language.

# What is PuzzleDB

PuzzleDB is a multi-data model database capable of handling key-value, relational, and document models. Additionally, PuzzleDB is a multi-interface database, compatible with existing database protocols such as PostgreSQL, MySQL, Redis, and MongoDB.

![concept](doc/img/concept.png)

PuzzleDB is a distributed database framework supporting various data models and protocols. It is designed as a flexible, scalable, and efficient database framework suitable for various environments.

![system](doc/img/system.png)

PuzzleDB accommodates existing query protocols such as PostgreSQL, MySQL, MongoDB, and Redis within a distributed, pluggable database framework. Consequently, developers can seamlessly start using PuzzleDB as a scalable, high-performance distributed database with existing database client drivers, eliminating any learning curve.

## Key Features

PuzzleDB has the following features:

- Flexibility: PuzzleDB allows for extensibility through its plugin architecture and pluggable modules for queries, data models, storage, and more.

- Scalability: PuzzleDB seamlessly transitions from an in-memory standalone storage plugin module to a scalable, shared-nothing, horizontally distributed database using an ordered distributed key-value store plugin module.

- Facility: PuzzleDB supports major database model and protocol plugin modules, such as PostgreSQL, Redis, MongoDB, and MySQL, simplifying application migration.

- Safety: PuzzleDB offers ACID-compliant plugin modules, enabling the development of intuitive and secure applications.

- Efficiency: PuzzleDB manages various database data models, including key-value, document, and relational, by consolidating them into a single core model.

# Get Started

See the following guide to learn about how to get started.

- [Quick Start](doc/quick-start.md)

# How does PuzzleDB work?

For information on the concept and architecture of PuzzleDB, refer to the following concept documents:

- [Design Concepts](doc/concept.md)

  - [Layer Concept](doc/layer-concept.md)

  - [Data Model](doc/data-model.md)

  - [Storage Concept](doc/storage-concept.md)

  - [Consistency Model](doc/consistency-model.md)

  - [Coordinator Concept](doc/coordinator-concept.md)

  - [Authentication Concept](doc/auth-concept.md)

  - [Plug-In Concept](doc/plugin-concept.md)

# Supported Protocols

PuzzleDB supports the following protocols:

- Compatibility

  - [PostgreSQL](doc/postgresql.md)

  - [MySQL](doc/mysql.md)

  - [MongoDB](doc/mongodb.md)

  - [Redis](doc/redis.md)

# Roadmap

PuzzleDB is currently in a technical preview release stage. At present, it is in the process of developing and testing a distributed plugin for the upcoming release. The development roadmap for PuzzleDB is outlined below.

![doc/img/roadmap](doc/img/roadmap.png)

# User Guides

- Get Started

  - [Quick Start](doc/quick-start.md)

    - [puzzledb-server](doc/cmd/server/puzzledb-server.md)

  - [Configuring PullzeDB](doc/configuring.md)

- Operation

  - [CLI (puzzledb-cli)](doc/cmd/cli/puzzledb-cli.md)

  - [Distributed Tracing](doc/tracing.md)

- Benchmarking

  - [puzzledb-bench](https://github.com/cybergarage/puzzledb-bench)

- Distribution

  - [Docker Hub (cybergarage/puzzledb)](https://hub.docker.com/repository/docker/cybergarage/puzzledb/general)

# Developer Guides

- References

  - [go-reference](https://pkg.go.dev/github.com/cybergarage/puzzledb-go)

- Contributing (Planning)

  - [Coding Guidelines](doc/coding_guideline.md)

- Building and Testing

  - [Build on macOS](doc/build-on-macos.md)

  - [Build on Ubuntu](doc/build-on-macos.md)

- Extending PuzzleDB

  - [Plug-In Concept](doc/plugin-concept.md)

  - [Plug-in Services](doc/plugin-types.md)

    - [Building User Plug-ins](doc/plugin-tutorial.md)

- Specification

  - [Coordinator Specification](doc/spec/coordinator-spec.md)

    - [Coordinator Messaging Specification](doc/spec/coordinator-msg-spec.md)

    - [Coordinator Key-Value Store Specification](doc/spec/coordinator-spec.md)

  - Store Specification

    - [Store Key-Value Specification](doc/spec/store-kv-spec.md)

  - Transversed Specifications

    - [Key-Value Store Specification](doc/spec/kv-store-spec.md)

# Related Projects

PuzzleDB is developed in collaboration with the following Cybergarage projects:

- [go-postgresql](https://github.com/cybergarage/go-postgresql) ![go postgresql](https://img.shields.io/github/v/tag/cybergarage/go-postgresql)

- [go-mysql](https://github.com/cybergarage/go-mysql) ![go mysql](https://img.shields.io/github/v/tag/cybergarage/go-mysql)

- [go-redis](https://github.com/cybergarage/go-redis) ![go redis](https://img.shields.io/github/v/tag/cybergarage/go-redis)

- [go-mongo](https://github.com/cybergarage/go-mongo) ![go mongo](https://img.shields.io/github/v/tag/cybergarage/go-mongo)

- [go-cbor](https://github.com/cybergarage/go-cbor) ![go cbor](https://img.shields.io/github/v/tag/cybergarage/go-cbor)

- [go-logger](https://github.com/cybergarage/go-logger) ![go logger](https://img.shields.io/github/v/tag/cybergarage/go-logger)

- [go-safecast](https://github.com/cybergarage/go-safecast) ![go safecast](https://img.shields.io/github/v/tag/cybergarage/go-safecast)

- [go-sqlparser](https://github.com/cybergarage/go-sqlparser) ![go sqlparser](https://img.shields.io/github/v/tag/cybergarage/go-sqlparser)

- [go-tracing](https://github.com/cybergarage/go-tracing) ![go tracing](https://img.shields.io/github/v/tag/cybergarage/go-tracing)

- [go-authenticator](https://github.com/cybergarage/go-authenticator) ![go authenticator](https://img.shields.io/github/v/tag/cybergarage/go-authenticator)

- [go-sasl](https://github.com/cybergarage/go-sasl) ![go sasl](https://img.shields.io/github/v/tag/cybergarage/go-sasl)

- [go-sqltest](https://github.com/cybergarage/go-sqltest) ![go sqltest](https://img.shields.io/github/v/tag/cybergarage/go-sqltest)

- [go-pict](https://github.com/cybergarage/go-pict) ![go pict](https://img.shields.io/github/v/tag/cybergarage/go-pict)
