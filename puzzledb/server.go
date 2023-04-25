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
	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/cybergarage/puzzledb-go/puzzledb/errors"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/document/cbor"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key/tuple"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
	etcd_coordinator "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/etcd"
	memdb_coordinator "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/memdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mongo"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/memdb"
)

// Server represents a server instance.
type Server struct {
	*Config
	*PluginManager
	*GrpcServer
	tracer.Tracer
}

// NewServer returns a new server instance.
func NewServer() *Server {
	server := &Server{
		GrpcServer:    nil,
		Config:        nil,
		PluginManager: NewPluginManagerWith(plugins.NewManager()),
		Tracer:        tracer.NullTracer,
	}
	server.GrpcServer = NewGrpcServerWith(server)
	return server
}

// NewServerWithConfig returns a new server instance with the specified configuradtion.
func NewServerWithConfig(config config.Config) *Server {
	server := NewServer()
	server.SetConfig(config)
	return server
}

// SetConfig sets the server configuration.
func (server *Server) SetConfig(config config.Config) {
	server.Config = NewConfigWith(config)
	server.Manager.SetConfig(config)
}

// Restart restarts the server.
func (server *Server) Restart() error {
	if err := server.Stop(); err != nil {
		return errors.Wrap(err)
	}
	return server.Start()
}

func (server *Server) reloadEmbeddedPlugins() error {
	services := []plugins.Service{
		cbor.NewCoder(),
		tuple.NewCoder(),
		store.NewStore(),
		fdb.NewStore(),
		memdb.NewStore(),
		coordinator.NewServiceWith(etcd_coordinator.NewCoordinator()),
		coordinator.NewServiceWith(memdb_coordinator.NewCoordinator()),
		mysql.NewService(),
		redis.NewService(),
		mongo.NewService(),
	}

	server.Manager.ReloadServices(services)

	return nil
}

func (server *Server) LoadPlugins() error {
	if err := server.reloadEmbeddedPlugins(); err != nil {
		return err
	}
	return nil
}

func (server *Server) setupPlugins() error {
	// Default services

	defaultKeyCoder, err := server.DefaultKeyCoderService()
	if err != nil {
		return errors.Wrap(err)
	}

	defaultDocCoder, err := server.DefaultDocumentCoderService()
	if err != nil {
		return errors.Wrap(err)
	}

	// KV store services

	for _, service := range server.KvStoreServices() {
		service.SetKeyCoder(defaultKeyCoder)
	}

	// Document store services

	defaultKvStore, err := server.DefaultKvStoreService()
	if err != nil {
		return errors.Wrap(err)
	}

	for _, service := range server.DocumentStoreServices() {
		service.SetKeyCoder(defaultKeyCoder)
		service.SetDocumentCoder(defaultDocCoder)
		store, ok := service.(*store.Store)
		if !ok {
			return plugins.NewErrInvalidService(service)
		}
		store.SetKvStore(defaultKvStore)
	}

	// Query services

	defaultCoodinator, err := server.DefaultCoordinatorService()
	if err != nil {
		return errors.Wrap(err)
	}

	defaultStore, err := server.DefaultStoreService()
	if err != nil {
		return errors.Wrap(err)
	}

	for _, service := range server.QueryServices() {
		service.SetConfig(server.Config)
		service.SetCoordinator(defaultCoodinator)
		service.SetStore(defaultStore)
		service.SetTracer(server.Tracer)
	}

	return nil
}

// SetTracer sets the tracing tracer.
func (server *Server) SetTracer(t tracer.Tracer) {
	server.Tracer = t
}

// Start starts the server.
func (server *Server) Start() error {
	log.Infof("%s/%s", ProductName, Version)

	if server.Config != nil {
		log.Infof("configuration loaded")
		log.Infof(server.Config.String())
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

	if err := server.GrpcServer.Start(); err != nil {
		return errors.Wrap(err)
	}

	log.Infof("%s", server.Manager.String())
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
