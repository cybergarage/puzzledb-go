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
	opentracing "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer/ot"
	opentelemetry "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer/otel"
)

// Server represents a server instance.
type Server struct {
	*Config
	*PluginManager
	*GrpcServer
	*PrometheusExporter
}

// NewServer returns a new server instance.
func NewServer() *Server {
	server := &Server{
		GrpcServer:         nil,
		Config:             nil,
		PluginManager:      NewPluginManagerWith(plugins.NewManager()),
		PrometheusExporter: nil,
	}
	conf, err := NewDefaultConfig()
	if err != nil {
		panic(err)
	}
	server.SetConfig(conf)
	server.GrpcServer = NewGrpcServerWith(server)
	server.PrometheusExporter = NewPrometheusExporterWith(server)

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
		return err
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
		opentelemetry.NewService(),
		opentracing.NewService(),
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
		return err
	}

	defaultDocCoder, err := server.DefaultDocumentCoderService()
	if err != nil {
		return err
	}

	defaultTracer, err := server.DefaultTracingService()
	if err != nil {
		return err
	}
	defaultTracer.SetServiceName(ProductName)

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

	defaultCoodinator, err := server.DefaultCoordinatorService()
	if err != nil {
		return err
	}

	defaultStore, err := server.DefaultStoreService()
	if err != nil {
		return err
	}

	for _, service := range server.QueryServices() {
		service.SetConfig(server.Config.Config)
		service.SetCoordinator(defaultCoodinator)
		service.SetStore(defaultStore)
		service.SetTracer(defaultTracer)
		port, err := server.QueryPortConfig(service.ServiceName())
		if err != nil {
			service.SetPort(port)
		}
	}

	return nil
}

// Start starts the server.
func (server *Server) Start() error { //nolint:gocognit
	// Setup logger

	ok, _ := server.Config.GetBool(loggerConfig, enabledConfig)
	if ok {
		level := log.LevelInfo
		levelStr, err := server.Config.GetString(loggerConfig, levelConfig)
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

	// Setup gRPC server

	ok, _ = server.GrpcServer.EnabledConfig()
	if ok {
		port, err := server.GrpcServer.PortConfig()
		if err == nil {
			server.GrpcServer.SetPort(port)
		}
		if err := server.GrpcServer.Start(); err != nil {
			if stopErr := server.Stop(); stopErr != nil {
				return errors.Join(err, stopErr)
			}
			return err
		}
	} else {
		log.Infof("gRPC server disabled")
	}

	// Setup metrics server

	ok, _ = server.PrometheusExporter.EnabledConfig()
	if ok {
		port, err := server.PrometheusExporter.PortConfig()
		if err == nil {
			server.PrometheusExporter.SetPort(port)
		}
		if err := server.PrometheusExporter.Start(); err != nil {
			if stopErr := server.Stop(); stopErr != nil {
				return errors.Join(err, stopErr)
			}
			return err
		}
	} else {
		log.Infof("prometheus exporter disabled")
	}

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

	log.Infof("%s (PID:%d) started", ProductName, os.Getpid())

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	var err error
	if stopErr := server.Manager.Stop(); stopErr != nil {
		err = errors.Join(err, stopErr)
	}
	ok, _ := server.PrometheusExporter.EnabledConfig()
	if ok {
		if stopErr := server.PrometheusExporter.Stop(); stopErr != nil {
			err = errors.Join(err, stopErr)
		}
	}
	ok, _ = server.GrpcServer.EnabledConfig()
	if ok {
		if stopErr := server.GrpcServer.Stop(); stopErr != nil {
			err = errors.Join(err, stopErr)
		}
	}
	log.Infof("%s (PID:%d) terminated", ProductName, os.Getpid())
	return err
}
