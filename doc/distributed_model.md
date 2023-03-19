# Distributed Model

PuzzleDB is a pluggable database that combines modules, and a pluggable module interface is defined based on a FoundationDB-like layering concept. PuzzleDB separates the query layer and data model from the storage layer, the lowest storage layer is defined as a simple Key-Value store like FoundationDB and early Google Spanner.

PuzzleDB defines the coordinator and storage function interfaces to run as standalone and distributed databases. PuzzleDB runs as a distributed multi-API and multi-model database with the distributed coordinator and storage plug-ins.

![](img/distributed_model.png)

## Storage Interface

PuzzleDB defines the low level storage interface as an ordered key-value store like early Google Spannerr. PuzzleDB expects that the storage plugin modules are implemented based on ordered key-value stores like FoundationDB rather than non-orders hashing key-value stores like MongoDB and Cassandra. 

## Coordinator Interface

PuzzleDB defines the coordinator interface to synchronize between PuzzleDB nodes. PuzzleDB expects that the coordinator modules are implemented based on existing distributed coordinator services such as Apache ZooKeeper or etcd.

## References

- [FoundationDB](https://www.foundationdb.org/)
- [Layer Concept — FoundationDB ](https://apple.github.io/foundationdb/layer-concept.html)
- [Whitepapers  |  Cloud Spanner  |  Google Cloud](https://cloud.google.com/spanner/docs/whitepapers)
- [What is Cloud Spanner? A gcpsketchnote cheat sheet | Google Cloud Blog](https://cloud.google.com/blog/en/topics/developers-practitioners/what-cloud-spanner?hl=en)
- [Spanner: Google's Globally-Distributed Database](https://research.google/pubs/pub39966/)