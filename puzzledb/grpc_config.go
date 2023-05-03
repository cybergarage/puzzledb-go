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
	"strings"

	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/grpc"
)

func (server *GrpcServer) ListConfig(context.Context, *pb.ListConfigRequest) (*pb.ListConfigResponse, error) {
	res := pb.ListConfigResponse{} //nolint:exhaustruct
	res.Values = strings.Split(server.Config.String(), "\n")
	return &res, nil
}

func (server *GrpcServer) GetConfig(context.Context, *pb.GetConfigRequest) (*pb.GetConfigResponse, error) {
	res := pb.GetConfigResponse{} //nolint:exhaustruct
	return &res, nil
}
