# Key-Value Store Specification

    In PuzzleDB, not only are records represented as key-value pairs, but schemas and indices are also represented as key-value data. This section describes the format of the key-value data used in PuzzleDB.

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
<th style="text-align: left;">Type</th>
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
<td style="text-align: left;"><p>1</p></td>
<td style="text-align: left;"><p>2</p></td>
<td style="text-align: left;"><p>3</p></td>
<td style="text-align: left;"><p>4</p></td>
<td style="text-align: left;"></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Object</p></td>
<td style="text-align: left;"><p>Header</p></td>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>Element</p></td>
<td style="text-align: left;"><p>Object</p></td>
</tr>
<tr class="odd">
<td style="text-align: left;"><p>Index</p></td>
<td style="text-align: left;"><p>Header</p></td>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>Element</p></td>
<td style="text-align: left;"><p>Key</p></td>
</tr>
<tr class="even">
<td style="text-align: left;"><p>Schema</p></td>
<td style="text-align: left;"><p>Header</p></td>
<td style="text-align: left;"><p>Database</p></td>
<td style="text-align: left;"><p>Collection</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>CBOR</p></td>
</tr>
</tbody>
</table>
