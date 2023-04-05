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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/document/cbor"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mongo"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
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

	server.LoadPlugins()

	return server
}

// NewServerWithConfig returns a new server instance with the specified configuration.
func NewServerWithConfig(config Config) *Server {
	server := NewServer()
	server.SetConfig(config)
	return server
}

// SetConfig sets the server configuration.
func (server *Server) SetConfig(config Config) {
	server.ServerConfig = NewServerConfigWith(config)
}

// Start starts the server.
func (server *Server) Start() error {
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

// Restart restarts the server.
func (server *Server) Restart() error {
	if err := server.Stop(); err != nil {
		return errors.Wrap(err)
	}
	return server.Start()
}

// LoadPlugins loads default plugin services.
func (server *Server) LoadPlugins() {
	services := []plugins.Service{}

	seralizer := cbor.NewSerializer()
	services = append(services, seralizer)

	store := store.NewStoreWithKvStore(memdb.NewStore())
	store.SetSerializer(seralizer)
	services = append(services, store)

	queryServices := []query.Service{
		mysql.NewService(),
		redis.NewService(),
		mongo.NewService(),
	}
	for _, queryService := range queryServices {
		queryService.SetStore(store)
		services = append(services, queryService)
	}

	for _, service := range services {
		server.Manager.Add(service)
	}
}
