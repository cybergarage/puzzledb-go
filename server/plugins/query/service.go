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
	"github.com/cybergarage/puzzledb-go/puzzledb/record"
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type Service interface {
	plugins.Service
	SetStore(store store.Store)
	SetSerializer(serializer record.Serializer)
	Store() store.Store
	Serializer() record.Serializer
}

type BaseService struct {
	store      store.Store
	serializer record.Serializer
}

func NewService() *BaseService {
	server := &BaseService{
		store:      nil,
		serializer: nil,
	}
	return server
}

func (service *BaseService) SetStore(store store.Store) {
	service.store = store
}

func (service *BaseService) SetSerializer(serializer record.Serializer) {
	service.serializer = serializer
}

func (service *BaseService) Store() store.Store {
	return service.store
}

func (service *BaseService) Serializer() record.Serializer {
	return service.serializer
}
