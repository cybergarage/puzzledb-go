# Plug-In Services

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
<tr class="header">
<th style="text-align: left;">Major Type</th>
<th style="text-align: left;">Sub Type</th>
<th style="text-align: left;">Description</th>
<th style="text-align: left;">Plug-ins</th>
<th style="text-align: left;">Distributed</th>
<th style="text-align: left;">Dependency</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Query</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>Query handler services</p></td>
<td style="text-align: left;"><p>Redis</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>MongoDB</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>MySQL</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>PostgreSQL (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Document)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Coordinator</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>Coordination services</p></td>
<td style="text-align: left;"><p>memdb</p></td>
<td style="text-align: left;"><p>X</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>etcd (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>ZooKeeper (Planning)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>FoundationDB (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Encoder</p></td>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>Document serializer services</p></td>
<td style="text-align: left;"><p>CBOR</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key</p></td>
<td style="text-align: left;"><p>Key serializer services</p></td>
<td style="text-align: left;"><p>Tuple</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Store</p></td>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>Doument store services</p></td>
<td style="text-align: left;"><p>Key-value based store</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Key-value)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key-value</p></td>
<td style="text-align: left;"><p>Key-value store services</p></td>
<td style="text-align: left;"><p>memdb</p></td>
<td style="text-align: left;"><p>X</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>FoundationDB</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Key-value Cahche (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Store (Key-value), Coordinator</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>TiKV (Planning)</p></td>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Extend</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>User-defined services</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
</tbody>
</table>

-   Distributed: Indicates whether the plug-in service supports distributed operation. The non-distributed plug-ins are provided for standalone operation or for internal testing of PuzzleDB.

-   Dependency: Indicates other plug-in service types required to run the plug-in service.

## Plug-In Interfaces

PuzzleDB defines the core plug-in interfaces based on the following concepts.

### Query Interface

PuzzleDB defines the query interface to support any database protocol such as Redis, MongoDB, and MySQL protocols. The query interface is kept to a minimal specification to support a wide variety of database protocols.

### Storage Interface

PuzzleDB defines the low-level storage interface as an ordered key-value store like early Google Spanner. PuzzleDB expects that the storage plug-in components are implemented based on ordered key-value stores like FoundationDB rather than non-ordered hashing key-value stores like MongoDB and Cassandra.

### Coordinator Interface

PuzzleDB defines the coordinator interface to synchronize between PuzzleDB nodes. PuzzleDB expects that the coordinator components are implemented based on existing distributed coordinator services such as Apache ZooKeeper or etcd.
