Key-Value Store Specification
=============================

In PuzzleDB, both the coordinator service and the store service use key-value stores to store data. The coordinator uses a key-value store to store states, messages, jobs, and indices, while the store uses a key-value store to store to store records, schemas, and indices.

Since both services may use the same key-value store implementation, the key-value store specifications are designed not to affect each other. This document provides a list of these specifications.

Coordinator Key-Value Specification
-----------------------------------

<table><colgroup><col style="width: 20%" /><col style="width: 20%" /><col style="width: 20%" /><col style="width: 20%" /><col style="width: 20%" /></colgroup><thead><tr class="header"><th>Category</th><th>Key Order</th><th></th><th></th><th>Value</th></tr></thead><tbody><tr class="odd"><td></td><td><p>0</p></td><td><p>1</p></td><td><p>2</p></td><td></td></tr><tr class="even"><td><p>Message</p></td><td><p>Header (M)</p></td><td><p>Logical Clock</p></td><td><p>-</p></td><td><p>CBOR (Message)</p></td></tr><tr class="odd"><td><p>State</p></td><td><p>Header (S)</p></td><td><p>State Type</p></td><td><p>(Key)</p></td><td><p>CBOR (State)</p></td></tr><tr class="even"><td><p>Job</p></td><td><p>Header (J)</p></td><td><p>Job ID</p></td><td><p>-</p></td><td><p>CBOR (Job)</p></td></tr></tbody></table>

Store Key-Value Specification
-----------------------------

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Category</th><th>Key Order</th><th></th><th></th><th></th><th>Value</th></tr></thead><tbody><tr class="odd"><td></td><td><p>0</p></td><td><p>1</p></td><td><p>2</p></td><td><p>3</p></td><td></td></tr><tr class="even"><td><p>Database</p></td><td><p>Header (D)</p></td><td><p>Database</p></td><td><p>-</p></td><td><p>-</p></td><td><p>CBOR (Options)</p></td></tr><tr class="odd"><td><p>Collection</p></td><td><p>Header ©</p></td><td><p>Database</p></td><td><p>Collection</p></td><td><p>-</p></td><td><p>CBOR (Schema)</p></td></tr><tr class="even"><td><p>Object</p></td><td><p>Header (O)</p></td><td><p>Database</p></td><td><p>Collection</p></td><td><p>Element</p></td><td><p>CBOR (Object)</p></td></tr><tr class="odd"><td><p>Index</p></td><td><p>Header (I)</p></td><td><p>Database</p></td><td><p>Collection</p></td><td><p>Element</p></td><td><p>Tuple (Key)</p></td></tr></tbody></table>
