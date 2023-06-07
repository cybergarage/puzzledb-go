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
	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	docStore "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kvcache"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type BaseService struct {
	plugins.Config
	coordinator coordinator.Coordinator
	store       store.Store
}

// NewBaseService returns a new query base service.
func NewBaseService() *BaseService {
	server := &BaseService{
		Config:      plugins.NewConfig(),
		store:       nil,
		coordinator: nil,
	}
	return server
}

// ServiceType returns the plug-in service type.
func (service *BaseService) ServiceType() plugins.ServiceType {
	return plugins.QueryService
}

// SetCoordinator sets the coordinator.
func (service *BaseService) SetCoordinator(coordinator coordinator.Coordinator) {
	service.coordinator = coordinator
}

// Coordinator returns the coordinator.
func (service *BaseService) Coordinator() coordinator.Coordinator {
	return service.coordinator
}

// PostSchemaMessage posts a schema message to the coordinator.
func (service *BaseService) PostSchemaMessage(key document.Key, e coordinator.EventType) error {
	obj, err := NewSchemaMessageObjectWith(key)
	if err != nil {
		return err
	}
	msg, err := coordinator.NewMessageWith(
		coordinator.CollectionMessage,
		e,
		obj,
	)
	if err != nil {
		return err
	}
	return service.coordinator.PostMessage(msg)
}

// OnMessageReceived is called when a message is received from the coordinator.
func (service *BaseService) OnMessageReceived(msg coordinator.Message) {
	switch msg.Type() { // nolint:gocritic, exhaustive
	case coordinator.CollectionMessage:
		store := service.Store()
		kvStore, ok := store.(kvcache.CacheStore)
		if !ok {
			return
		}
		var msgObj CollectionMessageObject
		if err := msg.UnmarshalTo(&msgObj); err != nil {
			log.Error(err)
			return
		}
		switch msg.Event() {
		case coordinator.CreatedEvent:
		case coordinator.UpdatedEvent, coordinator.DeletedEvent:
			if err := kvStore.EraseCollectionCache(msgObj.Database, msgObj.Collection); err != nil {
				log.Error(err)
			}
		}
	}
}

// SetStore sets the store.
func (service *BaseService) SetStore(store store.Store) {
	service.store = store

	// Register the cache key prefix for the database and collection.
	kvDocStore, ok := store.(docStore.DocumentKvStore)
	if !ok {
		return
	}
	kvStore := kvDocStore.KvStore()
	cacheStore, ok := kvStore.(kvcache.CacheStore)
	if !ok {
		return
	}
	cacheStore.EnableDatabaseCache()
	cacheStore.EnableCollectionCache()
}

// Store returns the store.
func (service *BaseService) Store() store.Store {
	return service.store
}
