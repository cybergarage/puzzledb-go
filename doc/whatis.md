# What is PuzzleDB

PuzzleDB is a multi-data model database that handles key-value model, relational model, and document model. In addition, PuzzleDB is a multi-API database and is compatible with existing database protocols such as MySQL, Redis, and MongoDB.

![concept](img/concept.png)

PuzzleDB supports existing query protocols such as MongoDB, Redia, and MySQL. Thus, developers can start using PuzzleDB as a scalable, high-performance distributed database with existing database client drivers without any learning curve.

The name PuzzleDB comes from the ability to combine multiple modules such as coordinators, storages, and existing database protocol handlers to form a database.

## Features

PuzzleDB has the following features:

-   Flexibility: PuzzleDB allows extensibility through its plugin architecture and pluggable modules for queries, data models, storage, and more.

-   Scalability: PuzzleDB seamlessly transitions from an in-memory standalone storage plugin module to a scalable, shared-nothing, horizontally distributed database using an ordered distributed key-value store plugin module.

-   Facility: PuzzleDB supports major database model and protocol plugin modules, such as Redis, MongoDB, and MySQL, simplifying application migration.

-   Safety: PuzzleDB offers ACID-compliant plugin modules, enabling the development of intuitive and secure applications.

-   Efficiency: PuzzleDB manages various database data models, including key-value, document, and relational, by consolidating them into a single core model.
