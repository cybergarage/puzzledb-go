<figure>
<img src="img/logo.png" alt="logo" />
</figure>

# Redis

PuzzleDB supports Redis commands based on [go-redis](https://github.com/cybergarage/go-redis), a database framework that makes it easy to implement Redis compatible servers using Go.

## Data Model

PuzzleDB is a multi-data model database and the core data model is a document model; PuzzleDB converts Redis data model the PuzzleDB data model as follows:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Redis</th>
<th style="text-align: left;">PuzzleDB</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Hash</p></td>
<td style="text-align: left;"><p>map</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>List</p></td>
<td style="text-align: left;"><p>array</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Sets</p></td>
<td style="text-align: left;"><p>array</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Sorted Sets</p></td>
<td style="text-align: left;"><p>array</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>String</p></td>
<td style="text-align: left;"><p>[]byte</p></td>
</tr>
</tbody>
</table>

## Supported Commands

PuzzleDB currently supports [Redis commands](https://redis.io/commands/) in stages. This section describes the status of Redis command support in PuzzleDB.

### Connection Commands

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Command</th>
<th style="text-align: left;">Redis Version</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>ECHO</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>PING</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>QUIT</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SELECT</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### Generic Commands

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Command</th>
<th style="text-align: left;">Redis Version</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>COPY</p></td>
<td style="text-align: left;"><p>6.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DEL</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>DUMP</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>EXISTS</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>EXPIRE</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>EXPIREAT</p></td>
<td style="text-align: left;"><p>1.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>EXPIRETIME</p></td>
<td style="text-align: left;"><p>7.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>KEYS</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>MOVE</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>MIGRATE</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>OBJECT ENCODING</p></td>
<td style="text-align: left;"><p>2.2.3</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>OBJECT FREQ</p></td>
<td style="text-align: left;"><p>4.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>OBJECT HELP</p></td>
<td style="text-align: left;"><p>6.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>OBJECT IDLETIME</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>OBJECT REFCOUNT</p></td>
<td style="text-align: left;"><p>2.2.3</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>PERSIST</p></td>
<td style="text-align: left;"><p>2.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>PEXPIRE</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>PEXPIREAT</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>PEXPIRETIME</p></td>
<td style="text-align: left;"><p>7.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>PTTL</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>RANDOMKEY</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>RENAME</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>RENAMENX</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>RESTORE</p></td>
<td style="text-align: left;"><p>2.8.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>SCAN</p></td>
<td style="text-align: left;"><p>2.8.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>SORT</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>SORT_RO</p></td>
<td style="text-align: left;"><p>7.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>TOUCH</p></td>
<td style="text-align: left;"><p>3.2.1</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>TTL</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>TYPE</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>UNLINK</p></td>
<td style="text-align: left;"><p>4.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>WAIT</p></td>
<td style="text-align: left;"><p>3.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### String Commands

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Support</th>
<th style="text-align: left;">Command</th>
<th style="text-align: left;">Redis Version</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>APPEND</p></td>
<td style="text-align: left;"><p>2.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DECR</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>DECRBY</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>GET</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>GETDEL</p></td>
<td style="text-align: left;"><p>6.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>GETEX</p></td>
<td style="text-align: left;"><p>6.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>GETRANGE</p></td>
<td style="text-align: left;"><p>2.4.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>GETSET</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>INCR</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>INCRBY</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>INCRBYFLOAT</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>LCS</p></td>
<td style="text-align: left;"><p>7.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>MGET</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>MSET</p></td>
<td style="text-align: left;"><p>1.0.1</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>MSETNX</p></td>
<td style="text-align: left;"><p>1.0.1</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>PSETNX</p></td>
<td style="text-align: left;"><p>2.6.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SET</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"><p>Any options are not supported yet</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>SETEX</p></td>
<td style="text-align: left;"><p>2.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SETNX</p></td>
<td style="text-align: left;"><p>2.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>SERANGE</p></td>
<td style="text-align: left;"><p>2.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>STRLEN</p></td>
<td style="text-align: left;"><p>2.2.0</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>SUBSTR</p></td>
<td style="text-align: left;"><p>1.0.0</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>
