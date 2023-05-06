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
	"strings"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/grpc"
)

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
