Coordinator Specification
=========================

The coordinator service provides distributed synchronization and coordination for PuzzleDB nodes, as well as a distributed key-value store. This document describes the specifications of the information stored in the key-value store provided by the coordinator service.

Coordinator Messaging
---------------------

The message box is used to store messages sent between PuzzleDB nodes in the cluster to notify any node and store status changes using the message key-value store.

### Message Clock

Logical clocks, such as the Lamport Clock, are important in distributed systems because they can order events across different nodes; in PuzzleDB, to manage the message clock for the coordinator service, the conceptual Lamport Clock algorithm:

![coordinator clock](img/coordinator_clock.png)

In practice, the coordinator node acts as a virtual message relay node between PuzzleDB nodes. The coordinator service uses the message clock to provide the total order of messages at all nodes in the system. To manage the message clock, PuzzleDB uses the Lamport Clock algorithm, which assigns a unique timestamp to each message sent by a node.

![coordinator message clock](img/coordinator_message_clock.png)

When sending a message, a PuzzleDB node obtains the latest logical clock from the coordinator node and uses it to timestamp the message. When a PuzzleDB node receives or retrieves a message from a coordinator node, it updates its own logical clock with the logical clock included in the message.

Message Object
--------------

The message object is encoded as a CBOR object and stored as the value of the message key-value record.

### Message Object Header

The message object has the following required header fields:

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Field Name</th><th>Data Type</th><th>Description</th></tr></thead><tbody><tr class="odd"><td><p>type</p></td><td><p>int</p></td><td><p>Message type</p></td></tr><tr class="even"><td><p>timestamp</p></td><td><p>time.Time</p></td><td><p>Generated phisical time</p></td></tr></tbody></table>

The type field is used to identify messages and the timestamp field is used to discard old messages.

### Message Object Value

The message value object is encoded as a CBOR object and stored as the value of the message key-value record to the coordinator service. The message object has a message type, an event type and a CBOR encoded message object, and the standard message object is defined as follows:

<table><colgroup><col style="width: 33%" /><col style="width: 33%" /><col style="width: 33%" /></colgroup><thead><tr class="header"><th>Field</th><th>Type</th><th>Value</th></tr></thead><tbody><tr class="odd"><td><p>ID</p></td><td><p>UUID</p></td><td><p>Destination node ID</p></td></tr><tr class="even"><td><p>Host</p></td><td><p>string</p></td><td><p>Destination host name</p></td></tr><tr class="odd"><td><p>Clock</p></td><td><p>uint64</p></td><td><p>Destination logical clock</p></td></tr><tr class="even"><td><p>Message Type</p></td><td><p>byte</p></td><td><p>Message type</p></td></tr><tr class="odd"><td><p>Event Type</p></td><td><p>byte</p></td><td><p>Event type</p></td></tr><tr class="even"><td><p>Object</p></td><td><p>[]byte</p></td><td><p>Message object (CBOR)</p></td></tr></tbody></table>

The message value object has a defined object for each type of message, and the object is stored as a simple byte sequence encoded in CBOR in the message value.

### Message Type and Event Type

The message types and the event types are reserved as follows:

<table><colgroup><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /><col style="width: 25%" /></colgroup><thead><tr class="header"><th>Message Type</th><th>Event Type</th><th>Occurrence Condition</th><th>Note</th></tr></thead><tbody><tr class="odd"><td><p>Object (O)</p></td><td><p>Created ©</p></td><td><p>Object created</p></td><td></td></tr><tr class="even"><td></td><td><p>Update (U)</p></td><td><p>Object updated</p></td><td></td></tr><tr class="odd"><td></td><td><p>Delete (D)</p></td><td><p>Object deleted</p></td><td></td></tr><tr class="even"><td><p>Database (D)</p></td><td><p>Created ©</p></td><td><p>Database created</p></td><td></td></tr><tr class="odd"><td></td><td><p>Update (U)</p></td><td><p>Database updated</p></td><td></td></tr><tr class="even"><td></td><td><p>Delete (D)</p></td><td><p>Database deleted</p></td><td></td></tr><tr class="odd"><td><p>Collection ©</p></td><td><p>Created ©</p></td><td><p>Schema created</p></td><td></td></tr><tr class="even"><td></td><td><p>Update (U)</p></td><td><p>Shcema updated</p></td><td></td></tr><tr class="odd"><td></td><td><p>Delete (D)</p></td><td><p>Schema deleted</p></td><td></td></tr><tr class="even"><td><p>User Defined (U)</p></td><td><p>-</p></td><td><p>-</p></td><td><p>User-defined message. Event types are not defined.</p></td></tr></tbody></table>

The object (O) message is notified when a store object is created, updated, or deleted in the coordinator store. The database (D) and collection © messages are notified by query services such as Redis, MySQL, and MongoDB.

The user (U) message preserves only the message type, and the event type and message value object are defined by the user.

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
