# Coordinator Messaging

The message box is used to store messages sent between PuzzleDB nodes in the cluster to notify any node and store status changes using the message key-value store.

## Message Clock

Logical clocks, such as the Lamport Clock, are important in distributed systems because they can order events across different nodes; in PuzzleDB, to manage the message clock for the coordinator service, the conceptual Lamport Clock algorithm:

![coordinator clock](img/coordinator_clock.png)

In practice, the coordinator node acts as a virtual message relay node between PuzzleDB nodes. The coordinator service uses the message clock to provide the total order of messages at all nodes in the system. To manage the message clock, PuzzleDB uses the Lamport Clock algorithm, which assigns a unique timestamp to each message sent by a node.

![coordinator message clock](img/coordinator_message_clock.png)

When sending a message, a PuzzleDB node obtains the latest logical clock from the coordinator node and uses it to timestamp the message. When a PuzzleDB node receives or retrieves a message from a coordinator node, it updates its own logical clock with the logical clock included in the message.

# Message Object

The message object is encoded as a CBOR object and stored as the value of the message key-value record.

## Message Object Header

The message object has the following required header fields:

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Field Name</th>
<th style="text-align: left;">Data Type</th>
<th style="text-align: left;">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>type</p></td>
<td style="text-align: left;"><p>int</p></td>
<td style="text-align: left;"><p>Message type</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>timestamp</p></td>
<td style="text-align: left;"><p>time.Time</p></td>
<td style="text-align: left;"><p>Generated phisical time</p></td>
</tr>
</tbody>
</table>

The type field is used to identify messages and the timestamp field is used to discard old messages.

## Message Object Value

The message value object is encoded as a CBOR object and stored as the value of the message key-value record to the coordinator service. The message object has a message type, an event type and a CBOR encoded message object, and the standard message object is defined as follows:

<table>
<colgroup>
<col style="width: 33%" />
<col style="width: 33%" />
<col style="width: 33%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Field</th>
<th style="text-align: left;">Type</th>
<th style="text-align: left;">Value</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>ID</p></td>
<td style="text-align: left;"><p>UUID</p></td>
<td style="text-align: left;"><p>Destination node ID</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Host</p></td>
<td style="text-align: left;"><p>string</p></td>
<td style="text-align: left;"><p>Destination host name</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Clock</p></td>
<td style="text-align: left;"><p>uint64</p></td>
<td style="text-align: left;"><p>Destination logical clock</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Message Type</p></td>
<td style="text-align: left;"><p>byte</p></td>
<td style="text-align: left;"><p>Message type</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Event Type</p></td>
<td style="text-align: left;"><p>byte</p></td>
<td style="text-align: left;"><p>Event type</p></td>
</tr>
<tr>
<td style="text-align: left;"><p>Object</p></td>
<td style="text-align: left;"><p>[]byte</p></td>
<td style="text-align: left;"><p>Message object (CBOR)</p></td>
</tr>
</tbody>
</table>

The message value object has a defined object for each type of message, and the object is stored as a simple byte sequence encoded in CBOR in the message value.

## Message Type and Event Type

The message types and the event types are reserved as follows:

<table>
<colgroup>
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
<col style="width: 25%" />
</colgroup>
<thead>
<tr>
<th style="text-align: left;">Message Type</th>
<th style="text-align: left;">Event Type</th>
<th style="text-align: left;">Occurrence Condition</th>
<th style="text-align: left;">Note</th>
</tr>
</thead>
<tbody>
<tr>
<td style="text-align: left;"><p>Object (O)</p></td>
<td style="text-align: left;"><p>Created (C)</p></td>
<td style="text-align: left;"><p>Object created</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Update (U)</p></td>
<td style="text-align: left;"><p>Object updated</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Delete (D)</p></td>
<td style="text-align: left;"><p>Object deleted</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Database (D)</p></td>
<td style="text-align: left;"><p>Created (C)</p></td>
<td style="text-align: left;"><p>Database created</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Update (U)</p></td>
<td style="text-align: left;"><p>Database updated</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Delete (D)</p></td>
<td style="text-align: left;"><p>Database deleted</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>Collection (C)</p></td>
<td style="text-align: left;"><p>Created (C)</p></td>
<td style="text-align: left;"><p>Schema created</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Update (U)</p></td>
<td style="text-align: left;"><p>Shcema updated</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"></td>
<td style="text-align: left;"><p>Delete (D)</p></td>
<td style="text-align: left;"><p>Schema deleted</p></td>
<td style="text-align: left;"></td>
</tr>
<tr>
<td style="text-align: left;"><p>User Defined (U)</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>-</p></td>
<td style="text-align: left;"><p>User-defined message. Event types are not defined.</p></td>
</tr>
</tbody>
</table>

The object (O) message is notified when a store object is created, updated, or deleted in the coordinator store. The database (D) and collection © messages are notified by query services such as Redis, MySQL, and MongoDB.

The user (U) message preserves only the message type, and the event type and message value object are defined by the user.
