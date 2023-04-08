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

-   Support for existing query protocols:

    -   Redis, MongoDB and MySQL

-   Support for multi-data models:

    -   Key-value, Document and Relational.

# Roadmap

PuzzleDB is currently in a technical preview release. Currently, PuzzleDB is in the process of developing and testing a distributed plugin for the next release. The development roadmap for PuzzleDB is shown below.

![doc/img/roadmap](doc/img/roadmap.png)

# How does PuzzleDB work?

-   For architecture of PuzzleDB, see the following concept document:

    -   [Architecture](doc/architecture.md)

    -   [Data Model](doc/data_model.md)

    -   [Consistency Model](doc/consistency_model.md) \*\*

# Supported Protocols

This Technology Preview version provides staged support for the following major database protocols.

-   [Redis](doc/redis.md)

-   [MongoDB](doc/mongodb.md)

-   [MySQL](doc/mysql.md)

Please refer to the following documents for information regarding the support status and limitations of each protocol.

# Overview

-   [What is PuzzleDB](doc/whatis.md)

# User Guides

-   See [Quick Start](doc/quick_start.md) to learn about how to get started.

# Developer Guides

-   [Contribute](doc/contributing.md)

    -   [Build on macOS](doc/build-on-macos.md)

    -   [Build on Ubuntu](doc/build-on-macos.md)
