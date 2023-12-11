Builing New Plug-ins
====================

This section describes the plug-in interface for adding your own plug-in services and registering them into PuzzleDB.

Plug-in interface
-----------------

Each plug-in service should implement the following `Service` interface, which is located in the `puzzledb/plugins` directory.

    type Service interface {
        // ServiceType returns the service type.
        ServiceType() ServiceType
        // ServiceName returns the service name.
        ServiceName() string
        // Start starts the service
        Start() error
        // Stop stops the service
        Stop() error
    }

Standard Plug-in interfaces
---------------------------

For the plugin services specified in the standards listed in the following table, refer to each plugin interface that is reserved in the `plugins` directory.

<table style="width:100%;"><colgroup><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /><col style="width: 16%" /></colgroup><thead><tr class="header"><th>Major Type</th><th>Sub Type</th><th>Description</th><th>Plug-ins</th><th>Distributed</th><th>Dependency</th></tr></thead><tbody><tr class="odd"><td><p>System</p></td><td><p>-</p></td><td><p>System services</p></td><td><p>gRPC</p></td><td><p>O</p></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>Actor</p></td><td><p>O</p></td><td><p>Coordinator</p></td></tr><tr class="odd"><td><p>Query</p></td><td><p>-</p></td><td><p>Query handler services</p></td><td><p>Redis</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>MongoDB</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>MySQL</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>PostgreSQL</p></td><td><p>O</p></td><td><p>Store (Document)</p></td></tr><tr class="odd"><td><p>Coordinator</p></td><td><p>-</p></td><td><p>Coordination services</p></td><td><p>memdb</p></td><td><p>X</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>etcd (Planning)</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>ZooKeeper (Planning)</p></td><td></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>FoundationDB (Planning)</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td><p>Coder</p></td><td><p>Document</p></td><td><p>Document coder services</p></td><td><p>CBOR</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="even"><td></td><td><p>Key</p></td><td><p>Key coder services</p></td><td><p>Tuple</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td><p>Store</p></td><td><p>Document</p></td><td><p>Doument store services</p></td><td><p>Key-value based store</p></td><td><p>O</p></td><td><p>Store (Key-value), Coder (Document), Coder (Key)</p></td></tr><tr class="even"><td></td><td><p>Key-value</p></td><td><p>Key-value store services</p></td><td><p>memdb</p></td><td><p>X</p></td><td><p>Coder (Document), Coder (Key)</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>FoundationDB</p></td><td><p>O</p></td><td><p>Coder (Document), Coder (Key)</p></td></tr><tr class="even"><td></td><td></td><td></td><td><p>TiKV (Planning)</p></td><td><p>O</p></td><td><p>-</p></td></tr><tr class="odd"><td></td><td></td><td></td><td><p>JunoDB (Planning)</p></td><td></td><td></td></tr><tr class="even"><td></td><td><p>Key-Value Cache</p></td><td><p>Key-value cache store services</p></td><td><p>Ristretto</p></td><td><p>O</p></td><td><p>Store (Key-value)</p></td></tr><tr class="odd"><td><p>Tracer</p></td><td><p>-</p></td><td><p>Distributed tracing services</p></td><td><p>OpenTelemetry</p></td><td><p>O</p></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>OpenTracing</p></td><td><p>O</p></td><td></td></tr><tr class="odd"><td><p>Metric</p></td><td><p>-</p></td><td><p>Metrics services</p></td><td><p>Prometheus</p></td><td><p>O</p></td><td></td></tr><tr class="even"><td></td><td></td><td></td><td><p>Graphite (Planning)</p></td><td><p>O</p></td><td></td></tr><tr class="odd"><td><p>Extend</p></td><td><p>-</p></td><td><p>User-defined services</p></td><td><p>-</p></td><td><p>-</p></td><td><p>-</p></td></tr></tbody></table>

For more information on plug-in implementation, please refer to the standard plug-ins located in the `plugins` directory.

Registering Plug-in
-------------------

To register your plug-in service, you should override `` Server::LoadPlugins()` `` as follows:

    import (
        "github.com/cybergarage/puzzledb-go/puzzledb"
    )

    type UserServer struct {
        *puzzledb.Server
        Host string
    }

    func NewServerWithConfig(config puzzledb.Config) *UserServer {
        server := &UserServer{
            Server: puzzledb.NewServerWithConfig(config),
        }
        return server
    }

    func (server *UserServer) LoadPlugins() error {
        if err := server.Server.LoadPlugins(); err != nil {
            return err
        }
        // Register your plug-in service
        var service puzzledb.Service = ....
        server.RegisterService(service)
        return nil
    }
