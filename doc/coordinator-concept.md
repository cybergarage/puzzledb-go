# Coordinator Concept

Coordinator services (e.g., ZooKeeper, etcd, Consul) manage configuration, synchronization, and coordination for distributed applications, helping maintain consistency and availability.

In distributed mode PuzzleDB runs as multiple instances. Each plugin service across instances is coordinated via the configured coordinator plugin.

<figure>
<img src="img/architecture.png" alt="architecture" />
</figure>

The coordinator plugin provides synchronization, membership management, and state propagation for PuzzleDB nodes.

## References

- Coordinator Services

  - [The Chubby lock service for loosely-coupled distributed systems](https://research.google/pubs/pub41344/)

  - [Apache ZooKeeper](https://zookeeper.apache.org/)

  - [Consul by HashiCorp](https://www.consul.io/)

  - [etcd by CoreOS](https://etcd.io/)

- [Distributed Coordination. How distributed systems reach consensus | by Imesha Sudasingha | Medium](https://loneidealist.medium.com/distributed-coordination-5eb8eabb2ff)

- [Apache Zookeeper vs etcd3. A comparison between distributed… | by Imesha Sudasingha | Medium](https://loneidealist.medium.com/apache-curator-vs-etcd3-9c1362600b26)
