Key-Value Store Specification
=============================

In PuzzleDB, records, schemas, and indices are all represented as key-value pairs. This section explains the format of the key-value data used in PuzzleDB.

Key Categories
--------------

The key-value store is a collection of key-value records, where each record is a key-value pair, consisting of a header as the key. The key-value store supports the following categories of key-value records:

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Category</th><th>Key Order</th><th></th><th></th><th></th><th>Value</th></tr></thead><tbody><tr class="odd"><td></td><td><p>0</p></td><td><p>1</p></td><td><p>2</p></td><td><p>3</p></td><td></td></tr><tr class="even"><td><p>Database</p></td><td><p>Header (D)</p></td><td><p>Database</p></td><td><p>-</p></td><td><p>-</p></td><td><p>CBOR (Options)</p></td></tr><tr class="odd"><td><p>Collection</p></td><td><p>Header ©</p></td><td><p>Database</p></td><td><p>Collection</p></td><td><p>-</p></td><td><p>CBOR (Schema)</p></td></tr><tr class="even"><td><p>Object</p></td><td><p>Header (O)</p></td><td><p>Database</p></td><td><p>Collection</p></td><td><p>Element</p></td><td><p>CBOR (Object)</p></td></tr><tr class="odd"><td><p>Index</p></td><td><p>Header (I)</p></td><td><p>Database</p></td><td><p>Collection</p></td><td><p>Element</p></td><td><p>Tuple (Key)</p></td></tr></tbody></table>

Key Header Specification
------------------------

The key header is a 2-byte header that is prepended to every key in the key-value store. The key header is reserved as follows:

<table><colgroup><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /></colgroup><thead><tr class="header"><th>Field Name</th><th>Size (bits)</th><th>Description</th><th>Example Value</th></tr></thead><tbody><tr class="odd"><td><p>Key category</p></td><td><p>8</p></td><td><p>The record key type</p></td><td><p>D:Database C:Collection O:Document I:Index</p></td></tr><tr class="even"><td><p>Version</p></td><td><p>4</p></td><td><p>The version number</p></td><td><p>0:reserved 1-7</p></td></tr><tr class="odd"><td><p>Value type</p></td><td><p>4</p></td><td><p>The record value type</p></td><td><p>0:reserved 1:CBOR 1:PRIMARY 2:SECONDARY</p></td></tr></tbody></table>

The key header begins with a 1-byte identifier for the key type, enabling key type-based searching. Duplication is tolerated because a value type is reserved for each key type.
