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

package puzzledb

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/errors"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
)

type PluginManager struct {
	*plugins.Manager
}

func NewPluginManagerWith(mgr *plugins.Manager) *PluginManager {
	return &PluginManager{
		Manager: mgr,
	}
}

func (mgr *PluginManager) RemoveDisabledServices(services []plugins.Service) []plugins.Service {
	enabledServices := []plugins.Service{}
	for _, service := range services {
		if mgr.IsEnabled(service) {
			enabledServices = append(enabledServices, service)
		}
	}
	return enabledServices
}

func (mgr *PluginManager) KeyCoderServices() []key.Service {
	services := []key.Service{}
	for _, service := range mgr.ServicesByType(plugins.CoderKeyService) {
		if s, ok := service.(key.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) DocumentCoderServices() []document.Service {
	services := []document.Service{}
	for _, service := range mgr.ServicesByType(plugins.CoderDocumentService) {
		if s, ok := service.(document.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) DocumentStoreServices() []store.Service {
	services := []store.Service{}
	for _, service := range mgr.ServicesByType(plugins.StoreDocumentService) {
		if s, ok := service.(store.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) KvStoreServices() []kv.Service {
	services := []kv.Service{}
	for _, service := range mgr.ServicesByType(plugins.StoreKvService) {
		if s, ok := service.(kv.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) QueryServices() []query.Service {
	services := []query.Service{}
	for _, service := range mgr.ServicesByType(plugins.QueryService) {
		if s, ok := service.(query.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) DefaultKeyCoderService() (key.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.CoderKeyService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	defaultKeyCoder, ok := defaultService.(key.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.CoderKeyService)
	}
	return defaultKeyCoder, nil
}

func (mgr *PluginManager) DefaultDocumentCoderService() (document.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.CoderDocumentService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	defaultDocCoder, ok := defaultService.(document.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.CoderDocumentService)
	}
	return defaultDocCoder, nil
}

func (mgr *PluginManager) DefaultKvStoreService() (kv.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.StoreKvService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	defaultKvStore, ok := defaultService.(kv.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.StoreKvService)
	}
	return defaultKvStore, nil
}

func (mgr *PluginManager) DefaultStoreService() (store.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.StoreDocumentService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	defaultDocStore, ok := defaultService.(store.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.StoreDocumentService)
	}
	return defaultDocStore, nil
}

func (mgr *PluginManager) EnabledKeyCoderServices() []key.Service {
	services := []key.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.CoderKeyService) {
		if s, ok := service.(key.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) EnabledDocumentCoderServices() []document.Service {
	services := []document.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.CoderDocumentService) {
		if s, ok := service.(document.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) EnabledDocumentStoreServices() []store.Service {
	services := []store.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.StoreDocumentService) {
		if s, ok := service.(store.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

func (mgr *PluginManager) EnabledKvStoreServices() []kv.Service {
	services := []kv.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.StoreKvService) {
		if s, ok := service.(kv.Service); ok {
			services = append(services, s)
		}
	}
	return services
}
