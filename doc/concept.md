# Design Concepts

PuzzleDB has a unique approach in the NewSQL field by using a simple key-value foundation for data model, indexes, and queries, enabling high scalability and ACID transactions.This section describes the architecture and design concepts of PuzzleDB.

# Layer Concept

PuzzleDB adopts a unique approach similar to FoundationDB and early Google Spanner. It offers high scalability and ACID transactions while constructing its data model, indexes, and query processing on a foundation of simple key-value storage without any query functionality.

![layer concept](img/layer_concept.png)

In contrast, PuzzleDB has loosely coupled the query API, data model, and storage engine, enabling users to build their database with a suitable combination for their specific use cases and workloads. In PuzzleDB, not only are records represented as key-value pairs, but schemas and indices are also represented as key-value data.

## References

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

# Data Model

PuzzleDB is a multi-data model database and the core data model is a document model, and the document model is constructed based on a key value model currently. PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as Key-Value objects.

<figure>
<img src="img/storage.png" alt="storage" />
</figure>

PuzzleDB defines a plug-in interface to the Key-Value store, which allows importing small local in-memory databases like memdb or large distributed databases like FoundationDB or TiKV.

## Document Model

PuzzleDB is a multi-data model database and the core data model is a document model like CosmosDB. PuzzleDB is a pluggable database that combines modules, and the storage layer modules must be as expressive as JSON or BSON like ARS (Atom-Record-Sequence) of CosmosDB.

PuzzleDB is a multi-model database, which converts any data models such as relational and document database models into the PuzzleDB data model as follows:

<table style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Type</th>
<th style="text-align: left;">PuzzleDB</th>
<th style="text-align: left;">Redis</th>
<th style="text-align: left;">MongoDB</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">PostgreSQL</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>map</p></td>
<td style="text-align: left;"><p>Hash</p></td>
<td style="text-align: left;"><p>Object</p></td>
<td style="text-align: left;"><p>COMPLEX</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>array</p></td>
<td style="text-align: left;"><p>List</p></td>
<td style="text-align: left;"><p>Array</p></td>
<td style="text-align: left;"><p>ARRAY</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Sets</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Sorted Sets</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>string</p></td>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>VARCHAR</p></td>
<td style="text-align: left;"><p>VARCHAR</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CHAR</p></td>
<td style="text-align: left;"><p>CHAR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Integer</p></td>
<td style="text-align: left;"><p>tiny</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>TINYINT</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>short</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>32-bit integer</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>long</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>64-bit integer</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Real</p></td>
<td style="text-align: left;"><p>float32</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>32-bit IEEE-754</p></td>
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>REAL</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>float64</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>64-bit IEEE-754</p></td>
<td style="text-align: left;"><p>DOUBLE (REAL)</p></td>
<td style="text-align: left;"><p>DOUBLE (REAL)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Time</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Date</p></td>
<td style="text-align: left;"><p>DATE DATETIME</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Timestamp</p></td>
<td style="text-align: left;"><p>TIME TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Special</p></td>
<td style="text-align: left;"><p>null</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Null</p></td>
<td style="text-align: left;"><p>NULL</p></td>
<td style="text-align: left;"><p>NULL</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>bool</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Boolean</p></td>
<td style="text-align: left;"><p>BOOLEAN (TINYINT(1))</p></td>
<td style="text-align: left;"><p>BOOLEAN</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>[]byte</p></td>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>Binary data</p></td>
<td style="text-align: left;"><p>BLOB (BYTEA)</p></td>
<td style="text-align: left;"><p>BINARY</p></td>
</tr>
</tbody>
</table>

### See also

-   [plugins.query.sql.NewDocumentElementTypeFrom()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/query/sql/type.go)

-   [plugins.query.mongo.BSONEncoder::EncodeBSON()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/query/mongo/encoder.go)

## Key-Value Object Model

PuzzleDB stores all database objects into key-value objects, and the key-value model is the core data model of PuzzleDB. The key-value model is a simple data model that stores data as a collection of key-value pairs. The key-value model is a flexible and scalable data model that can be used to store and retrieve data efficiently.

PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as key-value objects.

# Key Object

In PuzzleDB, records, schemas, and indices are all represented as key-value pairs. This section describes the format of the key object in detail.

## Key Header Specification

Every key object includes a header that specifies the key category, version, and the type of stored value. The key header is a 2-byte field prepended to every key in the key-value store. It is structured as follows:

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Field Name</th>
<th style="text-align: left;">Size (bits)</th>
<th style="text-align: left;">Description</th>
<th style="text-align: left;">Example Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Key category</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>The record key type</p></td>
<td style="text-align: left;"><p>D:Database C:Collection O:Document I:Index</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Version</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The version number</p></td>
<td style="text-align: left;"><p>0:reserved 1:Current</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Value type</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The record value type</p></td>
<td style="text-align: left;"><p>(Defined for each key category)</p></td>
</tr>
</tbody>
</table>

Key headers start with a one-byte identifier that indicates the type of key, enabling efficient searches based on key type. Currently, the version is fixed at `1`. The value type is specified individually for each key category. The values are specified as follows:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Value Types</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Index</p></td>
<td style="text-align: left;"><p>0:reserved 1:Primary 2:Secondary</p></td>
</tr>
</tbody>
</table>

## Key Categories

The key-value store consists of key-value records, where each record is defined by a key-value pair and includes a header as part of the key. The store supports the following categories of key-value records:

<table style="width:100%;">
<colgroup>
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Key Order</th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;">Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>5</p></td>
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>Header (D)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Options)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>Header (C)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Schema)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>Header (O)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>Primary Element Value</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Object)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Index</p></td>
<td style="text-align: left;"><p>Header (I)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>Secondary Element Name</p></td>
<td style="text-align: left;"><p>Secondary Element Value</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
</tbody>
</table>

Primary keys and secondary indices may comprise one or more columns. Although omitted in the table above, the combination of the element name and value for both objects and indices is repeated based on the index format. Additionally, since the primary key is stored in the key section of an index, the value section remains empty.

### See also

-   [plugins.coder.key.tuple.Coder::EncodeKey()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/coder/key/tuple/coder.go)

## Document (Value) Object

The document model is not natively implemented and is currently built on a key-value model with a coder plugin module. PuzzleDB provides a default coder, the CBOR (Concise Binary Object Representation ) plug-in module as the default coder.

PuzzleDB encodes a document data with a coder and stores it as a key-value data. The relationship between the default coder, CBOR data model, and the document data model is shown below.

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Type</th>
<th style="text-align: left;">PuzzleDB</th>
<th style="text-align: left;">CBOR</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>map</p></td>
<td style="text-align: left;"><p>5 (map)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>array</p></td>
<td style="text-align: left;"><p>4 (array)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>string</p></td>
<td style="text-align: left;"><p>3 (text string)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Integer</p></td>
<td style="text-align: left;"><p>tiny</p></td>
<td style="text-align: left;"><p>tiny</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>short</p></td>
<td style="text-align: left;"><p>short</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>long</p></td>
<td style="text-align: left;"><p>long</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Real</p></td>
<td style="text-align: left;"><p>float32</p></td>
<td style="text-align: left;"><p>7 (floating-point) 26</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>float64</p></td>
<td style="text-align: left;"><p>7 (floating-point) 27</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Time</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
<td style="text-align: left;"><p>6 (tag) 0</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Special</p></td>
<td style="text-align: left;"><p>null</p></td>
<td style="text-align: left;"><p>null</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>bool</p></td>
<td style="text-align: left;"><p>bool</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>[]byte</p></td>
<td style="text-align: left;"><p>binary</p></td>
</tr>
</tbody>
</table>

### See also

-   [plugins.coder.document.cbor.Coder::EncodeDocument()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/coder/document/cbor/coder.go)

## References

-   [A technical overview of Azure Cosmos DB | Azure Blog and Updates | Microsoft Azure](https://azure.microsoft.com/en-gb/blog/a-technical-overview-of-azure-cosmos-db/)

    -   [Azure Cosmos DB conceptual whitepapers](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

    -   [Schema-Agnostic Indexing with Azure DocumentDB](https://www.vldb.org/pvldb/vol8/p1668-shukla.pdf)

<!-- -->

-   [CBOR — Concise Binary Object Representation | Overview](http://cbor.io/)

# Storage Concepts

In PuzzleDB, the storage plugins are expected to be implemented as transaction-enabled, ordered sharding NoSQL storage systems, similar to Google Spanner or FoundationDB.

## Ordered Key-Value Store

PuzzleDB defines its storage interface as an ordered key-value store, akin to early Google Spanner and FoundationDB. PuzzleDB expects its storage plugin components to be implemented based on an ordered key-value store, in contrast to unordered hash-like key-value stores found in MongoDB and Cassandra. The implementation should be based on ACID-compliant ordered key-value stores.

FoundationDB and early Google Spanner utilize ordered key-value stores to support their unique features and capabilities in managing large-scale distributed databases. By organizing the keys in a sorted manner, these databases can optimize storage, retrieval, and update operations. This ordered structure also enables the databases to maintain consistency and achieve high performance in distributed environments.

Ordered key-value stores are a fundamental component of the storage layers in distributed databases like FoundationDB and Google Spanner. By maintaining keys in a sorted order, these systems can efficiently handle range queries and optimize various operations in large-scale distributed environments.

## Data Model

PuzzleDB is a multi-data model database and the core data model is a document model, and the document model is constructed based on a key value model currently. PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as Key-Value objects.

PuzzleDB is a multi-data model database and the core data model is a document model like CosmosDB. PuzzleDB is a pluggable database that combines modules, and the storage layer modules must be as expressive as JSON or BSON like ARS (Atom-Record-Sequence) of CosmosDB. For more detailed information about PuzzleDB’s data model, it is recommended to refer to the [Data Model](data-model.md) documents.

## References

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

<!-- -->

-   [A technical overview of Azure Cosmos DB | Azure Blog and Updates | Microsoft Azure](https://azure.microsoft.com/en-gb/blog/a-technical-overview-of-azure-cosmos-db/)

    -   [Azure Cosmos DB conceptual whitepapers](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

    -   [Schema-Agnostic Indexing with Azure DocumentDB](https://www.vldb.org/pvldb/vol8/p1668-shukla.pdf)

# Consistency Model

PuzzleDB is a multi-data model database; PuzzleDB is a pluggable database that combines modules, and the storage layer modules are expected to satisfy ACID-like interfaces.

PuzzleDB defines the top-level storage plug-in as a document model interface, and the storage interface consists of transaction and document interfaces.

<figure>
<img src="img/consistency_model.png" alt="consistency model" />
</figure>

While developers can omit the interface and implement the storage plug-ins based on non-ACID storage, such as contingent consistency model storage, PuzzleDB expects that storage modules are implemented based on ACID storages.

# Coordinator Concept

Coordinator services, such as Zookeeper and etcd, are distributed systems that play a crucial role in managing configuration data, synchronization, and coordination of distributed applications. They are designed to address the challenges of maintaining consistency and ensuring high availability in distributed environments.

In distributed mode, PuzzleDB operates under the assumption that it will be launched as multiple distributed instances. Each plugin service (such as query plugins) running on each instance should be coordinated using the coordinator service plugin.

![architecture](img/architecture.png)

The coordinator service provides distributed synchronization and coordination for PuzzleDB nodes. It is used to manage the distributed PuzzleDB nodes and synchronize the states of the nodes. The coordinator service plug-in is used to manage and synchronize the distributed PuzzleDB nodes.

## References

-   Coordinator Services

    -   [The Chubby lock service for loosely-coupled distributed systems](https://research.google/pubs/pub41344/)

    -   [Apache ZooKeeper](https://zookeeper.apache.org/)

    -   [Consul by HashiCorp](https://www.consul.io/)

    -   [etcd by CoreOS](https://etcd.io/)

-   [Distributed Coordination. How distributed systems reach consensus | by Imesha Sudasingha | Medium](https://loneidealist.medium.com/distributed-coordination-5eb8eabb2ff)

-   [Apache Zookeeper vs etcd3. A comparison between distributed… | by Imesha Sudasingha | Medium](https://loneidealist.medium.com/apache-curator-vs-etcd3-9c1362600b26)

# Authentication Methods

PuzzleDB includes a authenticator manager to manage the authentication for the query plugins.

<figure>
<img src="img/authenticator.png" alt="authenticator" />
</figure>

The authenticator manager supports multiple authentication methods, including username and password authentication, SASL (Simple Authentication and Security Layer) authentication, and certificate-based authentication.

## Authentication Plugins

PuzzleDB supports the following authentication methods for the query plugins.

-   Plain

-   SCRAM-SHA-256

-   Certificate (TLS Client Certificate)

-   MD5 (Not yes supported)

-   Crypt (Not yes supported)

-   LDAP (Not yes supported)

-   PAM (Not yes supported)

-   Kerberos (Not yes supported)

## Supported Authentication Methods

PuzzleDB supports the following authentication methods for the query plugins.

<table style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Method</th>
<th style="text-align: left;">Parameter</th>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">MongoDB</th>
<th style="text-align: left;">Redis</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Plain</p></td>
<td style="text-align: left;"><p>user</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>O</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>password</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>O</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>SCRAM-SHA-256</p></td>
<td style="text-align: left;"><p>user</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>password</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Certificate (TLS)</p></td>
<td style="text-align: left;"><p>common name</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>O</p></td>
</tr>
</tbody>
</table>

O:Supported, X:Unsupported, -:Not yes supported

## References

### PostgreSQL

-   [PostgreSQL: Documentation: Authentication Methods](https://www.postgresql.org/docs/current/auth-methods.html)

    -   [PostgreSQL: Documentation: The pg\_hba.conf File](https://www.postgresql.org/docs/current/auth-pg-hba-conf.html)

## MySQL

-   [MySQL: Connection Phase](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase.html)

-   [MySQL: Authentication Methods](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods.html)

    -   [MySQL: Old Password Authentication](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods.html#page_protocol_connection_phase_authentication_methods_old_password_authentication)

    -   [MySQL: Native Password Authentication](https://dev.mysql.com/doc/dev/mysql-server/latest/page_protocol_connection_phase_authentication_methods_native_password_authentication.html)

## MongoDB

-   [Security — MongoDB Manual](https://www.mongodb.com/docs/manual/security/)

    -   [Authentication — MongoDB Manual](https://www.mongodb.com/docs/manual/core/authentication/)

    -   [Configure Database User Authentication — MongoDB Atlas](https://www.mongodb.com/docs/atlas/security/config-db-auth/)

## Redis

-   [Security – Redis](https://redis.io/docs/management/security/)

    -   [AUTH | Redis](https://redis.io/commands/auth/)

# Plug-In Concepts

PuzzleDB is a pluggable database that amalgamates various components. It defines a pluggable component interface following a layering concept similar to FoundationDB. PuzzleDB separates the query layer and data model from the storage layer. The most basic storage layer is defined as a simple key-value store, much like FoundationDB and early Google Spanner.

![architecture](img/architecture.png)

PuzzleDB defines the coordinator and storage function interfaces to operate as standalone and distributed databases. Running with distributed coordinator and storage plug-ins, PuzzleDB functions as a distributed multi-API and multi-model database.

# Plug-In Service Types

PuzzleDB offers various types of plug-ins, including query, storage, and coordinator. These are categorized based on their support for distributed operations and their dependencies on other plug-ins. System plug-ins, responsible for managing configuration data and coordinating distributed nodes, are always activated by default. The database optimizes storage, retrieval, and update operations through a query interface that supports any database protocol, and a storage interface that employs an ordered key-value store, thereby maintaining consistency in distributed environments.

PuzzleDB provides default plug-in services that include query, storage, and coordinator plug-ins and defines the default plug-in types as follows:

<table style="width:100%;">
<colgroup>
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
<col style="width: 16%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Major Type</th>
<th style="text-align: left;">Sub Type</th>
<th style="text-align: left;">Description</th>
<th style="text-align: left;">Plug-ins</th>
<th style="text-align: left;">Distributed</th>
<th style="text-align: left;">Dependency</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>System</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>System services</p></td>
<td style="text-align: left;"><p>gRPC</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Actor</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Coordinator</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Query</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>Query handler services</p></td>
<td style="text-align: left;"><p>Redis</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>MongoDB</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>MySQL</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>PostgreSQL</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Coordinator</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>Coordination services</p></td>
<td style="text-align: left;"><p>memdb</p></td>
<td style="text-align: left;"><p>X</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>etcd (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>ZooKeeper (Planning)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>FoundationDB (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Coder</p></td>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>Document coder services</p></td>
<td style="text-align: left;"><p>CBOR</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key</p></td>
<td style="text-align: left;"><p>Key coder services</p></td>
<td style="text-align: left;"><p>Tuple</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Store</p></td>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>Doument store services</p></td>
<td style="text-align: left;"><p>Key-value based store</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Key-value), Coder (Document), Coder (Key)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key-value</p></td>
<td style="text-align: left;"><p>Key-value store services</p></td>
<td style="text-align: left;"><p>memdb</p></td>
<td style="text-align: left;"><p>X</p></td>
<td style="text-align: left;"><p>Coder (Document), Coder (Key)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>FoundationDB</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Coder (Document), Coder (Key)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>TiKV (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>JunoDB (Planning)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key-Value Cache</p></td>
<td style="text-align: left;"><p>Key-value cache store services</p></td>
<td style="text-align: left;"><p>Ristretto</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Key-value)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Tracer</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>Distributed tracing services</p></td>
<td style="text-align: left;"><p>OpenTelemetry</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>OpenTracing</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Metric</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>Metrics services</p></td>
<td style="text-align: left;"><p>Prometheus</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Graphite (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Extend</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>User-defined services</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
</tbody>
</table>

-   Distributed: Indicates whether the plug-in service supports distributed operation. The non-distributed plug-ins are provided for standalone operation or for internal testing of PuzzleDB.

-   Dependency: Indicates other plug-in service types required to run the plug-in service.

## Plug-In Interfaces

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

## References

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
