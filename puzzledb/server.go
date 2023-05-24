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
	"errors"
	"os"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	coord "github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/document/cbor"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key/tuple"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
	etcd_coordinator "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/etcd"
	fdb_coordinator "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/fdb"
	memdb_coordinator "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/memdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/metrics/prometheus"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mongo"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/memdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/system/actor"
	opentracing "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer/ot"
	opentelemetry "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer/otel"
)

// Server represents a server instance.
type Server struct {
	actorService *actor.Service
	*Config
	*PluginManager
	coord.Process
}

// NewServer returns a new server instance.
func NewServer() *Server {
	server := &Server{
		actorService:  nil,
		Config:        nil,
		PluginManager: NewPluginManagerWith(plugins.NewManager()),
		Process:       coord.NewProcess(),
	}
	conf, err := NewDefaultConfig()
	if err != nil {
		panic(err)
	}
	server.SetConfig(conf)

	return server
}

// NewServerWithConfig returns a new server instance with the specified configuradtion.
func NewServerWithConfig(conf config.Config) *Server {
	server := NewServer()
	server.SetConfig(NewConfigWith(conf))
	return server
}

// SetConfig sets a server configuration.
func (server *Server) SetConfig(conf *Config) {
	server.Config = conf
	server.Manager.SetConfig(conf)
}

// Restart restarts the server.
func (server *Server) Restart() error {
	if err := server.Stop(); err != nil {
		return err
	}
	return server.Start()
}

func (server *Server) reloadEmbeddedPlugins() error {
	server.actorService = actor.NewService()
	services := []plugins.Service{
		server.actorService,
		NewGrpcServiceWith(server),
		cbor.NewCoder(),
		tuple.NewCoder(),
		store.NewStore(),
		fdb.NewStore(),
		memdb.NewStore(),
		coordinator.NewServiceWith(etcd_coordinator.NewCoordinator()),
		coordinator.NewServiceWith(memdb_coordinator.NewCoordinator()),
		coordinator.NewServiceWith(fdb_coordinator.NewCoordinator()),
		mysql.NewService(),
		redis.NewService(),
		mongo.NewService(),
		opentelemetry.NewService(),
		opentracing.NewService(),
		prometheus.NewService(),
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
	server.PluginManager.SetConfig(server.Config)

	// Default services

	defaultKeyCoder, err := server.DefaultKeyCoderService()
	if err != nil {
		return err
	}

	defaultDocCoder, err := server.DefaultDocumentCoderService()
	if err != nil {
		return err
	}

	defaultCoodinator, err := server.DefaultCoordinatorService()
	if err != nil {
		return err
	}

	defaultStore, err := server.DefaultStoreService()
	if err != nil {
		return err
	}

	defaultTracer, err := server.DefaultTracingService()
	if err != nil {
		return err
	}
	defaultTracer.SetServiceName(ProductName)

	// Coordinator services

	for _, service := range server.CoordinatorServices() {
		service.SetProcess(server.Process)
		service.SetKeyCoder(defaultKeyCoder)
	}

	// Actor service

	server.actorService.SetCoordinator(defaultCoodinator)

	// KV store services

	for _, service := range server.KvStoreServices() {
		service.SetKeyCoder(defaultKeyCoder)
	}

	// Document store services

	defaultKvStore, err := server.DefaultKvStoreService()
	if err != nil {
		return err
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

	for _, service := range server.QueryServices() {
		service.SetCoordinator(defaultCoodinator)
		service.SetStore(defaultStore)
		service.SetTracer(defaultTracer)
	}

	return nil
}

// Start starts the server.
func (server *Server) Start() error { //nolint:gocognit
	// Setup logger

	ok, _ := server.Config.GetConfigBool(ConfigLogger, ConfigEnabled)
	if ok {
		level := log.LevelInfo
		levelStr, err := server.Config.GetConfigString(ConfigLogger, ConfigLevel)
		if err == nil {
			level = log.GetLevelFromString(levelStr)
		}
		log.SetSharedLogger(log.NewStdoutLogger(level))
	} else {
		log.SetSharedLogger(nil)
	}

	// Output version

	log.Infof("%s/%s", ProductName, Version)

	// Output logger settings

	log.Infof("logger (%s) started", log.GetLevelString(log.GetSharedLogger().Level()))

	// Setup configuration

	log.Infof("configuration loaded")
	log.Infof(server.Config.String())

	// Setup plugins

	if err := server.LoadPlugins(); err != nil {
		return err
	}

	if err := server.setupPlugins(); err != nil {
		return err
	}

	if err := server.Manager.Start(); err != nil {
		if stopErr := server.Stop(); stopErr != nil {
			return errors.Join(err, stopErr)
		}
		return err
	}

	log.Infof("%s", server.Manager.String())

	// Output success message

	log.Infof("%s (PID:%d) started", ProductName, os.Getpid())

	server.actorService.SetStatus(actor.StatusRunning)

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	var err error
	if stopErr := server.Manager.Stop(); stopErr != nil {
		err = errors.Join(err, stopErr)
	}

	log.Infof("%s (PID:%d) terminated", ProductName, os.Getpid())

	if err == nil {
		server.actorService.SetStatus(actor.StatusStopped)
	}

	return err
}
