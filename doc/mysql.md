# MySQL Compatibility

PuzzleDB supports MySQL commands based on [go-mysql](https://github.com/cybergarage/go-mysql), a database framework that makes it easy to implement MySQL compatible servers using Go.

<figure>
<img src="https://raw.githubusercontent.com/cybergarage/go-mysql/main/doc/img/framework.png" alt="framework" />
</figure>

The [go-mysql](https://github.com/cybergarage/go-mysql) framework automatically handles the MySQL protocol and system commands. Therefore, PuzzleDB achieves MySQL compatibility by implementing only simply handling DDL (Data Definition Language) and DML (Data Manipulation Language) query commands.

## Data Model

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts [MySQL data types](https://dev.mysql.com/doc/refman/8.0/en/data-types.html) into the PuzzleDB data model as follows:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">MySQL</th>
<th style="text-align: left;">PuzzleDB</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>COMPLEX</p></td>
<td style="text-align: left;"><p>map</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>ARRAY</p></td>
<td style="text-align: left;"><p>array</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>VARCHAR</p></td>
<td style="text-align: left;"><p>string</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>CHAR</p></td>
<td style="text-align: left;"><p>string</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>TINYINT</p></td>
<td style="text-align: left;"><p>tiny</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>short</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>int</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>long</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>FLOAT</p></td>
<td style="text-align: left;"><p>float32</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>DOUBLE (REAL)</p></td>
<td style="text-align: left;"><p>float64</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>DATE DATETIME</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>TIME TIMESTAMP</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>NULL</p></td>
<td style="text-align: left;"><p>null</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>BOOLEAN (TINYINT(1))</p></td>
<td style="text-align: left;"><p>bool</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>BLOB (BYTEA)</p></td>
<td style="text-align: left;"><p>[]byte</p></td>
</tr>
</tbody>
</table>

## Supported commands

PuzzleDB currently supports [MySQL statements](https://dev.mysql.com/doc/refman/8.0/en/sql-statements.html) in stages. This section describes the status of Redis command support in PuzzleDB.

### Data Definition Statements

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Statement</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>CREATE DATABASE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>CREATE TABLE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DROP DATABASE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DROP TABLE</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### Transaction Control Statements

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Statement</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>BEGIN</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>COMMIT</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>ROLLBACK</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### Data Manipulation Statements

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Statement</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DELETE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>INSERT</p></td>
<td style="text-align: left;"><p>Does not perform constraint checking on foreign keys</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"><p>Sub-queries and algebraic operations are currently not supported</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>UPDATE</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### Utility Statements

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Statement</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>USE</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>
