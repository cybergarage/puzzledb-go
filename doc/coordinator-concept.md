# Coordinator Concept

In distributed mode, PuzzleDB assumes that it will be launched as multiple distributed instances, and each plug-in service (such as query plug-ins) launched on each instance should be coordinated using the coordinator service plug-in.

![architecture](img/architecture.png)

The coordinator service plug-in is used to manage and synchronize the distributed PuzzleDB nodes.

## References

-   Coordinator Services

    -   [The Chubby lock service for loosely-coupled distributed systems](https://research.google/pubs/pub41344/)

    -   [Apache ZooKeeper](https://zookeeper.apache.org/)

    -   [Consul by HashiCorp](https://www.consul.io/)

    -   [etcd by CoreOS](https://etcd.io/)

-   [Distributed Coordination. How distributed systems reach consensus | by Imesha Sudasingha | Medium](https://loneidealist.medium.com/distributed-coordination-5eb8eabb2ff)

-   [Apache Zookeeper vs etcd3. A comparison between distributed… | by Imesha Sudasingha | Medium](https://loneidealist.medium.com/apache-curator-vs-etcd3-9c1362600b26)
