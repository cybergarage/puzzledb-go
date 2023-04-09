![doc/img/logo](doc/img/logo.png)

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/puzzledb-go) [![Go](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/puzzledb-go/actions/workflows/make.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/puzzledb-go.svg)](https://pkg.go.dev/github.com/cybergarage/puzzledb-go)

PuzzleDB is a high-performance, distributed, cloud-native, multi-API, multi-model database.

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

-   [Quick Start](doc/quick_start.md)

# Supported Database Protocols

This technology preview version partially supports major database models and protocols. Please refer to the following documents for details on support status and limitations.

-   [Redis Compatibility](doc/redis.md)

-   [MongoDB Compatibility](doc/mongodb.md)

-   [MySQL Compatibility](doc/mysql.md)

-   PostgreSQL (Planning)

# How does PuzzleDB work?

For architecture of PuzzleDB, see the following concept documents:

-   Design Docs

    -   [Architecture](doc/architecture.md)

    -   [Data Model](doc/data_model.md)

    -   [Consistency Model](doc/consistency_model.md)

# Roadmap

PuzzleDB is currently in a technical preview release. Currently, PuzzleDB is in the process of developing and testing a distributed plugin for the next release. The development roadmap for PuzzleDB is shown below.

![doc/img/roadmap](doc/img/roadmap.png)

# Developer Guides

-   [Quick Start](doc/quick_start.md)

    -   [Build on macOS](doc/build-on-macos.md)

    -   [Build on Ubuntu](doc/build-on-macos.md)
