<figure>
<img src="img/logo.png" alt="logo" />
</figure>

# MySQL

PuzzleDB supports MySQL commands based on [go-mysql](https://github.com/cybergarage/go-mysql), a database framework that makes it easy to implement MySQL compatible servers using Go.

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
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"></td>
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
