// Copyright (C) 2022 The PuzzleDB Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package puzzledb

import (
	"os"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb/errors"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/document/cbor"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key/tuple"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
	coordinator_etcd "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/etcd"
	coordinator_memdb "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/memdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mongo"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/memdb"
)

// Server represents a server instance.
type Server struct {
	*ServerConfig
	*plugins.Manager
}

// NewServer returns a new server instance.
func NewServer() *Server {
	server := &Server{
		ServerConfig: nil,
		Manager:      plugins.NewManager(),
	}

	return server
}

// NewServerWithConfig returns a new server instance with the specified configuradtion.
func NewServerWithConfig(config Config) *Server {
	server := NewServer()
	server.SetConfig(config)
	return server
}

// SetConfig sets the server configuration.
func (server *Server) SetConfig(config Config) {
	server.ServerConfig = NewServerConfigWith(config)
	server.Manager.SetConfig(config)
}

// Restart restarts the server.
func (server *Server) Restart() error {
	if err := server.Stop(); err != nil {
		return errors.Wrap(err)
	}
	return server.Start()
}

func (server *Server) loadEmbeddedPlugins() error {
	services := []plugins.Service{}

	docCoder := cbor.NewCoder()
	services = append(services, docCoder)

	keyCoder := tuple.NewCoder()
	services = append(services, keyCoder)

	kvStores := []kv.Service{
		memdb.NewStore(),
		fdb.NewStore(),
	}
	for _, kvStore := range kvStores {
		kvStore.SetKeyCoder(keyCoder)
		services = append(services, kvStore)
	}

	coordServices := []*coordinator.Service{
		coordinator.NewServiceWith(coordinator_memdb.NewCoordinator()),
		coordinator.NewServiceWith(coordinator_etcd.NewCoordinator()),
	}
	for _, coordService := range coordServices {
		services = append(services, coordService)
	}

	kvStore := kvStores[0]
	store := store.NewStore()
	store.SetKvStore(kvStore)
	store.SetDocumentCoder(docCoder)
	store.SetKeyCoder(keyCoder)
	services = append(services, store)

	queryServices := []query.Service{
		mysql.NewService(),
		redis.NewService(),
		mongo.NewService(),
	}
	for _, queryService := range queryServices {
		if err := queryService.SetStore(store); err != nil {
			return errors.Wrap(err)
		}
		services = append(services, queryService)
	}

	server.Manager.ReloadServices(services)

	return nil
}

func (server *Server) LoadPlugins() error {
	if err := server.loadEmbeddedPlugins(); err != nil {
		return err
	}
	return nil
}

func (server *Server) setupPlugins() error {
	// Query services

	defaultStore, err := server.DefaultService(plugins.StoreDocumentService)
	if err != nil {
		return errors.Wrap(err)
	}

	services := server.ServicesByType(plugins.QueryService)
	for _, service := range services {
		queryService, ok := service.(query.Service)
		if !ok {
			return plugins.NewErrInvalidService(service)
		}
		if err := queryService.SetStore(defaultStore); err != nil {
			return errors.Wrap(err)
		}
	}

	return nil
}

// Start starts the server.
func (server *Server) Start() error {
	log.Infof("%s/%s", ProductName, Version)

	if server.ServerConfig != nil {
		log.Infof("configuration loaded")
		log.Infof(server.ServerConfig.String())
	}

	if err := server.LoadPlugins(); err != nil {
		return errors.Wrap(err)
	}

	if err := server.setupPlugins(); err != nil {
		return errors.Wrap(err)
	}

	if err := server.Manager.Start(); err != nil {
		return errors.Wrap(err)
	}

	log.Infof("%s (PID:%d) started", ProductName, os.Getpid())

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	if err := server.Manager.Stop(); err != nil {
		return errors.Wrap(err)
	}
	log.Infof("%s (PID:%d) terminated", ProductName, os.Getpid())
	return nil
}
