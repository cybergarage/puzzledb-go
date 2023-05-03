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

	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/grpc"
)

// CreateDatabase creates a specified database.
func (client *Client) CreateDatabase(name string) error {
	c := pb.NewStoreClient(client.Conn)
	req := &pb.CreateDatabaseRequest{
		DatabaseName: name,
	}
	_, err := c.CreateDatabase(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

// RemoveDatabase removes a specified database.
func (client *Client) RemoveDatabase(name string) error {
	c := pb.NewStoreClient(client.Conn)
	req := &pb.RemoveDatabaseRequest{
		DatabaseName: name,
	}
	_, err := c.RemoveDatabase(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

// ListDatabases returns a list of database names.
func (client *Client) ListDatabases() ([]string, error) {
	c := pb.NewStoreClient(client.Conn)
	res, err := c.ListDatabases(context.Background(), &pb.ListDatabasesRequest{})
	if err != nil {
		return []string{}, err
	}
	return res.Databases, nil
}
