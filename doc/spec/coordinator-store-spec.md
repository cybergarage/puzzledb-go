# Coordinator Key-value Store

The coordinator sevice has a key-value store to commuicate with other nodes in PuzzleDB. The key-value store is a collection of key-value records, where each record is a key-value pair, consisting of a header as the key and a value as the value. The key-value store supports the following categories of key-value records:

![coordinator compo](img/coordinator_compo.png)

The coordinator service provides a distributed key-value store for PuzzleDB nodes. The key-value store is a collection of key-value records, where each record is a key-value pair, consisting of a header as the key and a value as the value.

## Key Categories

The key-value store is a collection of key-value records, where each record is a key-value pair, consisting of a header as the key. The key-value store supports the following categories of key-value records:

<table>
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Key Order</th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;">Value</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Message</p></td>
<td style="text-align: left;"><p>Header (M)</p></td>
<td style="text-align: left;"><p>Logical Clock</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Message)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>State</p></td>
<td style="text-align: left;"><p>Header (S)</p></td>
<td style="text-align: left;"><p>State Type</p></td>
<td style="text-align: left;"><p>(Key)</p></td>
<td style="text-align: left;"><p>CBOR (State)</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Job</p></td>
<td style="text-align: left;"><p>Header (J)</p></td>
<td style="text-align: left;"><p>Job ID</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Job)</p></td>
</tr>
</tbody>
</table>

The value of the coordinator store object is encoded and decoded in CBOR format as standard.

## Key Header Specification

The key header is a 2-byte header that is prepended to every key in the key-value store. The key header is reserved as follows:

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Field Name</th>
<th style="text-align: left;">Size (bits)</th>
<th style="text-align: left;">Description</th>
<th style="text-align: left;">Example Value</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Key category</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>The record key type</p></td>
<td style="text-align: left;"><p>N:Node M:Message J:Job</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Version</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The version number</p></td>
<td style="text-align: left;"><p>0:reserved 1-7</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Value type</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The record value type</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
</tbody>
</table>

The key header begins with a 1-byte identifier for the key type, enabling key type-based searching. Duplication is tolerated because a value type is reserved for each key type.

## State Objects

The coordinator service defines standard state objects to share state among the PuzzleDB nodes. The state object values are defined by category, but the standard state object header and key order are defined as follows:

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
<tr class="header">
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Key Order</th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;">Value</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Node</p></td>
<td style="text-align: left;"><p>Header (S)</p></td>
<td style="text-align: left;"><p>State Type (N)</p></td>
<td style="text-align: left;"><p>Cluster ID</p></td>
<td style="text-align: left;"><p>Node ID</p></td>
<td style="text-align: left;"><p>CBOR (Node Object)</p></td>
</tr>
</tbody>
</table>
