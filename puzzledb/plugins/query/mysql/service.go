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

package mysql

import (
	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Service represents a new MySQL service instance.
type Service struct {
	*mysql.BaseExecutor
	*mysql.Server
	*query.BaseService
}

// NewService returns a new MySQL service.
func NewService() query.Service {
	srv := &Service{
		BaseExecutor: mysql.NewBaseExecutor(),
		Server:       mysql.NewServer(),
		BaseService:  query.NewService(),
	}
	srv.Server.SetQueryExecutor(srv)
	return srv
}

// Type returns the plug-in service type.
func (service *Service) Type() plugins.ServiceType {
	return plugins.QueryService
}

// Name returns the plug-in service name.
func (service *Service) Name() string {
	return "mysql"
}

// Start starts the service.
func (service *Service) Start() error {
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

// Stop stops the service.
func (service *Service) CancelTransactionWithError(txn store.Transaction, err error) error {
	if txErr := txn.Cancel(); txErr != nil {
		return txErr
	}
	return err
}
