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
