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
