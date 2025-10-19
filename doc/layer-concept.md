# Layer Concept

PuzzleDB adopts an approach similar to FoundationDB and early Google Spanner: high scalability and ACID transactions built atop a simple ordered key‑value substrate without embedded query functionality.

![layer concept](img/layer_concept.png)

PuzzleDB loosely couples the query APIs, data model, and storage engine, enabling tailored compositions for specific workloads. Records, schemas, and indexes are all materialized as key‑value data.

# References

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
