# What is PuzzleDB

PuzzleDB is a multi-data model database capable of handling key-value, relational, and document models. Additionally, PuzzleDB is a multi-interface database, compatible with existing database protocols such as PostgreSQL, MySQL, Redis, and MongoDB.

![concept](img/concept.png)

PuzzleDB is a distributed database framework supporting various data models and protocols. It is designed as a flexible, scalable, and efficient database framework suitable for various environments.

![system](img/system.png)

PuzzleDB accommodates existing query protocols such as PostgreSQL, MySQL, MongoDB, and Redis within a distributed, pluggable database framework. Consequently, developers can seamlessly start using PuzzleDB as a scalable, high-performance distributed database with existing database client drivers, eliminating any learning curve.

## Key Features

PuzzleDB has the following features:

-   Flexibility: PuzzleDB allows for extensibility through its plugin architecture and pluggable modules for queries, data models, storage, and more.

-   Scalability: PuzzleDB seamlessly transitions from an in-memory standalone storage plugin module to a scalable, shared-nothing, horizontally distributed database using an ordered distributed key-value store plugin module.

-   Facility: PuzzleDB supports major database model and protocol plugin modules, such as PostgreSQL, Redis, MongoDB, and MySQL, simplifying application migration.

-   Safety: PuzzleDB offers ACID-compliant plugin modules, enabling the development of intuitive and secure applications.

-   Efficiency: PuzzleDB manages various database data models, including key-value, document, and relational, by consolidating them into a single core model.
