# PostgreSQL Compatibility

PuzzleDB supports PostgreSQL commands based on [go-postgresql](https://github.com/cybergarage/go-postgresql), a database framework that makes it easy to implement PostgreSQL compatible servers using Go.

<figure>
<img src="https://raw.githubusercontent.com/cybergarage/go-postgresql/master/doc/img/framework.png" alt="framework" />
</figure>

The [go-postgresql](https://github.com/cybergarage/go-postgresql) framework automatically handles the PostgreSQL protocol and system commands. Therefore, PuzzleDB achieves PostgreSQL compatibility by implementing only simply handling DDL (Data Definition Language) and DML (Data Manipulation Language) query commands.

## Data Model

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts [PostgreSQL: Data Types](https://www.postgresql.org/docs/current/datatype.html) into the PuzzleDB data model as follows:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">PostgreSQL</th>
<th style="text-align: left;">PuzzleDB</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>TEXT</p></td>
<td style="text-align: left;"><p>string</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>VARCHAR</p></td>
<td style="text-align: left;"><p>string</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>CHAR</p></td>
<td style="text-align: left;"><p>string</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>SMALLINT</p></td>
<td style="text-align: left;"><p>int16</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>INTEGER</p></td>
<td style="text-align: left;"><p>int32</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BIGINT</p></td>
<td style="text-align: left;"><p>int64</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DECIMAL</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>NUMERIC</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>REAL</p></td>
<td style="text-align: left;"><p>float32</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>DOUBLE PRECISION</p></td>
<td style="text-align: left;"><p>float64</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>TIMESTAMP</p></td>
<td style="text-align: left;"><p>timestamp</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>NULL</p></td>
<td style="text-align: left;"><p>null</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>BOOLEAN</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Bytea</p></td>
<td style="text-align: left;"><p>[]byte</p></td>
</tr>
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
<tr>
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Statement</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>CREATE DATABASE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>CREATE TABLE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>ALTER DATABSE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>ALTER TABLE</p></td>
<td style="text-align: left;"><p>ADD CLUMN, ADD INDEX, DROP COLUMN</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DROP DATABASE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
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
<tr>
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Statement</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>BEGIN</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>COMMIT</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>ROLLBACK</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>SAVEPOINT</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>RELEASE SAVEPOINT</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>ROLLBACK TO SAVEPOINT</p></td>
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
<tr>
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Statement</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DELETE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>INSERT</p></td>
<td style="text-align: left;"><p>Does not perform constraint checking on foreign keys</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"><p>Sub-queries and algebraic operations are currently not supported</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>UPDATE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>TRUNCATE</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>VACCUM</p></td>
<td style="text-align: left;"><p>(No Current Action)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>COPY</p></td>
<td style="text-align: left;"><p>Support only text format</p></td>
</tr>
</tbody>
</table>

### Functions

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Type</th>
<th style="text-align: left;">Function</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Aggregation</p></td>
<td style="text-align: left;"><p>COUNT</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>MIN</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>MAX</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>AVG</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Mathematic</p></td>
<td style="text-align: left;"><p>ABS</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>FLOOR</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CEIL</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>Arithmetic</p></td>
<td style="text-align: left;"><p>+</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>*</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>/</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>%</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>
