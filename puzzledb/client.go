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

	"google.golang.org/grpc"
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
func (client *Client) Open(host string) error {
	addr := net.JoinHostPort(client.Host, strconv.Itoa(client.Port))
	conn, err := grpc.Dial(addr)
	if err != nil {
		return err
	}
	client.Conn = conn
	return nil
}

func (client *Client) Execute(args []string) error {
	return nil
}
