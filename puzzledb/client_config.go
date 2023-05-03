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

func (client *Client) GetConfig(name string) (string, error) {
	c := pb.NewConfigClient(client.Conn)
	req := &pb.GetConfigRequest{
		Name: name,
	}
	res, err := c.GetConfig(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.Value, nil
}

func (client *Client) ListConfig() ([]string, error) {
	c := pb.NewConfigClient(client.Conn)
	res, err := c.ListConfig(context.Background(), &pb.ListConfigRequest{})
	if err != nil {
		return []string{}, err
	}
	return res.Values, nil
}
