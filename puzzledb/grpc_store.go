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

	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/api"
)

func (server *GrpcServer) ListDatabases(context.Context, *pb.ListDatabasesRequest) (*pb.ListDatabasesResponse, error) {
	res := pb.ListDatabasesResponse{} //nolint:exhaustruct
	defaultStore, err := server.DefaultStoreService()
	if err != nil {
		return &res, err
	}
	dbs, err := defaultStore.ListDatabases()
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
