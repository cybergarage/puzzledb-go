# MySQL Compatibility

PuzzleDB supports MySQL commands based on [go-postgresql](https://github.com/cybergarage/go-postgresql), a database framework that makes it easy to implement MySQL compatible servers using Go.

<figure>
<img src="https://raw.githubusercontent.com/cybergarage/go-postgresql/master/doc/img/framework.png" alt="framework" />
</figure>

The [go-postgresql](https://github.com/cybergarage/go-postgresql) framework automatically handles the MySQL protocol and system commands. Therefore, PuzzleDB achieves PostgreSQL compatibility by implementing only simply handling DDL (Data Definition Language) and DML (Data Manipulation Language) query commands.

## Data Model

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts [PostgreSQL: Data Types](https://www.postgresql.org/docs/current/datatype.html) into the PuzzleDB data model as follows:

<table>
<colgroup>
<col style="width: 100%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Unresolved directive in postgresql.adoc - include::model/postgresql.csv[]</th>
</tr>
</thead>
<tbody>
</tbody>
</table>

## Supported commands

PuzzleDB currently supports [PostgreSQL: Basic Statements](https://www.postgresql.org/docs/current/plpgsql-statements.html) in stages. This section describes the status of Redis command support in PuzzleDB.

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
