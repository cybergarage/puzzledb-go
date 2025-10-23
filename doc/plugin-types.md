# Plugin Service Types

PuzzleDB offers several plugin categories (query, storage, coordinator, system). They are classified by distributed capability and dependency requirements. System plugins (configuration, coordination) are always active by default. Query plugins expose database protocols; storage plugins implement an ordered key‑value store to maintain consistency in distributed environments.

PuzzleDB provides default query, storage, coordinator, tracing, and metrics plugins. Types are defined below:

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
<td style="text-align: left;"><p>Job</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>Job services</p></td>
<td style="text-align: left;"><p>memdb</p></td>
<td style="text-align: left;"><p>X</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>FoundationDB</p></td>
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

- Distributed: Whether the plugin supports distributed operation (non‑distributed ones serve standalone or testing use cases).

- Dependency: Other plugin types required for activation.

## Plugin Interfaces

PuzzleDB defines plugin categories and interfaces as follows.

### System Plugins

System plugins manage configuration, synchronization, and coordination of distributed PuzzleDB nodes.

These are always activated by default. Some (e.g., gRPC) are independent; others (e.g., Actor) depend on additional plugins.

### Query Interface

Redis, MongoDB, MySQL, and PostgreSQL each use distinct wire protocols for handling queries. PuzzleDB’s query interface aims to support any database protocol with a minimal abstraction.

The abstraction is intentionally minimal to ease implementation of additional protocols.

### Storage Interface

The storage interface is an ACID‑compliant ordered key‑value abstraction (similar to early Spanner / FoundationDB), enabling efficient range operations and strong consistency.

Ordered storage optimizes range scans, point lookups, and transactional workloads in large‑scale distributed environments.

Maintaining keys in sorted order enables efficient range queries and predictable performance.

### Coordinator Interface

Coordinator plugins integrate external services (ZooKeeper, etcd, Consul) for cluster membership, leader election, and distributed state.

They provide synchronization and coordination primitives for PuzzleDB nodes.

### Tracer Interface

Tracing plugins implement distributed trace collection (e.g., OpenTelemetry, OpenTracing) for end‑to‑end request analysis and latency diagnostics.

The tracer interface is minimal to facilitate implementation of diverse tracing backends.

### Metrics Interface

Metrics plugins collect, store, and export time‑series performance data for monitoring and alerting.

The metrics interface is minimal to enable integration with systems like Prometheus or Graphite.
