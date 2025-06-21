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
<td style="text-align: left;"><p>FoundationDB</p></td>
<td style="text-align: left;"><p>O</p></td>
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

- Distributed: Indicates whether the plug-in service supports distributed operation. The non-distributed plug-ins are provided for standalone operation or for internal testing of PuzzleDB.

- Dependency: Indicates other plug-in service types required to run the plug-in service.

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
