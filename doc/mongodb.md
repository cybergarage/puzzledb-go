<figure>
<img src="img/logo.png" alt="logo" />
</figure>

# MongoDB

PuzzleDB supports MongoDB API based on [go-mongo](https://github.com/cybergarage/go-mongo), a database framework that makes it easy to implement MongoDB compatible servers using Go.

# Supported Commands

PuzzleDB currently supports [MongoDB database commands](https://www.mongodb.com/docs/manual/reference/command/) in stages. This section describes the status of MongoDB command support in PuzzleDB.

## User Commands

### Diagnostic Commands

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Supported</th>
<th style="text-align: left;">Diagnostic Command</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>buildInfo</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>

### Replication Commands

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Supported</th>
<th style="text-align: left;">Replication Command</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>isMaster</p></td>
<td style="text-align: left;"><p>Always returns true</p></td>
</tr>
</tbody>
</table>

### Query and Write Operation Commands

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Supported</th>
<th style="text-align: left;">Query Command</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>delete</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>find</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>findAndModify</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>getMore</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>insert</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>resetError</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>O</p></td>
<td style="text-align: left;"><p>update</p></td>
<td style="text-align: left;"></td>
</tr>
</tbody>
</table>
