PostgreSQL Compatibility
========================

PuzzleDB supports PostgreSQL commands based on [go-postgresql](https://github.com/cybergarage/go-postgresql), a database framework that makes it easy to implement PostgreSQL compatible servers using Go.

![framework](https://raw.githubusercontent.com/cybergarage/go-postgresql/master/doc/img/framework.png)

The [go-postgresql](https://github.com/cybergarage/go-postgresql) framework automatically handles the PostgreSQL protocol and system commands. Therefore, PuzzleDB achieves PostgreSQL compatibility by implementing only simply handling DDL (Data Definition Language) and DML (Data Manipulation Language) query commands.

Data Model
----------

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts [PostgreSQL: Data Types](https://www.postgresql.org/docs/current/datatype.html) into the PuzzleDB data model as follows:

<table><colgroup><col style="width: 50%" /><col style="width: 50%" /></colgroup><thead><tr class="header"><th>PostgreSQL</th><th>PuzzleDB</th></tr></thead><tbody><tr class="odd"><td><p>TEXT</p></td><td><p>string</p></td></tr><tr class="even"><td><p>VARCHAR</p></td><td><p>string</p></td></tr><tr class="odd"><td><p>CHAR</p></td><td><p>string</p></td></tr><tr class="even"><td><p>SMALLINT</p></td><td><p>int16</p></td></tr><tr class="odd"><td><p>INTEGER</p></td><td><p>int32</p></td></tr><tr class="even"><td><p>BIGINT</p></td><td><p>int64</p></td></tr><tr class="odd"><td><p>DECIMAL</p></td><td><p>-</p></td></tr><tr class="even"><td><p>NUMERIC</p></td><td><p>-</p></td></tr><tr class="odd"><td><p>REAL</p></td><td><p>float32</p></td></tr><tr class="even"><td><p>DOUBLE PRECISION</p></td><td><p>float64</p></td></tr><tr class="odd"><td><p>TIMESTAMP</p></td><td><p>timestamp</p></td></tr><tr class="even"><td><p>NULL</p></td><td><p>null</p></td></tr><tr class="odd"><td><p>BOOLEAN</p></td><td><p>-</p></td></tr><tr class="even"><td><p>Bytea</p></td><td><p>[]byte</p></td></tr></tbody></table>

Supported commands
------------------

PuzzleDB currently supports [PostgreSQL: Basic Statements](https://www.postgresql.org/docs/current/plpgsql-statements.html) in stages. This section describes the status of Redis command support in PuzzleDB.

### Data Definition Statements

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Statement</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>CREATE DATABASE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>CREATE TABLE</p></td><td></td></tr><tr class="odd"><td><p>-</p></td><td><p>ALTER DATABSE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>ALTER TABLE</p></td><td><p>ADD CLUMN, ADD INDEX, DROP COLUMN</p></td></tr><tr class="odd"><td><p>O</p></td><td><p>DROP DATABASE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>DROP TABLE</p></td><td></td></tr></tbody></table>

### Transaction Control Statements

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Statement</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>BEGIN</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>COMMIT</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td><p>ROLLBACK</p></td><td></td></tr><tr class="even"><td><p>-</p></td><td><p>SAVEPOINT</p></td><td></td></tr><tr class="odd"><td><p>-</p></td><td><p>RELEASE SAVEPOINT</p></td><td></td></tr><tr class="even"><td><p>-</p></td><td><p>ROLLBACK TO SAVEPOINT</p></td><td></td></tr></tbody></table>

### Data Manipulation Statements

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Statement</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>DELETE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>INSERT</p></td><td><p>Does not perform constraint checking on foreign keys</p></td></tr><tr class="odd"><td><p>O</p></td><td><p>SELECT</p></td><td><p>Sub-queries and algebraic operations are currently not supported</p></td></tr><tr class="even"><td><p>O</p></td><td><p>UPDATE</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td><p>TRUNCATE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>VACCUM</p></td><td><p>(No Current Action)</p></td></tr><tr class="odd"><td><p>O</p></td><td><p>COPY</p></td><td><p>Support only text format</p></td></tr></tbody></table>

### Functions

<table><colgroup><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /></colgroup><thead><tr class="header"><th>Support</th><th>Type</th><th>Function</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>Aggregation</p></td><td><p>COUNT</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td></td><td><p>MIN</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td></td><td><p>MAX</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td></td><td><p>AVG</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td><p>Mathematic</p></td><td><p>ABS</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td></td><td><p>FLOOR</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td></td><td><p>CEIL</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>Arithmetic</p></td><td><p>+</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td></td><td><p>-</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td></td><td><p>*</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td></td><td><p>/</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td></td><td><p>%</p></td><td></td></tr></tbody></table>
