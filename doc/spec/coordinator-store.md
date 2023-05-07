# Coordinator Store Specification

The Coordinator service provides distributed synchronization and coordination for PuzzleDB nodes, as well as a distributed key-value store. This document describes the specifications of the information stored in the key-value store provided by the Coordinator service.

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
<td style="text-align: left;"><p>Node</p></td>
<td style="text-align: left;"><p>Header</p></td>
<td style="text-align: left;"><p>Host (IP address)</p></td>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>CBOR (State)</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Event</p></td>
<td style="text-align: left;"><p>Header</p></td>
<td style="text-align: left;"><p>Event type</p></td>
<td style="text-align: left;"><p>Tuple (Key)</p></td>
<td style="text-align: left;"><p>CBOR (Message)</p></td>
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
<td style="text-align: left;"><p>N:Node E:Event</p></td>
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
