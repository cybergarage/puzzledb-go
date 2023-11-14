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

	pb "github.com/cybergarage/puzzledb-go/puzzledb/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client represens a gRPC client.
type Client struct {
	Host string
	Port int
	Conn *grpc.ClientConn
}

// NewClient returns a new gRPC client.
func NewClient() *Client {
	client := &Client{
		Host: "",
		Port: DefaultGrpcPort,
		Conn: nil,
	}
	return client
}

// SetPort sets a port number.
func (client *Client) SetPort(port int) {
	client.Port = port
}

// SetHost sets a host name.
func (client *Client) SetHost(host string) {
	client.Host = host
}

// Open opens a gRPC connection.
func (client *Client) Open() error {
	addr := net.JoinHostPort(client.Host, strconv.Itoa(client.Port))
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	client.Conn = conn
	return nil
}

func (client *Client) Close() error {
	if client.Conn == nil {
		return nil
	}
	err := client.Conn.Close()
	if err != nil {
		return err
	}
	client.Conn = nil
	return nil
}

func (client *Client) Check() (bool, error) {
	c := pb.NewHealthClient(client.Conn)
	req := &pb.HealthCheckRequest{} //nolint:exhaustruct
	res, err := c.Check(context.Background(), req)
	if err != nil {
		return false, err
	}
	if res.GetStatus() != pb.HealthCheckResponse_SERVING {
		return false, nil
	}
	return true, nil
}

func (client *Client) GetVersion() (string, error) {
	c := pb.NewConfigClient(client.Conn)
	req := &pb.GetVersionRequest{}
	res, err := c.GetVersion(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.GetValue(), nil
}

func (client *Client) GetConfig(name string) (string, error) {
	c := pb.NewConfigClient(client.Conn)
	req := &pb.GetConfigRequest{
		Name: name,
	}
	res, err := c.GetConfig(context.Background(), req)
	if err != nil {
		return "", err
	}
	return res.GetValue(), nil
}

func (client *Client) ListConfig() ([]string, error) {
	c := pb.NewConfigClient(client.Conn)
	res, err := c.ListConfig(context.Background(), &pb.ListConfigRequest{})
	if err != nil {
		return []string{}, err
	}
	return res.GetValues(), nil
}

func (client *Client) GetMetric(name string) (map[string]string, error) {
	c := pb.NewMetricClient(client.Conn)
	req := &pb.GetMetricRequest{
		Name: name,
	}
	res, err := c.GetMetric(context.Background(), req)
	if err != nil {
		return nil, err
	}
	metricsMap := make(map[string]string)
	values := res.GetValues()
	for n, name := range res.GetNames() {
		if len(values) <= n {
			metricsMap[name] = ""
			continue
		}
		metricsMap[name] = values[n]
	}
	return metricsMap, nil
}

func (client *Client) ListMetric() ([]string, error) {
	c := pb.NewMetricClient(client.Conn)
	res, err := c.ListMetric(context.Background(), &pb.ListMetricRequest{})
	if err != nil {
		return []string{}, err
	}
	return res.GetValues(), nil
}

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
	return res.GetDatabases(), nil
}
