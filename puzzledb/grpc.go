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
	"net"
	"strconv"

	"github.com/cybergarage/go-logger/log"
	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/api"
	"google.golang.org/grpc"
)

const (
	// DefaultGrpcPort is the default port number of the gRPC server.
	DefaultGrpcPort = 50053
)

type GrpcServer struct {
	*Server
	grpcServer *grpc.Server
	Addr       string
	Port       int
	pb.UnimplementedStoreAPIServer
}

// NewGrpcServerWith returns a new GrpcServer.
func NewGrpcServerWith(server *Server) *GrpcServer {
	return &GrpcServer{
		Server:                      server,
		grpcServer:                  nil,
		Addr:                        "",
		Port:                        DefaultGrpcPort,
		UnimplementedStoreAPIServer: pb.UnimplementedStoreAPIServer{},
	}
}

// SetPort sets a port number of the server.
func (server *GrpcServer) SetPort(port int) {
	server.Port = port
}

// Start starts the server.
func (server *GrpcServer) Start() error {
	var err error
	addr := net.JoinHostPort(server.Addr, strconv.Itoa(server.Port))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server.grpcServer = grpc.NewServer()
	pb.RegisterStoreAPIServer(server.grpcServer, server)
	go func() {
		if err := server.grpcServer.Serve(listener); err != nil {
			log.Error(err)
		}
	}()

	log.Infof("gRPC (%s) started", addr)

	return nil
}

// Stop stops the Grpc server.
func (server *GrpcServer) Stop() error {
	if server.grpcServer != nil {
		server.grpcServer.Stop()
		server.grpcServer = nil
	}

	addr := net.JoinHostPort(server.Addr, strconv.Itoa(server.Port))
	log.Infof("gRPC (%s) terminated", addr)

	return nil
}
