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
