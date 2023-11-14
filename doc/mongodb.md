# MongoDB Comptibility

PuzzleDB supports MongoDB API based on [go-mongo](https://github.com/cybergarage/go-mongo), a database framework that makes it easy to implement MongoDB compatible servers using Go.

<figure>
<img src="https://raw.githubusercontent.com/cybergarage/go-mongo/master/doc/img/framework.png" alt="framework" />
</figure>

The [go-mongo](https://github.com/cybergarage/go-mongo) framework automatically handles the MongoDB protocol and system commands. Therefore, PuzzleDB achieves MongoDB compatibility by implementing only user query commands.

## Data Model

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts MongoDB data model, [BSON (Binary JSON)](https://bsonspec.org/), into the PuzzleDB data model as follows:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">MongoDB</th>
<th style="text-align: left;">PuzzleDB</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Object</p></td>
<td style="text-align: left;"><p>map</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Array</p></td>
<td style="text-align: left;"><p>array</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>string</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32-bit integer</p></td>
<td style="text-align: left;"><p>int</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>64-bit integer</p></td>
<td style="text-align: left;"><p>long</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>32-bit IEEE-754</p></td>
<td style="text-align: left;"><p>float32</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>64-bit IEEE-754</p></td>
<td style="text-align: left;"><p>float64</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Date</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Timestamp</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Null</p></td>
<td style="text-align: left;"><p>null</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Boolean</p></td>
<td style="text-align: left;"><p>bool</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Binary data</p></td>
<td style="text-align: left;"><p>[]byte</p></td>
</tr>
</tbody>
</table>

## Supported Commands

PuzzleDB currently supports [MongoDB database commands](https://www.mongodb.com/docs/manual/reference/command/) in stages. This section describes the status of MongoDB command support in PuzzleDB.

### Diagnostic Commands

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Diagnostic Command</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>buildInfo</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### Replication Commands

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Replication Command</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>isMaster</p></td>
<td style="text-align: left;"><p>Always returns true</p></td>
</tr>
</tbody>
</table>

### Query and Write Operation Commands

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Query Command</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>delete</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>find</p></td>
<td style="text-align: left;"><p>Only $eq operation supported</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>findAndModify</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>getMore</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>insert</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>resetError</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>update</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

## Indexing

Currently, PuzzleDB automatically indexes all sigle fields of inserted and updated documents by default, as CosmosDB does. In the future, PuzzleDB will support indexing of only the specified fields like MongoDB or more smart indexing like CosmosDB.

## References

-   [Conceptual whitepapers - Azure Cosmos DB | Microsoft Learn](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

-   [Azure Cosmos DB indexing policies | Microsoft Learn](https://learn.microsoft.com/en-us/azure/cosmos-db/index-policy)

-   [Schema-Agnostic Indexing with Azure DocumentDB](https://www.microsoft.com/en-us/research/publication/schema-agnostic-indexing-azure-documentdb/)

-   [Manage indexing in Azure Cosmos DB for MongoDB | Microsoft Learn](https://learn.microsoft.com/en-us/azure/cosmos-db/mongodb/indexing)
