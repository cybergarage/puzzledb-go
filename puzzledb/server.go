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
	std_tls "crypto/tls"
	"errors"
	std_log "log"
	"net"
	"net/http"
	_ "net/http/pprof" //nolint:gosec
	"os"
	"strconv"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/puzzledb-go/puzzledb/auth"
	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	auth_plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/auth"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/document/cbor"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key/tuple"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
	fdb_coordinator "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/fdb"
	memdb_coordinator "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator/core/memdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/metrics/prometheus"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mongo"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/postgresql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/fdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv/memdb"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kvcache/ristretto"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/system/actor"
	opentracing "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer/ot"
	opentelemetry "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer/otel"
	"github.com/cybergarage/puzzledb-go/puzzledb/tls"
)

// Server represents a server instance.
type Server struct {
	actor *actor.Service
	Config
	tlsConfig *std_tls.Config
	auth.AuthManager
	*PluginManager
	cluster.Node
	pprofStarted bool
}

// NewServer returns a new server instance.
func NewServer() *Server {
	conf, err := NewDefaultConfig()
	if err != nil {
		panic(err)
	}
	return NewServerWithConfig(conf)
}

// NewServerWithConfig returns a new server instance with the specified configuradtion.
func NewServerWithConfig(conf config.Config) *Server {
	server := &Server{
		actor:         nil,
		Config:        nil,
		tlsConfig:     nil,
		AuthManager:   auth.NewAuthManager(),
		PluginManager: NewPluginManagerWith(plugins.NewManager()),
		Node:          cluster.NewNode(),
		pprofStarted:  false,
	}
	server.SetConfig(conf)
	return server
}

// SetConfig sets a server configuration.
func (server *Server) SetConfig(conf Config) {
	server.Config = conf
	server.Manager.SetConfig(conf)
}

// SetTLSConfig sets a TLS configuration.
func (server *Server) SetTLSConfig(tlsConfig *std_tls.Config) {
	server.tlsConfig = tlsConfig
}

// IsTLSEnabled returns true if TLS is enabled.
func (server *Server) IsTLSEnabled() bool {
	tlsConf, err := tls.NewConfigWith(server.Config, ConfigTLS)
	if err != nil {
		return false
	}
	return tlsConf.TLSEnabled()
}

// TLSConfig returns a TLS configuration.
func (server *Server) TLSConfig() (*std_tls.Config, bool) {
	if server.tlsConfig != nil {
		return server.tlsConfig, true
	}
	tlsConf, err := tls.NewConfigWith(server.Config, ConfigTLS)
	if err != nil {
		return nil, false
	}
	tlsConfig, err := tlsConf.TLSConfig()
	if err != nil {
		return nil, false
	}
	return tlsConfig, true
}

// Restart restarts the server.
func (server *Server) Restart() error {
	if err := server.Stop(); err != nil {
		return err
	}
	return server.Start()
}

func (server *Server) reloadEmbeddedPlugins() error {
	server.actor = actor.NewService()
	services := []plugins.Service{
		cbor.NewCoder(),
		tuple.NewCoder(),
		coordinator.NewServiceWith(fdb_coordinator.NewCoordinator()),
		coordinator.NewServiceWith(memdb_coordinator.NewCoordinator()),
		// coordinator.NewServiceWith(etcd_coordinator.NewCoordinator()),
		store.NewStore(),
		fdb.NewStore(),
		memdb.NewStore(),
		ristretto.NewStore(),
		postgresql.NewService(),
		mysql.NewService(),
		redis.NewService(),
		mongo.NewService(),
		opentelemetry.NewService(),
		opentracing.NewService(),
		prometheus.NewService(),
		NewGrpcServiceWith(server),
		auth_plugin.NewService(),
		server.actor,
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
	var err error

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

	// Coordinator services

	if services, err := server.CoordinatorServices(); err != nil {
		return err
	} else {
		for _, service := range services {
			service.SetNode(server.Node)
			service.SetKeyCoder(defaultKeyCoder)
		}
	}

	// Actor service

	server.actor.SetCoordinator(defaultCoodinator)

	// KV store services

	if services, err := server.KvStoreServices(); err != nil {
		return err
	} else {
		for _, service := range services {
			service.SetKeyCoder(defaultKeyCoder)
		}
	}

	defaultKvStore, err := server.DefaultKvStoreService()
	if err != nil {
		return err
	}

	// KV cache store services

	if services, err := server.KvCacheStoreServices(); err != nil {
		return err
	} else {
		for _, service := range services {
			service.SetStore(defaultKvStore)
		}
	}

	if server.Manager.IsServiceTypeConfigEnabled(plugins.StoreKvCacheService) {
		defaultKvCacheStore, err := server.DefaultKvCacheStoreService()
		if err == nil {
			defaultKvStore = defaultKvCacheStore
		}
	}

	// Document store services

	if services, err := server.DocumentStoreServices(); err != nil {
		return err
	} else {
		for _, service := range services {
			service.SetKeyCoder(defaultKeyCoder)
			service.SetDocumentCoder(defaultDocCoder)
			store, ok := service.(*store.Store)
			if !ok {
				return plugins.NewErrInvalidService(service)
			}
			store.SetKvStore(defaultKvStore)
		}
	}
	// Setup tracer

	defaultTracer := tracer.NewNullTracer()
	ok := server.Manager.IsServiceTypeConfigEnabled(plugins.TracingService)
	if ok {
		defaultTracer, err = server.DefaultTracingService()
		if err != nil {
			return err
		}
	}
	defaultTracer.SetPackageName(PackageName)
	defaultTracer.SetServiceName(ProductName)

	// TLS configuration

	tlsConf, err := tls.NewConfigWith(server.Config, ConfigTLS)
	if err != nil {
		return err
	}

	tlsConfig := server.tlsConfig
	if tlsConf.TLSEnabled() {
		if tlsConfig == nil {
			tlsConfig, err = tlsConf.TLSConfig()
			if err != nil {
				return err
			}
		}
	} else {
		tlsConfig = nil
	}

	// Authenticator service

	authService, err := server.DefaultAuthenticatorService()
	if err != nil {
		return err
	}

	// Query services

	if services, err := server.QueryServices(); err != nil {
		return err
	} else {
		for _, service := range services {
			service.SetCoordinator(defaultCoodinator)
			service.SetStore(defaultStore)
			service.SetTracer(defaultTracer)
			service.SetTLSConfig(tlsConfig)

			service.SetCertificateAuthenticator(authService)
			service.SetCredentialAuthenticator(authService)

			err := defaultCoodinator.AddObserver(service)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Start starts the server.
func (server *Server) Start() error { //nolint:gocognit
	// Setup logger

	ok, _ := server.Config.LookupConfigBool(ConfigLogger, ConfigEnabled)
	if ok {
		level := log.LevelInfo
		levelStr, err := server.Config.LookupConfigString(ConfigLogger, ConfigLevel)
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

	// Setup pprof

	ok, _ = server.Config.LookupConfigBool(ConfigPprof, ConfigEnabled)
	if ok && !server.pprofStarted {
		port, err := server.Config.LookupConfigInt(ConfigPprof, ConfigPort)
		if err != nil {
			return err
		}
		go func() {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			std_log.Println(http.ListenAndServe(addr, nil)) //nolint:errcheck,gosec
		}()
	}

	// Setup configuration

	log.Info("configuration loading...")
	log.Info(server.Config.UsedConfigFile())
	log.Info(server.Config.String())

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

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	var err error
	if stopErr := server.Manager.Stop(); stopErr != nil {
		err = errors.Join(err, stopErr)
	}

	log.Infof("%s (PID:%d) terminated", ProductName, os.Getpid())

	return err
}
