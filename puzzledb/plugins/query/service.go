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
	"github.com/cybergarage/go-tracing/tracer"
	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// StoreService represents a store interface for query service.
type StoreService interface {
	// SetStore sets the store.
	SetStore(store store.Store)
	// Store returns the store.
	Store() store.Store
}

// CoordinatorService represents a coordinator interface for query service.
type CoordinatorService interface {
	// SetCoordinator sets the coordinator.
	SetCoordinator(coordinator coordinator.Coordinator)
	// Coordinator returns the coordinator.
	Coordinator() coordinator.Coordinator
	// Observer is an interface to receive a message from the coordinator.
	coordinator.Observer
	// PostDatabaseCreateMessage posts a create database message to the coordinator.
	PostDatabaseCreateMessage(database string) error
	// PostDatabaseDeleteMessage posts a update database message to the coordinator.
	PostDatabaseUpdateMessage(database string) error
	// PostDatabaseDropMessage posts a drop database message to the coordinator.
	PostDatabaseDropMessage(database string) error
	// PostCollectionCreateMessage posts a create collection message to the coordinator.
	PostCollectionCreateMessage(database string, collection string) error
	// PostCollectionDeleteMessage posts a update collection message to the coordinator.
	PostCollectionUpdateMessage(database string, collection string) error
	// PostCollectionDropMessage posts a drop collection message to the coordinator.
	PostCollectionDropMessage(database string, collection string) error
}

// Service represents a query service.
type Service interface {
	plugins.Service
	CoordinatorService
	StoreService
	// SetTracer sets the tracing tracer.
	SetTracer(t tracer.Tracer)
	// SetPort sets the listen port.
	SetPort(port int)
}
