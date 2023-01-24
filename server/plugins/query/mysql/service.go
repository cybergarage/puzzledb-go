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
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins/query"
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

// Start starts the service.
func (srv *Service) Start() error {
	if err := srv.Server.Start(); err != nil {
		return err
	}
	return nil
}

// Stop stops the service.
func (srv *Service) Stop() error {
	if err := srv.Server.Stop(); err != nil {
		return err
	}
	return nil
}
