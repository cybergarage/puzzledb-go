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

package plugins

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

type Manager struct {
	*plugins.Manager
}

func NewManager() *Manager {
	server := puzzledbtest.NewServer()
	server.LoadPlugins() //nolint:errcheck
	return &Manager{
		Manager: server.Manager,
	}
}

func (mgr *Manager) EnabledKeyCoderServices() []key.Service {
	services := []key.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.CoderKeyService) {
		if s, ok := service.(key.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *Manager) EnabledDocumentCoderServices() []document.Service {
	services := []document.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.CoderDocumentService) {
		if s, ok := service.(document.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *Manager) EnabledDocumentStoreServices() []store.Service {
	services := []store.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.StoreDocumentService) {
		if s, ok := service.(store.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *Manager) EnabledKvStoreServices() []kv.Service {
	services := []kv.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.StoreKvService) {
		if s, ok := service.(kv.Service); ok {
			services = append(services, s)
		}
	}
	return services
}
