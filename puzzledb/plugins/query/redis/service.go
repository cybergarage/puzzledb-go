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
	"strconv"

	"github.com/cybergarage/go-redis/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Service represents a new Redis service instance.
type Service struct {
	*redis.Server
	*query.BaseService
}

// NewService returns a new Redis service instance.
func NewService() Service {
	service := &Service{
		Server:      redis.NewServer(),
		BaseService: query.NewBaseService(),
	}
	service.Server.SetCommandHandler(service)
	return service
}

// ServiceName returns the plug-in service name.
func (service *Service) ServiceName() string {
	return "redis"
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

// GetDatabase returns the database with the specified ID.
func (service *Service) GetDatabase(ctx context.Context, id int) (store.Database, error) {
	store := service.Store()
	name := strconv.Itoa(id)
	db, err := store.GetDatabase(ctx, name)
	if err == nil {
		return db, nil
	}
	err = store.CreateDatabase(ctx, name)
	if err != nil {
		return nil, err
	}
	return store.GetDatabase(ctx, name)
}
