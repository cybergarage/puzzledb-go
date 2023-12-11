Coordinator Key-value Store
===========================

The coordinator sevice has a key-value store to commuicate with other nodes in PuzzleDB. The key-value store is a collection of key-value records, where each record is a key-value pair, consisting of a header as the key and a value as the value. The key-value store supports the following categories of key-value records:

![coordinator compo](img/coordinator_compo.png)

The coordinator service provides a distributed key-value store for PuzzleDB nodes. The key-value store is a collection of key-value records, where each record is a key-value pair, consisting of a header as the key and a value as the value.

Key Categories
--------------

The key-value store is a collection of key-value records, where each record is a key-value pair, consisting of a header as the key. The key-value store supports the following categories of key-value records:

<table><colgroup><col style="width: 20%" /><col style="width: 20%" /><col style="width: 20%" /><col style="width: 20%" /><col style="width: 20%" /></colgroup><thead><tr class="header"><th>Category</th><th>Key Order</th><th></th><th></th><th>Value</th></tr></thead><tbody><tr class="odd"><td></td><td><p>0</p></td><td><p>1</p></td><td><p>2</p></td><td></td></tr><tr class="even"><td><p>Message</p></td><td><p>Header (M)</p></td><td><p>Logical Clock</p></td><td><p>-</p></td><td><p>CBOR (Message)</p></td></tr><tr class="odd"><td><p>State</p></td><td><p>Header (S)</p></td><td><p>State Type</p></td><td><p>(Key)</p></td><td><p>CBOR (State)</p></td></tr><tr class="even"><td><p>Job</p></td><td><p>Header (J)</p></td><td><p>Job ID</p></td><td><p>-</p></td><td><p>CBOR (Job)</p></td></tr></tbody></table>

The value of the coordinator store object is encoded and decoded in CBOR format as standard.

Key Header Specification
------------------------

The key header is a 2-byte header that is prepended to every key in the key-value store. The key header is reserved as follows:

<table><colgroup><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /></colgroup><thead><tr class="header"><th>Field Name</th><th>Size (bits)</th><th>Description</th><th>Example Value</th></tr></thead><tbody><tr class="odd"><td><p>Key category</p></td><td><p>8</p></td><td><p>The record key type</p></td><td><p>N:Node M:Message J:Job</p></td></tr><tr class="even"><td><p>Version</p></td><td><p>4</p></td><td><p>The version number</p></td><td><p>0:reserved 1-7</p></td></tr><tr class="odd"><td><p>Value type</p></td><td><p>4</p></td><td><p>The record value type</p></td><td><p>0:reserved 1:CBOR</p></td></tr></tbody></table>

The key header begins with a 1-byte identifier for the key type, enabling key type-based searching. Duplication is tolerated because a value type is reserved for each key type.

State Objects
-------------

The coordinator service defines standard state objects to share state among the PuzzleDB nodes. The state object values are defined by category, but the standard state object header and key order are defined as follows:

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Category</th><th>Key Order</th><th></th><th></th><th></th><th>Value</th></tr></thead><tbody><tr class="odd"><td></td><td><p>0</p></td><td><p>1</p></td><td><p>2</p></td><td><p>3</p></td><td></td></tr><tr class="even"><td><p>Node</p></td><td><p>Header (S)</p></td><td><p>State Type (N)</p></td><td><p>Cluster ID</p></td><td><p>Node ID</p></td><td><p>CBOR (Node Object)</p></td></tr></tbody></table>
