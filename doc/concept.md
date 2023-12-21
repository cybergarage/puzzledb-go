Design Concepts
===============

PuzzleDB has a unique approach in the NewSQL field by using a simple key-value foundation for data model, indexes, and queries, enabling high scalability and ACID transactions.This section describes the architecture and design concepts of PuzzleDB.

Plug-In Concepts
================

PuzzleDB is a pluggable database that amalgamates various components. It defines a pluggable component interface following a layering concept similar to FoundationDB. PuzzleDB separates the query layer and data model from the storage layer. The most basic storage layer is defined as a simple key-value store, much like FoundationDB and early Google Spanner.

![architecture](img/architecture.png)

PuzzleDB defines the coordinator and storage function interfaces to operate as standalone and distributed databases. Running with distributed coordinator and storage plug-ins, PuzzleDB functions as a distributed multi-API and multi-model database.

Plug-In Service Types
=====================

PuzzleDB offers various types of plug-ins, including query, storage, and coordinator. These are categorized based on their support for distributed operations and their dependencies on other plug-ins. System plug-ins, responsible for managing configuration data and coordinating distributed nodes, are always activated by default. The database optimizes storage, retrieval, and update operations through a query interface that supports any database protocol, and a storage interface that employs an ordered key-value store, thereby maintaining consistency in distributed environments.

PuzzleDB provides default plug-in services that include query, storage, and coordinator plug-ins and defines the default plug-in types as follows:

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Major Type</th><th>Sub Type</th><th>Description</th><th>Plug-ins</th><th>Distributed</th><th>Dependency</th></tr></thead><tbody><tr class="odd"><td><p>System</p></td><td><p>-</p></td><td><p>System services</p></td><td><p>gRPC</p></td><td><p>O</p></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>Actor</p></td><td><p>O</p></td><td><p>Coordinator</p></td></tr><tr class="odd"><td><p>Query</p></td><td><p>-</p></td><td><p>Query handler services</p></td><td><p>Redis</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>MongoDB</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>MySQL</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>PostgreSQL</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="odd"><td><p>Coordinator</p></td><td><p>-</p></td><td><p>Coordination services</p></td><td><p>memdb</p></td><td><p>X</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>etcd (Planning)</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>ZooKeeper (Planning)</p></td><td></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>FoundationDB (Planning)</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td><p>Coder</p></td><td><p>Document</p></td><td><p>Document coder services</p></td><td><p>CBOR</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td><p>Key</p></td><td><p>Key coder services</p></td><td><p>Tuple</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td><p>Store</p></td><td><p>Document</p></td><td><p>Doument store services</p></td><td><p>Key-value based store</p></td><td><p>O</p></td><td><p>Store (Key-value), Coder (Document), Coder (Key)</p></td></tr><tr class="even"><td></td><td><p>Key-value</p></td><td><p>Key-value store services</p></td><td><p>memdb</p></td><td><p>X</p></td><td><p>Coder (Document), Coder (Key)</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>FoundationDB</p></td><td><p>O</p></td><td><p>Coder (Document), Coder (Key)</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>TiKV (Planning)</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>JunoDB (Planning)</p></td><td></td><td></td></tr><tr class="even"><td></td><td><p>Key-Value Cache</p></td><td><p>Key-value cache store services</p></td><td><p>Ristretto</p></td><td><p>O</p></td><td><p>Store (Key-value)</p></td></tr><tr class="odd"><td><p>Tracer</p></td><td><p>-</p></td><td><p>Distributed tracing services</p></td><td><p>OpenTelemetry</p></td><td><p>O</p></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>OpenTracing</p></td><td><p>O</p></td><td></td></tr><tr class="odd"><td><p>Metric</p></td><td><p>-</p></td><td><p>Metrics services</p></td><td><p>Prometheus</p></td><td><p>O</p></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>Graphite (Planning)</p></td><td><p>O</p></td><td></td></tr><tr class="odd"><td><p>Extend</p></td><td><p>-</p></td><td><p>User-defined services</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr></tbody></table>

-   Distributed: Indicates whether the plug-in service supports distributed operation. The non-distributed plug-ins are provided for standalone operation or for internal testing of PuzzleDB.

-   Dependency: Indicates other plug-in service types required to run the plug-in service.

Plug-In Interfaces
------------------

PuzzleDB defines the plug-in categories and interfaces based on the following concepts.

### System Plug-Ins

System plug-ins are used to manage the PuzzleDB system. They are used to manage the configuration data, synchronization, and coordination of distributed PuzzleDB nodes. System plug-ins are used to manage and synchronize the distributed PuzzleDB nodes.

Unlike other plugins, system plugins are always activated as default plugins. Some, such as the gRPC plugin, work independently, while others, such as the Actor service, depend on other plugins to function.

### Query Interface

Redis, MongoDB, and MySQL are popular database management systems, each with its own communication protocol for handling database queries. These protocols enable clients to interact with the database server, performing various operations such as inserting, updating, retrieving, or deleting data.

PuzzleDB defines the query interface to support any database protocol such as Redis, MongoDB, and MySQL protocols. The query interface is kept to a minimal specification to support a wide variety of database protocols.

### Storage Interface

PuzzleDB defines the storage interface as an ordered key-value store, similar to early Google Spanner and FoundationDB. PuzzleDB expects its storage plugin components to be implemented based on an ordered key-value store, in contrast to unordered hash-like key-value stores found in MongoDB and Cassandra. The implementation should be based on ACID-compliant ordered key-value stores.

FoundationDB and early Google Spanner utilize ordered key-value stores to support their unique features and capabilities in managing large-scale distributed databases. By organizing the keys in a sorted manner, these databases can optimize storage, retrieval, and update operations. This ordered structure also enables the databases to maintain consistency and achieve high performance in distributed environments.

Ordered key-value stores are a fundamental component of the storage layers in distributed databases like FoundationDB and Google Spanner. By maintaining keys in a sorted order, these systems can efficiently handle range queries and optimize various operations in large-scale distributed environments.

### Coordinator Interface

Coordinator services, such as Zookeeper and etcd, are distributed systems that play a crucial role in managing the configuration data, synchronization, and coordination of distributed applications. They are designed to handle the challenges of maintaining consistency and ensuring high availability in distributed environments.

The coordinator service provides distributed synchronization and coordination for PuzzleDB nodes. It is used to manage the distributed PuzzleDB nodes and synchronize the states of the nodes. The coordinator service plug-in is used to manage and synchronize the distributed PuzzleDB nodes.

### Tracer Interface

Distributed tracing is a monitoring technique for analyzing and troubleshooting distributed systems like microservices and cloud-based applications. It tracks requests as they flow through various services, identifying bottlenecks and performance issues. Unique trace IDs tag requests, and spans represent each step in the request lifecycle. Visualization tools display interactions between components, aiding in issue detection and system optimization. Distributed tracing is essential for modern software systems, helping improve performance and reliability.

PuzzleDB defines the tracer service interface to support any distributed tracing protocol such as OpenTracing and OpenTelemetry. The tracer interface is kept to a minimal specification to support a wide variety of tracer protocols.

### Metrics Interface

Metric service is a tool or platform used for collecting, storing, and analyzing metric data. Metric data is time-series data that describes the behavior and performance of a system or application over time. Metric services allow organizations to monitor their systems and applications in real-time, gain insights into performance trends, and detect and troubleshoot issues.

PuzzleDB defines the metrics service interface to support any metrics servicel such as Prometheus and Graphite. The metrics interface is kept to a minimal specification to support a wide variety of metrics services.

References
----------

-   [FoundationDB](https://www.foundationdb.org/)

    -   [Layer Concept — FoundationDB](https://apple.github.io/foundationdb/layer-concept.html)

    -   [Announcing FoundationDB Document Layer](https://www.foundationdb.org/blog/announcing-document-layer/)

<!-- -->

-   [Google Cloud Spanner](https://cloud.google.com/spanner/)

    -   [Whitepapers | Cloud Spanner | Google Cloud](https://cloud.google.com/spanner/docs/whitepapers)

    -   [What is Cloud Spanner? A gcpsketchnote cheat sheet | Google Cloud Blog](https://cloud.google.com/blog/en/topics/developers-practitioners/what-cloud-spanner?hl=en)

    -   [F1: a distributed SQL database that scales: Proceedings of the VLDB Endowment: Vol 6, No 11](https://dl.acm.org/doi/10.14778/2536222.2536232)

    -   [Spanner: Google’s Globally-Distributed Database](https://research.google/pubs/pub39966/)

    -   [Spanner: Becoming a SQL System](https://dl.acm.org/doi/10.1145/3035918.3056103)

Layer Concept
=============

PuzzleDB adopts a unique approach similar to FoundationDB and early Google Spanner. It offers high scalability and ACID transactions while constructing its data model, indexes, and query processing on a foundation of simple key-value storage without any query functionality.

![layer concept](img/layer_concept.png)

In contrast, PuzzleDB has loosely coupled the query API, data model, and storage engine, enabling users to build their database with a suitable combination for their specific use cases and workloads. In PuzzleDB, not only are records represented as key-value pairs, but schemas and indices are also represented as key-value data.

References
----------

-   [FoundationDB](https://www.foundationdb.org/)

    -   [Layer Concept — FoundationDB](https://apple.github.io/foundationdb/layer-concept.html)

    -   [Announcing FoundationDB Document Layer](https://www.foundationdb.org/blog/announcing-document-layer/)

<!-- -->

-   [Google Cloud Spanner](https://cloud.google.com/spanner/)

    -   [Whitepapers | Cloud Spanner | Google Cloud](https://cloud.google.com/spanner/docs/whitepapers)

    -   [What is Cloud Spanner? A gcpsketchnote cheat sheet | Google Cloud Blog](https://cloud.google.com/blog/en/topics/developers-practitioners/what-cloud-spanner?hl=en)

    -   [F1: a distributed SQL database that scales: Proceedings of the VLDB Endowment: Vol 6, No 11](https://dl.acm.org/doi/10.14778/2536222.2536232)

    -   [Spanner: Google’s Globally-Distributed Database](https://research.google/pubs/pub39966/)

    -   [Spanner: Becoming a SQL System](https://dl.acm.org/doi/10.1145/3035918.3056103)

Data Model
==========

PuzzleDB is a multi-data model database and the core data model is a document model, and the document model is constructed based on a key value model currently. PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as Key-Value objects.

![storage](img/storage.png)

PuzzleDB defines a plug-in interface to the Key-Value store, which allows importing small local in-memory databases like memdb or large distributed databases like FoundationDB or TiKV.

Document Model
--------------

PuzzleDB is a multi-data model database and the core data model is a document model like CosmosDB. PuzzleDB is a pluggable database that combines modules, and the storage layer modules must be as expressive as JSON or BSON like ARS (Atom-Record-Sequence) of CosmosDB.

PuzzleDB is a multi-model database, which converts any data models such as relational and document database models into the PuzzleDB data model as follows:

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Type</th><th>PuzzleDB</th><th>Redis</th><th>MongoDB</th><th>MySQL</th><th>PostgreSQL</th></tr></thead><tbody><tr class="odd"><td><p>Collection</p></td><td><p>map</p></td><td><p>Hash</p></td><td><p>Object</p></td><td><p>COMPLEX</p></td><td></td></tr><tr class="even"><td></td><td><p>array</p></td><td><p>List</p></td><td><p>Array</p></td><td><p>ARRAY</p></td><td></td></tr><tr class="odd"><td></td><td></td><td><p>Sets</p></td><td></td><td></td><td></td></tr><tr class="even"><td></td><td></td><td><p>Sorted Sets</p></td><td></td><td></td><td></td></tr><tr class="odd"><td><p>String</p></td><td><p>string</p></td><td><p>String</p></td><td><p>String</p></td><td><p>TEXT</p></td><td><p>TEXT</p></td></tr><tr class="even"><td></td><td></td><td></td><td></td><td><p>VARCHAR</p></td><td><p>VARCHAR</p></td></tr><tr class="odd"><td></td><td></td><td></td><td></td><td><p>CHAR</p></td><td><p>CHAR</p></td></tr><tr class="even"><td><p>Integer</p></td><td><p>tiny</p></td><td></td><td></td><td><p>TINYINT</p></td><td></td></tr><tr class="odd"><td></td><td><p>short</p></td><td></td><td></td><td><p>SMALLINT</p></td><td><p>SMALLINT</p></td></tr><tr class="even"><td></td><td><p>int</p></td><td></td><td><p>32-bit integer</p></td><td><p>INTEGER</p></td><td><p>INTEGER</p></td></tr><tr class="odd"><td></td><td><p>long</p></td><td></td><td><p>64-bit integer</p></td><td><p>BIGINT</p></td><td><p>BIGINT</p></td></tr><tr class="even"><td><p>Real</p></td><td><p>float32</p></td><td></td><td><p>32-bit IEEE-754</p></td><td><p>FLOAT</p></td><td><p>REAL</p></td></tr><tr class="odd"><td></td><td><p>float64</p></td><td></td><td><p>64-bit IEEE-754</p></td><td><p>DOUBLE (REAL)</p></td><td><p>DOUBLE (REAL)</p></td></tr><tr class="even"><td><p>Time</p></td><td><p>time.Time</p></td><td></td><td><p>Date</p></td><td><p>DATE DATETIME</p></td><td></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>Timestamp</p></td><td><p>TIME TIMESTAMP</p></td><td><p>TIMESTAMP</p></td></tr><tr class="even"><td><p>Special</p></td><td><p>null</p></td><td></td><td><p>Null</p></td><td><p>NULL</p></td><td><p>NULL</p></td></tr><tr class="odd"><td></td><td><p>bool</p></td><td></td><td><p>Boolean</p></td><td><p>BOOLEAN (TINYINT(1))</p></td><td><p>BOOLEAN</p></td></tr><tr class="even"><td></td><td><p>[]byte</p></td><td><p>String</p></td><td><p>Binary data</p></td><td><p>BLOB (BYTEA)</p></td><td><p>BINARY</p></td></tr></tbody></table>

Key-Value Model
---------------

PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as Key-Value objects.

The document model is not natively implemented and is currently built on a key-value model with a coder plugin module. PuzzleDB provides a default coder, the CBOR (Concise Binary Object Representation ) plug-in module as the default coder.

PuzzleDB encodes a document data with a coder and stores it as a key-value data. The relationship between the default coder, CBOR data model, and the document data model is shown below.

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Type</th><th>PuzzleDB</th><th>CBOR</th></tr></thead><tbody><tr class="odd"><td><p>Collection</p></td><td><p>map</p></td><td><p>5 (map)</p></td></tr><tr class="even"><td></td><td><p>array</p></td><td><p>4 (array)</p></td></tr><tr class="odd"><td><p>String</p></td><td><p>string</p></td><td><p>3 (text string)</p></td></tr><tr class="even"><td><p>Integer</p></td><td><p>tiny</p></td><td><p>tiny</p></td></tr><tr class="odd"><td></td><td><p>short</p></td><td><p>short</p></td></tr><tr class="even"><td></td><td><p>int</p></td><td><p>int</p></td></tr><tr class="odd"><td></td><td><p>long</p></td><td><p>long</p></td></tr><tr class="even"><td><p>Real</p></td><td><p>float32</p></td><td><p>7 (floating-point) 26</p></td></tr><tr class="odd"><td></td><td><p>float64</p></td><td><p>7 (floating-point) 27</p></td></tr><tr class="even"><td><p>Time</p></td><td><p>time.Time</p></td><td><p>6 (tag) 0</p></td></tr><tr class="odd"><td><p>Special</p></td><td><p>null</p></td><td><p>null</p></td></tr><tr class="even"><td></td><td><p>bool</p></td><td><p>bool</p></td></tr><tr class="odd"><td></td><td><p>[]byte</p></td><td><p>binary</p></td></tr></tbody></table>

References
----------

-   [A technical overview of Azure Cosmos DB | Azure Blog and Updates | Microsoft Azure](https://azure.microsoft.com/en-gb/blog/a-technical-overview-of-azure-cosmos-db/)

    -   [Azure Cosmos DB conceptual whitepapers](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

    -   [Schema-Agnostic Indexing with Azure DocumentDB](https://www.vldb.org/pvldb/vol8/p1668-shukla.pdf)

<!-- -->

-   [CBOR — Concise Binary Object Representation | Overview](http://cbor.io/)

Unresolved directive in concept.adoc - include::storage-model.adoc\[leveloffset=+1\] :leveloffset: +1

Consistency Model
=================

PuzzleDB is a multi-data model database; PuzzleDB is a pluggable database that combines modules, and the storage layer modules are expected to satisfy ACID-like interfaces.

PuzzleDB defines the top-level storage plug-in as a document model interface, and the storage interface consists of transaction and document interfaces.

![consistency model](img/consistency_model.png)

While developers can omit the interface and implement the storage plug-ins based on non-ACID storage, such as contingent consistency model storage, PuzzleDB expects that storage modules are implemented based on ACID storages.

Unresolved directive in concept.adoc - include::coordinator-model.adoc\[leveloffset=+1\]
