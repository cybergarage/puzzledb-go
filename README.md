<figure>
<img src="doc/img/logo.png" alt="PuzzleDB" />
</figure>

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/puzzledb-go) [![Go](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml) [![go-Reference](https://pkg.go.dev/badge/github.com/cybergarage/puzzledb-go.svg)](https://pkg.go.dev/github.com/cybergarage/puzzledb-go) [![go-Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/puzzledb-go) [![codecov](https://codecov.io/gh/cybergarage/puzzledb-go/branch/main/graph/badge.svg?token=C3Q82XPE44)](https://codecov.io/gh/cybergarage/puzzledb-go)

PuzzleDB is a high-performance, distributed, cloud‑native, multi‑API, multi‑model database (Technology Preview) implemented in Go.

# What is PuzzleDB

PuzzleDB is a multi‑data‑model database capable of handling key‑value, relational, and document models. It is also multi‑interface, speaking existing protocols (PostgreSQL, MySQL, Redis, MongoDB).

![concept](doc/img/concept.png)

PuzzleDB is a distributed framework supporting diverse models and protocols, designed for flexibility, scalability, and efficiency.

![system](doc/img/system.png)

By accommodating existing query protocols, developers can adopt PuzzleDB with standard client drivers and minimal learning curve.

## Key Features

PuzzleDB provides:

- Extensibility – Modular plugin architecture for queries, models, storage, coordination, tracing, metrics.

- Scalability – Seamless path from local in‑memory to shared‑nothing horizontal distribution via ordered key‑value storage.

- Compatibility – Supports major database protocol plugins (PostgreSQL, Redis, MongoDB, MySQL) easing migration.

- Reliability – ACID‑compliant storage and transaction semantics for correctness.

- Consolidation – Unified internal model representing key‑value, document, and relational data efficiently.

# Get Started

Start with the quick start guide:

- [Quick Start](doc/quick-start.md)

# Architecture & Concepts

Core architecture and design documents:

- [Design Concepts](doc/concept.md)

  - [Layer Concept](doc/layer-concept.md)

  - [Data Model](doc/data-model.md)

  - [Storage Concept](doc/storage-concept.md)

  - [Consistency Model](doc/consistency-model.md)

  - [Coordinator Concept](doc/coordinator-concept.md)

  - [Authentication Concept](doc/auth-concept.md)

  - [Plugin Concept](doc/plugin-concept.md)

# Supported Protocols

PuzzleDB currently supports the following database protocols:

- Compatibility

  - [PostgreSQL](doc/postgresql.md)

  - [MySQL](doc/mysql.md)

  - [MongoDB](doc/mongodb.md)

  - [Redis](doc/redis.md)

# Roadmap

PuzzleDB is in a technical preview stage. A distributed plugin stack is under active development. See roadmap:

![doc/img/roadmap](doc/img/roadmap.png)

# User Guides

- Get Started

  - [Quick Start](doc/quick-start.md)

    - [puzzledb-server](doc/cmd/server/puzzledb-server.md)

  - [Configuring PuzzleDB](doc/configuring.md)

- Operation

  - [CLI (puzzledb-cli)](doc/cmd/cli/puzzledb-cli.md)

  - [gRPC API](doc/grpc-api.md)

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

  - [Build on Ubuntu](doc/build-on-ubuntu.md)

- Extending PuzzleDB

  - [Plugin Concept](doc/plugin-concept.md)

  - [Plugin Services](doc/plugin-types.md)

    - [Building User Plugins](doc/plugin-tutorial.md)

- Specification

  - [Coordinator Specification](doc/spec/coordinator-spec.md)

    - [Coordinator Messaging Specification](doc/spec/coordinator-msg-spec.md)

    - [Coordinator Key-Value Store Specification](doc/spec/coordinator-spec.md)

  - Store Specifications

    - [Store Key-Value Specification](doc/spec/store-kv-spec.md)

  - Transverse Specifications

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
