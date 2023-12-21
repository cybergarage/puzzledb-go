Data Model
==========

PuzzleDB is a multi-data model database and the core data model is a document model, and the document model is constructed based on a key value model currently. PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as Key-Value objects.

![storage](img/storage.png)

PuzzleDB defines a plug-in interface to the Key-Value store, which allows importing small local in-memory databases like memdb or large distributed databases like FoundationDB or TiKV.

Document Model
--------------

PuzzleDB is a multi-data model database and the core data model is a document model like CosmosDB. PuzzleDB is a pluggable database that combines modules, and the storage layer modules must be as expressive as JSON or BSON like ARS (Atom-Record-Sequence) of CosmosDB.

PuzzleDB is a multi-model database, which converts any data models such as relational and document database models into the PuzzleDB data model as follows:

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Type</th><th>PuzzleDB</th><th>Redis</th><th>MongoDB</th><th>MySQL</th><th>PostgreSQL</th></tr></thead><tbody><tr class="odd"><td><p>Collection</p></td><td><p>map</p></td><td><p>Hash</p></td><td><p>Object</p></td><td><p>COMPLEX</p></td><td></td></tr><tr class="even"><td></td><td><p>array</p></td><td><p>List</p></td><td><p>Array</p></td><td><p>ARRAY</p></td><td></td></tr><tr class="odd"><td></td><td></td><td><p>Sets</p></td><td></td><td></td><td></td></tr><tr class="even"><td></td><td></td><td><p>Sorted Sets</p></td><td></td><td></td><td></td></tr><tr class="odd"><td><p>String</p></td><td><p>string</p></td><td><p>String</p></td><td><p>String</p></td><td><p>TEXT</p></td><td><p>TEXT</p></td></tr><tr class="even"><td></td><td></td><td></td><td></td><td><p>VARCHAR</p></td><td><p>VARCHAR</p></td></tr><tr class="odd"><td></td><td></td><td></td><td></td><td><p>CHAR</p></td><td><p>CHAR</p></td></tr><tr class="even"><td><p>Integer</p></td><td><p>tiny</p></td><td></td><td></td><td><p>TINYINT</p></td><td></td></tr><tr class="odd"><td></td><td><p>short</p></td><td></td><td></td><td><p>SMALLINT</p></td><td><p>SMALLINT</p></td></tr><tr class="even"><td></td><td><p>int</p></td><td></td><td><p>32-bit integer</p></td><td><p>INTEGER</p></td><td><p>INTEGER</p></td></tr><tr class="odd"><td></td><td><p>long</p></td><td></td><td><p>64-bit integer</p></td><td><p>BIGINT</p></td><td><p>BIGINT</p></td></tr><tr class="even"><td><p>Real</p></td><td><p>float32</p></td><td></td><td><p>32-bit IEEE-754</p></td><td><p>FLOAT</p></td><td><p>REAL</p></td></tr><tr class="odd"><td></td><td><p>float64</p></td><td></td><td><p>64-bit IEEE-754</p></td><td><p>DOUBLE (REAL)</p></td><td><p>DOUBLE (REAL)</p></td></tr><tr class="even"><td><p>Time</p></td><td><p>time.Time</p></td><td></td><td><p>Date</p></td><td><p>DATE DATETIME</p></td><td></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>Timestamp</p></td><td><p>TIME TIMESTAMP</p></td><td><p>TIMESTAMP</p></td></tr><tr class="even"><td><p>Special</p></td><td><p>null</p></td><td></td><td><p>Null</p></td><td><p>NULL</p></td><td><p>NULL</p></td></tr><tr class="odd"><td></td><td><p>bool</p></td><td></td><td><p>Boolean</p></td><td><p>BOOLEAN (TINYINT(1))</p></td><td><p>BOOLEAN</p></td></tr><tr class="even"><td></td><td><p>[]byte</p></td><td><p>String</p></td><td><p>Binary data</p></td><td><p>BLOB (BYTEA)</p></td><td><p>BINARY</p></td></tr></tbody></table>

Key-Value Model
---------------

PuzzleDB represents all database objects such as data objects, schema objects, and index objects as document data. Document data are ultimately stored as Key-Value objects.

The document model is not natively implemented and is currently built on a key-value model with a coder plugin module. PuzzleDB provides a default coder, the CBOR (Concise Binary Object Representation ) plug-in module as the default coder.

PuzzleDB encodes a document data with a coder and stores it as a key-value data. The relationship between the default coder, CBOR data model, and the document data model is shown below.

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Type</th><th>PuzzleDB</th><th>CBOR</th></tr></thead><tbody><tr class="odd"><td><p>Collection</p></td><td><p>map</p></td><td><p>5 (map)</p></td></tr><tr class="even"><td></td><td><p>array</p></td><td><p>4 (array)</p></td></tr><tr class="odd"><td><p>String</p></td><td><p>string</p></td><td><p>3 (text string)</p></td></tr><tr class="even"><td><p>Integer</p></td><td><p>tiny</p></td><td><p>tiny</p></td></tr><tr class="odd"><td></td><td><p>short</p></td><td><p>short</p></td></tr><tr class="even"><td></td><td><p>int</p></td><td><p>int</p></td></tr><tr class="odd"><td></td><td><p>long</p></td><td><p>long</p></td></tr><tr class="even"><td><p>Real</p></td><td><p>float32</p></td><td><p>7 (floating-point) 26</p></td></tr><tr class="odd"><td></td><td><p>float64</p></td><td><p>7 (floating-point) 27</p></td></tr><tr class="even"><td><p>Time</p></td><td><p>time.Time</p></td><td><p>6 (tag) 0</p></td></tr><tr class="odd"><td><p>Special</p></td><td><p>null</p></td><td><p>null</p></td></tr><tr class="even"><td></td><td><p>bool</p></td><td><p>bool</p></td></tr><tr class="odd"><td></td><td><p>[]byte</p></td><td><p>binary</p></td></tr></tbody></table>

References
----------

-   [A technical overview of Azure Cosmos DB | Azure Blog and Updates | Microsoft Azure](https://azure.microsoft.com/en-gb/blog/a-technical-overview-of-azure-cosmos-db/)

    -   [Azure Cosmos DB conceptual whitepapers](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

    -   [Schema-Agnostic Indexing with Azure DocumentDB](https://www.vldb.org/pvldb/vol8/p1668-shukla.pdf)

<!-- -->

-   [CBOR — Concise Binary Object Representation | Overview](http://cbor.io/)
