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

package puzzledbtest

import (
	"github.com/cybergarage/puzzledb-go/puzzledb"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

// Server represents an example server.
type Server struct {
	*puzzledb.Server
	Host string
}

// NewServer returns a new testserver instance.
func NewServer() *Server {
	server := &Server{
		Server: puzzledb.NewServer(),
		Host:   LocalHost,
	}

	conf, err := puzzledb.NewConfigWithString(puzzledbConf)
	if err != nil {
		panic(err)
	}
	server.SetConfig(conf)

	return server
}

// NewServerWithConfig returns a new test server instance with the specified configuration.
func NewServerWithConfig(config config.Config) *Server {
	server := &Server{
		Server: puzzledb.NewServerWithConfig(config),
		Host:   LocalHost,
	}
	return server
}

func (server *Server) Store() *Store {
	store, _ := server.DefaultStoreService()
	return NewStoreWith(store)
}
