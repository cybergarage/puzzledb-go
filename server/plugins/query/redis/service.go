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

package redis

import (
	"github.com/cybergarage/go-redis/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Service represents a new MySQL service instance.
type Service struct {
	*redis.Server
	query.Service
}

// NewServiceWithStore returns a new MySQL service instance with the specified　Store.
func NewServiceWithStore(store store.Store) *Service {
	service := &Service{
		Server:  redis.NewServer(),
		Service: *query.NewService(),
	}
	service.Server.SetCommandHandler(service)
	return service
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
