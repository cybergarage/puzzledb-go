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
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/cybergarage/go-logger/log"
	pc "github.com/cybergarage/puzzledb-go/puzzledb/context"
	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/grpc"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	pb.UnimplementedHealthServer
	pb.UnimplementedMetricServer
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
		UnimplementedHealthServer: pb.UnimplementedHealthServer{},
		UnimplementedMetricServer: pb.UnimplementedMetricServer{},
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
	pb.RegisterHealthServer(server.grpcServer, server)
	pb.RegisterMetricServer(server.grpcServer, server)
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
		log.Infof("gRPC Request: %s", info.FullMethod)
	} else {
		log.Errorf("gRPC Request: %s", info.FullMethod)
	}

	return resp, err
}

func (server *GrpcServer) Check(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	res := pb.HealthCheckResponse{} //nolint:exhaustruct
	res.Status = pb.HealthCheckResponse_SERVING
	return &res, nil
}

func (server *GrpcServer) ListConfig(context.Context, *pb.ListConfigRequest) (*pb.ListConfigResponse, error) {
	res := pb.ListConfigResponse{} //nolint:exhaustruct
	res.Values = strings.Split(server.Config.String(), "\n")
	return &res, nil
}

func (server *GrpcServer) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	v, err := server.Config.Get(req.Name)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%s not found", req.Name))
	}
	res := pb.GetConfigResponse{} //nolint:exhaustruct
	res.Value = fmt.Sprintf("%v", v)
	return &res, nil
}

func (server *GrpcServer) GetVersion(context.Context, *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	res := pb.GetVersionResponse{} //nolint:exhaustruct
	res.Value = fmt.Sprintf("%v", Version)
	return &res, nil
}

func (server *GrpcServer) ListMetric(context.Context, *pb.ListMetricRequest) (*pb.ListMetricResponse, error) {
	res := pb.ListMetricResponse{} //nolint:exhaustruct
	metrics, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return &res, err
	}
	for _, metric := range metrics {
		res.Values = append(res.Values, metric.GetName())
	}
	return &res, nil
}

func (server *GrpcServer) GetMetric(ctx context.Context, req *pb.GetMetricRequest) (*pb.GetMetricResponse, error) {
	res := pb.GetMetricResponse{} //nolint:exhaustruct
	metrics, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return &res, err
	}
	for _, metric := range metrics {
		if metric.GetName() != req.Name {
			res.Value = metric.String()
			return &res, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%s not found", req.Name))
}

func (server *GrpcServer) CreateDatabase(ctx context.Context, req *pb.CreateDatabaseRequest) (*pb.StatusResponse, error) {
	res := pb.StatusResponse{} //nolint:exhaustruct
	defaultStore, err := server.DefaultStoreService()
	if err != nil {
		return &res, err
	}
	err = defaultStore.CreateDatabase(pc.NewContext(), req.DatabaseName)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func (server *GrpcServer) RemoveDatabase(ctx context.Context, req *pb.RemoveDatabaseRequest) (*pb.StatusResponse, error) {
	res := pb.StatusResponse{} //nolint:exhaustruct
	defaultStore, err := server.DefaultStoreService()
	if err != nil {
		return &res, err
	}
	err = defaultStore.RemoveDatabase(pc.NewContext(), req.DatabaseName)
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func (server *GrpcServer) ListDatabases(context.Context, *pb.ListDatabasesRequest) (*pb.ListDatabasesResponse, error) {
	res := pb.ListDatabasesResponse{} //nolint:exhaustruct
	defaultStore, err := server.DefaultStoreService()
	if err != nil {
		return &res, err
	}
	dbs, err := defaultStore.ListDatabases(pc.NewContext())
	if err != nil {
		return &res, err
	}
	res.Databases = []string{}
	for _, db := range dbs {
		res.Databases = append(res.Databases, db.Name())
	}
	return &res, nil
}

func (server *GrpcServer) ListCollections(context.Context, *pb.ListCollectionsRequest) (*pb.ListCollectionsResponse, error) {
	res := pb.ListCollectionsResponse{} //nolint:exhaustruct
	return &res, nil
}
