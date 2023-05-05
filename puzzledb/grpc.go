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
	"context"
	"net"
	"strconv"

	"github.com/cybergarage/go-logger/log"
	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/grpc"
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
	pb.UnimplementedStoreServer
	pb.UnimplementedConfigServer
}

// NewGrpcServerWith returns a new GrpcServer.
func NewGrpcServerWith(server *Server) *GrpcServer {
	return &GrpcServer{
		Server:                    server,
		grpcServer:                nil,
		Addr:                      "",
		Port:                      DefaultGrpcPort,
		UnimplementedStoreServer:  pb.UnimplementedStoreServer{},
		UnimplementedConfigServer: pb.UnimplementedConfigServer{},
	}
}

// SetPort sets a port number of the server.
func (server *GrpcServer) SetPort(port int) {
	server.Port = port
}

// EnabledConfig returns a port number for the specified query service name.
func (server *GrpcServer) EnabledConfig() (bool, error) {
	return server.Config.GetBool(apiConfig, grpcConfig, enabledConfig)
}

// PortConfig returns a port number for the specified query service name.
func (server *GrpcServer) PortConfig() (int, error) {
	return server.Config.GetInt(apiConfig, grpcConfig, portConfig)
}

// Start starts the server.
func (server *GrpcServer) Start() error {
	var err error
	addr := net.JoinHostPort(server.Addr, strconv.Itoa(server.Port))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(loggingUnaryInterceptor))
	pb.RegisterStoreServer(server.grpcServer, server)
	pb.RegisterConfigServer(server.grpcServer, server)
	go func() {
		if err := server.grpcServer.Serve(listener); err != nil {
			log.Error(err)
		}
	}()

	log.Infof("gRPC server (%s) started", addr)

	return nil
}

// Stop stops the Grpc server.
func (server *GrpcServer) Stop() error {
	if server.grpcServer != nil {
		server.grpcServer.Stop()
		server.grpcServer = nil
	}

	addr := net.JoinHostPort(server.Addr, strconv.Itoa(server.Port))
	log.Infof("gRPC server (%s) terminated", addr)

	return nil
}

func loggingUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	resp, err := handler(ctx, req)

	if err == nil {
		log.Infof("gRPC Method: %s, Request: %v, Response: %v",
			info.FullMethod,
			req,
			resp)
	} else {
		log.Errorf("gRPC Method: %s, Request: %v, Response: %v",
			info.FullMethod,
			req,
			err.Error())
	}

	return resp, err
}
