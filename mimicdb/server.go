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

package mimicdb

import (
	"github.com/cybergarage/mimicdb/mimicdb/errors"
	"github.com/cybergarage/mimicdb/mimicdb/plugins"
	"github.com/cybergarage/mimicdb/mimicdb/plugins/query/mysql"
	"github.com/cybergarage/mimicdb/mimicdb/plugins/store/memdb"
)

// Server represents a server instance.
type Server struct {
	*plugins.Services
}

// NewServer returns a new server instance.
func NewServer() *Server {
	server := &Server{
		Services: plugins.NewServices(),
	}

	server.LoadPlugins()

	return server
}

// Start starts the server.
func (server *Server) Start() error {
	if err := server.Services.Start(); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// Stop stops the server.
func (server *Server) Stop() error {
	if err := server.Services.Stop(); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

// LoadPlugins loads default plugin services.
func (server *Server) LoadPlugins() {
	store := memdb.NewStore()
	server.Services.Add(store)

	server.Services.Add(mysql.NewServiceWithStore(store))
}
