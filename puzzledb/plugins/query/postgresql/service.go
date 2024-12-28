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

package postgresql

import (
	"github.com/cybergarage/go-postgresql/postgresql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query/sql"
)

// Service represents a new PostgreSQL service instance.
type Service struct {
	postgresql.Server
	*sql.Service
}

// NewService returns a new PostgreSQL service.
func NewService() query.Service {
	service := &Service{
		Server:  postgresql.NewServer(),
		Service: sql.NewService(),
	}
	service.SetQueryExecutor(service)
	service.SetBulkQueryExecutor(service)
	service.SetErrorHandler(service)
	service.SetSQLExecutor(service.Service)
	return service
}

// ServiceName returns the plug-in service name.
func (service *Service) ServiceName() string {
	return "postgresql"
}

// Start starts the service.
func (service *Service) Start() error {
	port, err := service.GetServiceConfigPort(service)
	if err == nil {
		service.SetPort(port)
	}
	tlsConfig, err := service.TLSConfig()
	if err != nil {
		service.Server.SetTLSConfig(tlsConfig)
	} else {
		service.Server.SetTLSConfig(nil)
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
