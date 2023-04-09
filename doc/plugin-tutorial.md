# Builing User Plug-ins

This section describes the plug-in interface for adding your own plug-in services and registering them into PuzzleDB.

## Plug-in interface

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

## Registering Plug-in

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

## Examples

For more information on plug-in implementation, please refer to the standard plug-ins located in the `plugins` directory.
