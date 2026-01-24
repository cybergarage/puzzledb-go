# Data Model

PuzzleDB is a multi‑model database. The core logical model is a document model layered atop an ordered key‑value store. All objects (data, schema, indexes) are represented as documents and ultimately persisted as key‑value pairs.

<figure>
<img src="img/storage.png" alt="storage" />
</figure>

PuzzleDB defines a storage plugin interface enabling use of local in‑memory stores (e.g. memdb) or large distributed stores (e.g. FoundationDB, TiKV).

## Document Model

The document model must be expressive (JSON / BSON level) similar to ARS (Atom‑Record‑Sequence) in CosmosDB.

PuzzleDB maps external data models (relational, document, key‑value) into its internal document representation:

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
<th style="text-align: left;">Type</th>
<th style="text-align: left;">PuzzleDB</th>
<th style="text-align: left;">Redis</th>
<th style="text-align: left;">MongoDB</th>
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">PostgreSQL</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>map</p></td>
<td style="text-align: left;"><p>Hash</p></td>
<td style="text-align: left;"><p>Object</p></td>
<td style="text-align: left;"><p>COMPLEX</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>array</p></td>
<td style="text-align: left;"><p>List</p></td>
<td style="text-align: left;"><p>Array</p></td>
<td style="text-align: left;"><p>ARRAY</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Sets</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Sorted Sets</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>string</p></td>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>TEXT</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>VARCHAR</p></td>
<td style="text-align: left;"><p>VARCHAR</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CHAR</p></td>
<td style="text-align: left;"><p>CHAR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Integer</p></td>
<td style="text-align: left;"><p>tiny</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>TINYINT</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>short</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>SMALLINT</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>32-bit integer</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>INTEGER</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>long</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>64-bit integer</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>BIGINT</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Real</p></td>
<td style="text-align: left;"><p>float32</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>32-bit IEEE-754</p></td>
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>REAL</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>float64</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>64-bit IEEE-754</p></td>
<td style="text-align: left;"><p>DOUBLE (REAL)</p></td>
<td style="text-align: left;"><p>DOUBLE (REAL)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Time</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Date</p></td>
<td style="text-align: left;"><p>DATE DATETIME</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Timestamp</p></td>
<td style="text-align: left;"><p>TIME TIMESTAMP</p></td>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Special</p></td>
<td style="text-align: left;"><p>null</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Null</p></td>
<td style="text-align: left;"><p>NULL</p></td>
<td style="text-align: left;"><p>NULL</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>bool</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Boolean</p></td>
<td style="text-align: left;"><p>BOOLEAN (TINYINT(1))</p></td>
<td style="text-align: left;"><p>BOOLEAN</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>[]byte</p></td>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>Binary data</p></td>
<td style="text-align: left;"><p>BLOB (BYTEA)</p></td>
<td style="text-align: left;"><p>BINARY</p></td>
</tr>
</tbody>
</table>

### See also

- [plugins.query.sql.NewDocumentElementTypeFrom()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/query/sql/type.go)

- [plugins.query.mongo.BSONEncoder::EncodeBSON()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/query/mongo/encoder.go)

## Key-Value Object Model

The ordered key‑value model underpins all persisted objects. It provides flexible, scalable storage and efficient range operations.

All higher‑level objects are encoded as documents and stored as key‑value entries.

# Key Object

In PuzzleDB, records, schemas, and indices are all represented as key-value pairs. This section describes the format of the key object in detail.

## Key Header Specification

Every key object includes a header that specifies the key category, version, and the type of stored value. The key header is a 2-byte field prepended to every key in the key-value store. It is structured as follows:

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Field Name</th>
<th style="text-align: left;">Size (bits)</th>
<th style="text-align: left;">Description</th>
<th style="text-align: left;">Example Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Category</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>The record object category</p></td>
<td style="text-align: left;"><p>D:Database C:Collection O:Document I:Index</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Format</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The record object format</p></td>
<td style="text-align: left;"><p>(Defined for each key category)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Version</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The record format version</p></td>
<td style="text-align: left;"><p>0:reserved 1:Current</p></td>
</tr>
</tbody>
</table>

Key headers start with a one-byte identifier that indicates the type of key, enabling efficient searches based on key type. Currently, the version is fixed at `1`. The value type is specified individually for each key category. The values are specified as follows:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Value Types</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Index</p></td>
<td style="text-align: left;"><p>0:reserved 1:Primary 2:Secondary</p></td>
</tr>
</tbody>
</table>

## Key Categories

The key-value store consists of key-value records, where each record is defined by a key-value pair and includes a header as part of the key. The store supports the following categories of key-value records:

<table style="width:100%;">
<colgroup>
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Key Order</th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;">Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>5</p></td>
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>Header (D)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Options)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>Header (C)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Schema)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>Header (O)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>Primary Element Value</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Object)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Index</p></td>
<td style="text-align: left;"><p>Header (I)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>Secondary Element Name</p></td>
<td style="text-align: left;"><p>Secondary Element Value</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
</tbody>
</table>

Primary keys and secondary indices may comprise one or more columns. Although omitted in the table above, the combination of the element name and value for both objects and indices is repeated based on the index format. Additionally, since the primary key is stored in the key section of an index, the value section remains empty.

### See also

- [plugins.coder.key.tuple.Coder::EncodeKey()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/coder/key/tuple/coder.go)

### Document (Value) Object

The document abstraction is realized via coder plugins layered on key‑value storage. The default coder uses CBOR (Concise Binary Object Representation).

Documents are encoded using the active coder (CBOR by default) and persisted as key‑value entries. The relationship between the document model and CBOR encoding is shown below.

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Type</th>
<th style="text-align: left;">PuzzleDB</th>
<th style="text-align: left;">CBOR</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>map</p></td>
<td style="text-align: left;"><p>5 (map)</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>array</p></td>
<td style="text-align: left;"><p>4 (array)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>string</p></td>
<td style="text-align: left;"><p>3 (text string)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Integer</p></td>
<td style="text-align: left;"><p>tiny</p></td>
<td style="text-align: left;"><p>tiny</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>short</p></td>
<td style="text-align: left;"><p>short</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>int</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>long</p></td>
<td style="text-align: left;"><p>long</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Real</p></td>
<td style="text-align: left;"><p>float32</p></td>
<td style="text-align: left;"><p>7 (floating-point) 26</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>float64</p></td>
<td style="text-align: left;"><p>7 (floating-point) 27</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Time</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
<td style="text-align: left;"><p>6 (tag) 0</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Special</p></td>
<td style="text-align: left;"><p>null</p></td>
<td style="text-align: left;"><p>null</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>bool</p></td>
<td style="text-align: left;"><p>bool</p></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>[]byte</p></td>
<td style="text-align: left;"><p>binary</p></td>
</tr>
</tbody>
</table>

#### See also

- [plugins.coder.document.cbor.Coder::EncodeDocument()](https://github.com/cybergarage/puzzledb-go/blob/main/puzzledb/plugins/coder/document/cbor/coder.go)

### References

- [A technical overview of Azure Cosmos DB | Azure Blog and Updates | Microsoft Azure](https://azure.microsoft.com/en-gb/blog/a-technical-overview-of-azure-cosmos-db/)

  - [Azure Cosmos DB conceptual whitepapers](https://learn.microsoft.com/en-us/azure/cosmos-db/whitepapers)

  - [Schema-Agnostic Indexing with Azure DocumentDB](https://www.vldb.org/pvldb/vol8/p1668-shukla.pdf)

<!-- -->

- [CBOR — Concise Binary Object Representation | Overview](http://cbor.io/)
