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

// Service represents a new MySQL service instance.
type Service struct {
	*postgresql.Server
	*sql.Service
}

// NewService returns a new MySQL service.
func NewService() query.Service {
	service := &Service{
		Server:  postgresql.NewServer(),
		Service: sql.NewService(),
	}
	service.SetTransactionExecutor(service)
	service.SetQueryExecutor(service)
	service.SetBulkExecutor(service)
	service.SetErrorHandler(service)
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
	if err == nil && tlsConfig != nil {
		service.Server.SetServerKeyFile(tlsConfig.TLSKeyFile())
		service.Server.SetServerCertFile(tlsConfig.TLSCertFile())
		service.Server.SetRootCertFiles(tlsConfig.TLSCAFiles()...)
		service.Server.SetTLSEnabled(true)
		_, err := tlsConfig.TLSConfig()
		if err != nil {
			return err
		}
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
