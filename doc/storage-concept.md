# Storage Concepts

Storage plugins should provide transaction‑enabled, ordered, sharded NoSQL capabilities (similar to Google Spanner or FoundationDB).

## Ordered Key-Value Store

PuzzleDB defines its storage interface as an ACID‑compliant ordered key‑value store (similar to early Spanner / FoundationDB), contrasting with unordered hash‑based stores.

Ordered stores optimize range scans, point lookups, and transactional semantics in large‑scale distributed deployments.

Sorted keys enable efficient range queries and predictable operational performance.

## Data Model

PuzzleDB’s logical layer is a document model encoded onto the ordered key‑value substrate. All objects (data, schema, indexes) are documents persisted as key‑value entries.

The document model must be expressive (JSON/BSON level) similar to CosmosDB’s ARS. See [Data Model](data-model.md) for details.

## References

- [FoundationDB](https://www.foundationdb.org/)

  - [Layer Concept — FoundationDB](https://apple.github.io/foundationdb/layer-concept.html)

  - [Announcing FoundationDB Document Layer](https://www.foundationdb.org/blog/announcing-document-layer/)

<!-- -->

- [Google Cloud Spanner](https://cloud.google.com/spanner/)

  - [Whitepapers | Cloud Spanner | Google Cloud](https://cloud.google.com/spanner/docs/whitepapers)

  - [What is Cloud Spanner? A gcpsketchnote cheat sheet | Google Cloud Blog](https://cloud.google.com/blog/en/topics/developers-practitioners/what-cloud-spanner?hl=en)

  - [F1: a distributed SQL database that scales: Proceedings of the VLDB Endowment: Vol 6, No 11](https://dl.acm.org/doi/10.14778/2536222.2536232)

  - [Spanner: Google’s Globally-Distributed Database](https://research.google/pubs/pub39966/)

  - [Spanner: Becoming a SQL System](https://dl.acm.org/doi/10.1145/3035918.3056103)

<!-- -->

- [A technical overview of Azure Cosmos DB | Azure Blog and Updates | Microsoft Azure](https://azure.microsoft.com/en-gb/blog/a-technical-overview-of-azure-cosmos-db/)

  - [Azure Cosmos DB conceptual whitepapers](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

  - [Schema-Agnostic Indexing with Azure DocumentDB](https://www.vldb.org/pvldb/vol8/p1668-shukla.pdf)
