# Architecture

PuzzleDB is a pluggable database that combines components, and a pluggable component interface is defined based on a FoundationDB-like layering concept. PuzzleDB separates the query layer and data model from the storage layer, the lowest storage layer is defined as a simple Key-Value store like FoundationDB and early Google Spanner.

![](img/architecture.png)

PuzzleDB defines the coordinator and storage function interfaces to run as standalone and distributed databases. PuzzleDB runs as a distributed multi-API and multi-model database with the distributed coordinator and storage plug-ins.

## Plug-In Concepts

PuzzleDB defines a core plug-in interface and basic component plug-in interfaces such as query, storage, and coordinator plug-in components based on the following concepts.

### Query Interface

PuzzleDB defines the query interface to support any database protocols such as Redis, MongoDB and MySQL protocols. The query interface is kept to a minimal specification in order to support a wide variety of database protocols.

### Storage Interface

PuzzleDB defines the low level storage interface as an ordered key-value store like early Google Spannerr. PuzzleDB expects that the storage plug-in components are implemented based on ordered key-value stores like FoundationDB rather than non-orders hashing key-value stores like MongoDB and Cassandra. 

### Coordinator Interface

PuzzleDB defines the coordinator interface to synchronize between PuzzleDB nodes. PuzzleDB expects that the coordinator components are implemented based on existing distributed coordinator services such as Apache ZooKeeper or etcd.

## References

- [FoundationDB](https://www.foundationdb.org/)
- [Layer Concept — FoundationDB ](https://apple.github.io/foundationdb/layer-concept.html)
- [Whitepapers  |  Cloud Spanner  |  Google Cloud](https://cloud.google.com/spanner/docs/whitepapers)
- [What is Cloud Spanner? A gcpsketchnote cheat sheet | Google Cloud Blog](https://cloud.google.com/blog/en/topics/developers-practitioners/what-cloud-spanner?hl=en)
- [Spanner: Google's Globally-Distributed Database](https://research.google/pubs/pub39966/)