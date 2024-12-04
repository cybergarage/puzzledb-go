// Copyright (C) 2020 The go-mysql Authors. All rights reserved.
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

package mysql

import (
	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/sql"
)

// Service represents a test server.
// This Service struct behave as ${hoge}CommandExecutor.
type Service struct {
	mysql.Server
	*sql.Service
}

// NewService returns a test server instance.
func NewService() query.Service {
	service := &Service{
		Server:  mysql.NewServer(),
		Service: sql.NewService(),
	}
	service.SetQueryExecutor(service)
	service.SetSQLExecutor(service.Service)
	return service
}

// ServiceName returns the plug-in service name.
func (service *Service) ServiceName() string {
	return "mysql"
}

// Start starts the service.
func (service *Service) Start() error {
	port, err := service.GetServiceConfigPort(service)
	if err == nil {
		service.SetPort(port)
	}
	tlsConfig, err := service.TLSConfig()
	if err != nil {
		service.Server.SetTLSEnabled(true)
		service.Server.SetTLSConfig(tlsConfig)
	} else {
		service.Server.SetTLSEnabled(false)
	}

	if err := service.Server.Start(); err != nil {
		return err
	}
	return nil
}

// Stop stops the service.
func (service *Service) Stop() error {
	if err := service.Server.Stop(); err != nil {
		return err
	}
	return nil
}
