MongoDB Comptibility
====================

PuzzleDB supports MongoDB API based on [go-mongo](https://github.com/cybergarage/go-mongo), a database framework that makes it easy to implement MongoDB compatible servers using Go.

![framework](https://raw.githubusercontent.com/cybergarage/go-mongo/master/doc/img/framework.png)

The [go-mongo](https://github.com/cybergarage/go-mongo) framework automatically handles the MongoDB protocol and system commands. Therefore, PuzzleDB achieves MongoDB compatibility by implementing only user query commands.

Data Model
----------

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts MongoDB data model, [BSON (Binary JSON)](https://bsonspec.org/), into the PuzzleDB data model as follows:

<table><colgroup><col style="width: 50%" /><col style="width: 50%" /></colgroup><thead><tr class="header"><th>MongoDB</th><th>PuzzleDB</th></tr></thead><tbody><tr class="odd"><td><p>Object</p></td><td><p>map</p></td></tr><tr class="even"><td><p>Array</p></td><td><p>array</p></td></tr><tr class="odd"><td><p>String</p></td><td><p>string</p></td></tr><tr class="even"><td><p>32-bit integer</p></td><td><p>int</p></td></tr><tr class="odd"><td><p>64-bit integer</p></td><td><p>long</p></td></tr><tr class="even"><td><p>32-bit IEEE-754</p></td><td><p>float32</p></td></tr><tr class="odd"><td><p>64-bit IEEE-754</p></td><td><p>float64</p></td></tr><tr class="even"><td><p>Date</p></td><td><p>time.Time</p></td></tr><tr class="odd"><td><p>Timestamp</p></td><td><p>time.Time</p></td></tr><tr class="even"><td><p>Null</p></td><td><p>null</p></td></tr><tr class="odd"><td><p>Boolean</p></td><td><p>bool</p></td></tr><tr class="even"><td><p>Binary data</p></td><td><p>[]byte</p></td></tr></tbody></table>

Supported Commands
------------------

PuzzleDB currently supports [MongoDB database commands](https://www.mongodb.com/docs/manual/reference/command/) in stages. This section describes the status of MongoDB command support in PuzzleDB.

### Diagnostic Commands

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Diagnostic Command</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>buildInfo</p></td><td></td></tr></tbody></table>

### Replication Commands

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Replication Command</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>isMaster</p></td><td><p>Always returns true</p></td></tr></tbody></table>

### Query and Write Operation Commands

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Query Command</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>delete</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>find</p></td><td><p>Only $eq operation supported</p></td></tr><tr class="odd"><td><p>-</p></td><td><p>findAndModify</p></td><td></td></tr><tr class="even"><td><p>-</p></td><td><p>getMore</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td><p>insert</p></td><td></td></tr><tr class="even"><td><p>-</p></td><td><p>resetError</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td><p>update</p></td><td></td></tr></tbody></table>

Indexing
--------

Currently, PuzzleDB automatically indexes all sigle fields of inserted and updated documents by default, as CosmosDB does. In the future, PuzzleDB will support indexing of only the specified fields like MongoDB or more smart indexing like CosmosDB.

References
----------

-   [Conceptual whitepapers - Azure Cosmos DB | Microsoft Learn](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

-   [Azure Cosmos DB indexing policies | Microsoft Learn](https://learn.microsoft.com/en-us/azure/cosmos-db/index-policy)

-   [Schema-Agnostic Indexing with Azure DocumentDB](https://www.microsoft.com/en-us/research/publication/schema-agnostic-indexing-azure-documentdb/)

-   [Manage indexing in Azure Cosmos DB for MongoDB | Microsoft Learn](https://learn.microsoft.com/en-us/azure/cosmos-db/mongodb/indexing)
