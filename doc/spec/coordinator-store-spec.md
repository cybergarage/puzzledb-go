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

## Key Header Specification

The key header is a 2-byte header that is prepended to every key in the key-value store. The key header is defined as follows:

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

The key header begins with a 1-byte identifier for the key type, enabling key type-based searching. Duplication is tolerated because a value type is defined for each key type.

## Message Objects

The coordinator service defines standard message objects for communication between PuzzleDB nodes. The standard message object is defined as follows:

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Field</th>
<th style="text-align: left;">Type</th>
<th style="text-align: left;">Value</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>ID</p></td>
<td style="text-align: left;"><p>UUID</p></td>
<td style="text-align: left;"><p>Destination node ID</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Host</p></td>
<td style="text-align: left;"><p>string</p></td>
<td style="text-align: left;"><p>Destination host name</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Clock</p></td>
<td style="text-align: left;"><p>uint64</p></td>
<td style="text-align: left;"><p>Destination logical clock</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Category</p></td>
<td style="text-align: left;"><p>byte</p></td>
<td style="text-align: left;"><p>Message category</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Type</p></td>
<td style="text-align: left;"><p>byte</p></td>
<td style="text-align: left;"><p>Message type</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Object</p></td>
<td style="text-align: left;"><p>[]byte</p></td>
<td style="text-align: left;"><p>Message object (CBOR)</p></td>
</tr>
</tbody>
</table>

The coordinator service defines standard message category and type ot the message objects too. The standard message category and type are defined as follows:

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr class="header">
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Type</th>
<th style="text-align: left;">Occurrence Condition</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td style="text-align: left;"><p>Object (O)</p></td>
<td style="text-align: left;"><p>Created ©</p></td>
<td style="text-align: left;"><p>Object created</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Update (U)</p></td>
<td style="text-align: left;"><p>Object updated</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Delete (D)</p></td>
<td style="text-align: left;"><p>Object deleted</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Schema (S)</p></td>
<td style="text-align: left;"><p>Created ©</p></td>
<td style="text-align: left;"><p>Schema created</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Update (U)</p></td>
<td style="text-align: left;"><p>Shcema updated</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Delete (D)</p></td>
<td style="text-align: left;"><p>Schema deleted</p></td>
</tr>
</tbody>
</table>

## State Objects

The coordinator service defines standard state objects to share state among the PuzzleDB nodes. The standard state object is defined as follows:

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
