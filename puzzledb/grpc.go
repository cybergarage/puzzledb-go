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
	pb "github.com/cybergarage/puzzledb-go/puzzledb/api/gen/go/v1"
	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	pc "github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	// DefaultGrpcPort is the default port number of the gRPC server.
	DefaultGrpcPort = 50053
)

// gRPCService represents a gRPC service.
type gRPCService struct {
	plugins.Config
	*Server
	grpcServer *grpc.Server
	Addr       string
	Port       int
	pb.UnimplementedStoreServer
	pb.UnimplementedConfigServer
	pb.UnimplementedHealthServer
	pb.UnimplementedMetricServer
}

// NewGrpcServiceWith returns a new GrpcServer.
func NewGrpcServiceWith(server *Server) *gRPCService {
	return &gRPCService{
		Config:                    plugins.NewConfig(),
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

// SetConfig sets a manager configuration.
func (service *gRPCService) SetConfig(c config.Config) {
	service.Config.SetConfig(c)
}

// ServiceName returns the plug-in service name.
func (service *gRPCService) ServiceName() string {
	return "grpc"
}

// ServiceType returns the plug-in service type.
func (service *gRPCService) ServiceType() plugins.ServiceType {
	return plugins.SystemService
}

// SetPort sets a port number of the service.
func (service *gRPCService) SetPort(port int) {
	service.Port = port
}

// Start starts the service.
func (service *gRPCService) Start() error {
	var err error
	port, err := service.LookupServiceConfigPort(service)
	if err == nil {
		service.SetPort(port)
	}
	addr := net.JoinHostPort(service.Addr, strconv.Itoa(service.Port))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	service.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(loggingUnaryInterceptor))
	pb.RegisterStoreServer(service.grpcServer, service)
	pb.RegisterConfigServer(service.grpcServer, service)
	pb.RegisterHealthServer(service.grpcServer, service)
	pb.RegisterMetricServer(service.grpcServer, service)
	go func() {
		if err := service.grpcServer.Serve(listener); err != nil {
			log.Error(err)
		}
	}()

	log.Infof("gRPC server (%s) started", addr)

	return nil
}

// Stop stops the Grpc server.
func (service *gRPCService) Stop() error {
	if service.grpcServer != nil {
		service.grpcServer.Stop()
		service.grpcServer = nil
	}

	addr := net.JoinHostPort(service.Addr, strconv.Itoa(service.Port))
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

func (service *gRPCService) Check(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	res := pb.HealthCheckResponse{}        //nolint:exhaustruct
	switch service.Server.actor.Status() { //nolint:exhaustive
	case cluster.NodeUp:
		res.Status = pb.HealthCheckResponse_SERVING
	default:
		res.Status = pb.HealthCheckResponse_NOT_SERVING
	}
	return &res, nil
}

func (service *gRPCService) ListConfig(context.Context, *pb.ListConfigRequest) (*pb.ListConfigResponse, error) {
	res := pb.ListConfigResponse{} //nolint:exhaustruct
	res.Values = strings.Split(service.Server.Config.String(), "\n")
	return &res, nil
}

func (service *gRPCService) GetConfig(ctx context.Context, req *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	v, err := service.Config.LookupConfigObject(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "%s not found", req.GetName())
	}
	res := pb.GetConfigResponse{} //nolint:exhaustruct
	res.Value = fmt.Sprintf("%v", v)
	return &res, nil
}

func (service *gRPCService) GetVersion(context.Context, *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	res := pb.GetVersionResponse{} //nolint:exhaustruct
	res.Value = Version
	return &res, nil
}

func (service *gRPCService) ListMetric(context.Context, *pb.ListMetricRequest) (*pb.ListMetricResponse, error) {
	res := pb.ListMetricResponse{} //nolint:exhaustruct
	metrics, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return &res, err
	}
	for _, metric := range metrics {
		res.Values = append(res.GetValues(), metric.GetName())
	}
	return &res, nil
}

func (service *gRPCService) GetMetric(ctx context.Context, req *pb.GetMetricRequest) (*pb.GetMetricResponse, error) {
	res := pb.GetMetricResponse{} //nolint:exhaustruct
	metrics, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return &res, err
	}
	for _, metric := range metrics {
		if metric.GetName() != req.GetName() {
			continue
		}
		metricName := metric.GetName()
		for _, metric := range metric.GetMetric() {
			for _, label := range metric.GetLabel() {
				lableName := fmt.Sprintf("%s.%s", metricName, label.GetName())
				res.Names = append(res.GetNames(), lableName)
				res.Values = append(res.GetValues(), label.GetValue())
			}
		}
	}
	return nil, status.Errorf(codes.NotFound, "%s not found", req.GetName())
}

func (service *gRPCService) CreateDatabase(ctx context.Context, req *pb.CreateDatabaseRequest) (*pb.StatusResponse, error) {
	res := pb.StatusResponse{} //nolint:exhaustruct
	defaultStore, err := service.DefaultStoreService()
	if err != nil {
		return &res, err
	}
	err = defaultStore.CreateDatabase(pc.NewContext(), req.GetDatabaseName())
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func (service *gRPCService) RemoveDatabase(ctx context.Context, req *pb.RemoveDatabaseRequest) (*pb.StatusResponse, error) {
	res := pb.StatusResponse{} //nolint:exhaustruct
	defaultStore, err := service.DefaultStoreService()
	if err != nil {
		return &res, err
	}
	err = defaultStore.RemoveDatabase(pc.NewContext(), req.GetDatabaseName())
	if err != nil {
		return &res, err
	}
	return &res, nil
}

func (service *gRPCService) ListDatabases(context.Context, *pb.ListDatabasesRequest) (*pb.ListDatabasesResponse, error) {
	res := pb.ListDatabasesResponse{} //nolint:exhaustruct
	defaultStore, err := service.DefaultStoreService()
	if err != nil {
		return &res, err
	}
	dbs, err := defaultStore.ListDatabases(pc.NewContext())
	if err != nil {
		return &res, err
	}
	res.Databases = []string{}
	for _, db := range dbs {
		res.Databases = append(res.GetDatabases(), db.Name())
	}
	return &res, nil
}

func (service *gRPCService) ListCollections(context.Context, *pb.ListCollectionsRequest) (*pb.ListCollectionsResponse, error) {
	res := pb.ListCollectionsResponse{} //nolint:exhaustruct
	return &res, nil
}
