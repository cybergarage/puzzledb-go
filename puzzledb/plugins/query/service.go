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

package query

import (
	"fmt"

	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type Service interface {
	plugins.Service
	SetStore(store plugins.Service) error
	Store() store.Store
}

type BaseService struct {
	store store.Store
}

func NewService() *BaseService {
	server := &BaseService{
		store: nil,
	}
	return server
}

// ServiceType returns the plug-in service type.
func (service *BaseService) ServiceType() plugins.ServiceType {
	return plugins.QueryService
}

func (service *BaseService) SetStore(v plugins.Service) error {
	store, ok := v.(store.Store)
	if !ok {
		return plugins.NewErrInvalid(fmt.Sprintf("%s (%s)", v.ServiceName(), v.ServiceType().String()))
	}
	service.store = store
	return nil
}

func (service *BaseService) Store() store.Store {
	return service.store
}
