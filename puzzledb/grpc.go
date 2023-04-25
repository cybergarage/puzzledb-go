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
	net.Listener
	pb.UnimplementedStoreAPIServer
}

// NewGrpcServerWith returns a new GrpcServer.
func NewGrpcServerWith(server *Server) *GrpcServer {
	return &GrpcServer{
		Server:                      server,
		grpcServer:                  nil,
		Addr:                        "",
		Port:                        DefaultGrpcPort,
		Listener:                    nil,
		UnimplementedStoreAPIServer: pb.UnimplementedStoreAPIServer{},
	}
}

// Start starts the server.
func (server *GrpcServer) Start() error {
	var err error
	addr := net.JoinHostPort(server.Addr, strconv.Itoa(server.Port))
	server.Listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server.grpcServer = grpc.NewServer()
	pb.RegisterStoreAPIServer(server.grpcServer, server)
	go server.grpcServer.Serve(server.Listener)
	return nil
}

// Stop stops the Grpc server.
func (server *GrpcServer) Stop() error {
	if server.grpcServer != nil {
		server.grpcServer.Stop()
		server.grpcServer = nil
	}
	if server.Listener != nil {
		err := server.Listener.Close()
		if err != nil {
			return err
		}
		server.Listener = nil
	}
	return nil
}
