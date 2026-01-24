# Key Object

In PuzzleDB, records, schemas, and indices are all represented as key-value pairs. This section describes the format of the key object in detail.

## Key Header Specification

Every key object includes a header that specifies the key category, version, and the type of stored value. The key header is a 2-byte field prepended to every key in the key-value store. It is structured as follows:

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Field Name</th>
<th style="text-align: left;">Size (bits)</th>
<th style="text-align: left;">Description</th>
<th style="text-align: left;">Example Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Category</p></td>
<td style="text-align: left;"><p>8</p></td>
<td style="text-align: left;"><p>The record object category</p></td>
<td style="text-align: left;"><p>D:Database C:Collection O:Document I:Index</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Format</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The record object format</p></td>
<td style="text-align: left;"><p>(Defined for each key category)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Version</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>The record format version</p></td>
<td style="text-align: left;"><p>0:reserved 1:Current</p></td>
</tr>
</tbody>
</table>

Key headers start with a one-byte identifier that indicates the type of key, enabling efficient searches based on key type. Currently, the version is fixed at `1`. The value type is specified individually for each key category. The values are specified as follows:

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Value Types</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>0:reserved 1:CBOR</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Index</p></td>
<td style="text-align: left;"><p>0:reserved 1:Primary 2:Secondary</p></td>
</tr>
</tbody>
</table>

## Key Categories

The key-value store consists of key-value records, where each record is defined by a key-value pair and includes a header as part of the key. The store supports the following categories of key-value records:

<table style="width:100%;">
<colgroup>
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
<col style="width: 11%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Category</th>
<th style="text-align: left;">Key Order</th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;"></th>
<th style="text-align: left;">Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>0</p></td>
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"><p>5</p></td>
<td style="text-align: left;"><p>6</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>Header (D)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Options)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>Header (C)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Schema)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Document</p></td>
<td style="text-align: left;"><p>Header (O)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>Primary Element Value</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR (Object)</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Index</p></td>
<td style="text-align: left;"><p>Header (I)</p></td>
<td style="text-align: left;"><p>Database Name</p></td>
<td style="text-align: left;"><p>Collection Name</p></td>
<td style="text-align: left;"><p>Secondary Element Name</p></td>
<td style="text-align: left;"><p>Secondary Element Value</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>Primary Element Name</p></td>
<td style="text-align: left;"><p>-</p></td>
</tr>
</tbody>
</table>

Primary keys and secondary indices may comprise one or more columns. Although omitted in the table above, the combination of the element name and value for both objects and indices is repeated based on the index format. Additionally, since the primary key is stored in the key section of an index, the value section remains empty.
