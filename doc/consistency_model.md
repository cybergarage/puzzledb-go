![](img/logo.png)

# Consistency Model

PuzzleDB is a multi-data model database; PuzzleDB is a pluggable database that combines modules, and the storage layer modules are expected to satisfy ACID-like interfaces.

PuzzleDB defines the top-level storage plug-in as a document model interface, and the storage interface consists of transaction and document interfaces.

![](img/consistency_model.png)

While developers can omit the interface and implement the storage plug-ins based on non-ACID storage, such as contingent consistency model storage, PuzzleDB expects that storage modules are implemented based on ACID storages.
