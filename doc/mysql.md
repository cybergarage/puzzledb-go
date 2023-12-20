MySQL Compatibility
===================

PuzzleDB supports MySQL commands based on [go-mysql](https://github.com/cybergarage/go-mysql), a database framework that makes it easy to implement MySQL compatible servers using Go.

![framework](https://raw.githubusercontent.com/cybergarage/go-mysql/main/doc/img/framework.png)

The [go-mysql](https://github.com/cybergarage/go-mysql) framework automatically handles the MySQL protocol and system commands. Therefore, PuzzleDB achieves MySQL compatibility by implementing only simply handling DDL (Data Definition Language) and DML (Data Manipulation Language) query commands.

Data Model
----------

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts [MySQL data types](https://dev.mysql.com/doc/refman/8.0/en/data-types.html) into the PuzzleDB data model as follows:

<table><colgroup><col style="width: 50%" /><col style="width: 50%" /></colgroup><thead><tr class="header"><th>MySQL</th><th>PuzzleDB</th></tr></thead><tbody><tr class="odd"><td><p>COMPLEX</p></td><td><p>map</p></td></tr><tr class="even"><td><p>ARRAY</p></td><td><p>array</p></td></tr><tr class="odd"><td><p>VARCHAR</p></td><td><p>string</p></td></tr><tr class="even"><td><p>CHAR</p></td><td><p>string</p></td></tr><tr class="odd"><td><p>TINYINT</p></td><td><p>tiny</p></td></tr><tr class="even"><td><p>SMALLINT</p></td><td><p>short</p></td></tr><tr class="odd"><td><p>INTEGER</p></td><td><p>int</p></td></tr><tr class="even"><td><p>BIGINT</p></td><td><p>long</p></td></tr><tr class="odd"><td><p>FLOAT</p></td><td><p>float32</p></td></tr><tr class="even"><td><p>DOUBLE (REAL)</p></td><td><p>float64</p></td></tr><tr class="odd"><td><p>DATE DATETIME</p></td><td><p>time.Time</p></td></tr><tr class="even"><td><p>TIME TIMESTAMP</p></td><td><p>time.Time</p></td></tr><tr class="odd"><td><p>NULL</p></td><td><p>null</p></td></tr><tr class="even"><td><p>BOOLEAN (TINYINT(1))</p></td><td><p>bool</p></td></tr><tr class="odd"><td><p>BLOB (BYTEA)</p></td><td><p>[]byte</p></td></tr></tbody></table>

Supported commands
------------------

PuzzleDB currently supports [MySQL statements](https://dev.mysql.com/doc/refman/8.0/en/sql-statements.html) in stages. This section describes the status of Redis command support in PuzzleDB.

### Data Definition Statements

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Statement</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>CREATE DATABASE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>CREATE TABLE</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td><p>DROP DATABASE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>DROP TABLE</p></td><td></td></tr></tbody></table>

### Transaction Control Statements

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Statement</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>BEGIN</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>COMMIT</p></td><td></td></tr><tr class="odd"><td><p>O</p></td><td><p>ROLLBACK</p></td><td></td></tr></tbody></table>

### Data Manipulation Statements

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Statement</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>DELETE</p></td><td></td></tr><tr class="even"><td><p>O</p></td><td><p>INSERT</p></td><td><p>Does not perform constraint checking on foreign keys</p></td></tr><tr class="odd"><td><p>O</p></td><td><p>SELECT</p></td><td><p>Sub-queries and algebraic operations are currently not supported</p></td></tr><tr class="even"><td><p>O</p></td><td><p>UPDATE</p></td><td></td></tr></tbody></table>

### Utility Statements

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Support</th><th>Statement</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>O</p></td><td><p>USE</p></td><td></td></tr></tbody></table>
