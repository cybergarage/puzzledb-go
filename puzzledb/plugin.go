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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/auth"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coder/key"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/metrics"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kv"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/store/kvcache"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer"
)

// PluginManager is a manager for plugins.
type PluginManager struct {
	*plugins.Manager
}

// NewPluginManager returns a new PluginManager.
func NewPluginManagerWith(mgr *plugins.Manager) *PluginManager {
	return &PluginManager{
		Manager: mgr,
	}
}

// RemoveDisabledServices removes disabled services.
func (mgr *PluginManager) RemoveDisabledServices(services []plugins.Service) []plugins.Service {
	enabledServices := []plugins.Service{}
	for _, service := range services {
		if mgr.IsServiceConfigEnabled(service) {
			enabledServices = append(enabledServices, service)
		}
	}
	return enabledServices
}

// KeyCoderServices returns key coder services.
func (mgr *PluginManager) KeyCoderServices() ([]key.Service, error) {
	services := []key.Service{}
	for _, service := range mgr.ServicesByType(plugins.CoderKeyService) {
		if s, ok := service.(key.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// DocumentCoderServices returns document coder services.
func (mgr *PluginManager) DocumentCoderServices() ([]document.Service, error) {
	services := []document.Service{}
	for _, service := range mgr.ServicesByType(plugins.CoderDocumentService) {
		if s, ok := service.(document.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// CoordinatorServices returns coordinator services.
func (mgr *PluginManager) CoordinatorServices() ([]coordinator.Service, error) {
	services := []coordinator.Service{}
	for _, service := range mgr.ServicesByType(plugins.CoordinatorService) {
		if s, ok := service.(coordinator.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// DocumentStoreServices returns document store services.
func (mgr *PluginManager) DocumentStoreServices() ([]store.Service, error) {
	services := []store.Service{}
	for _, service := range mgr.ServicesByType(plugins.StoreDocumentService) {
		if s, ok := service.(store.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// KvStoreServices returns KV store services.
func (mgr *PluginManager) KvStoreServices() ([]kv.Service, error) {
	services := []kv.Service{}
	for _, service := range mgr.ServicesByType(plugins.StoreKvService) {
		if s, ok := service.(kv.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// KvCacheStoreServices returns KV cache store services.
func (mgr *PluginManager) KvCacheStoreServices() ([]kvcache.Service, error) {
	services := []kvcache.Service{}
	for _, service := range mgr.ServicesByType(plugins.StoreKvCacheService) {
		if s, ok := service.(kvcache.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// QueryServices returns query services.
func (mgr *PluginManager) QueryServices() ([]query.Service, error) {
	services := []query.Service{}
	for _, service := range mgr.ServicesByType(plugins.QueryService) {
		if s, ok := service.(query.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// TracingServices returns tracing services.
func (mgr *PluginManager) TracingServices() ([]tracer.Service, error) {
	services := []tracer.Service{}
	for _, service := range mgr.ServicesByType(plugins.TracingService) {
		if s, ok := service.(tracer.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// AuthenticatorServices returns authenticator services.
func (mgr *PluginManager) AuthenticatorServices() ([]auth.Service, error) {
	services := []auth.Service{}
	for _, service := range mgr.ServicesByType(plugins.AuthenticatorService) {
		if s, ok := service.(auth.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.AuthenticatorService)
		}
	}
	return services, nil
}

// MetricsServices returns metrics services.
func (mgr *PluginManager) MetricsServices() ([]metrics.Service, error) {
	services := []metrics.Service{}
	for _, service := range mgr.ServicesByType(plugins.MetricsService) {
		if s, ok := service.(metrics.Service); ok {
			services = append(services, s)
		} else {
			return nil, newErrInvalidService(service, plugins.CoderKeyService)
		}
	}
	return services, nil
}

// DefaultCoordinatorService returns the default coordinator service.
func (mgr *PluginManager) DefaultCoordinatorService() (coordinator.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.CoordinatorService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	service, ok := defaultService.(coordinator.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.CoordinatorService)
	}
	return service, nil
}

// DefaultKeyCoderService returns the default key coder service.
func (mgr *PluginManager) DefaultKeyCoderService() (key.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.CoderKeyService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	service, ok := defaultService.(key.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.CoderKeyService)
	}
	return service, nil
}

// DefaultDocumentCoderService returns the default document coder service.
func (mgr *PluginManager) DefaultDocumentCoderService() (document.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.CoderDocumentService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	service, ok := defaultService.(document.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.CoderDocumentService)
	}
	return service, nil
}

// DefaultKvStoreService returns the default KV store service.
func (mgr *PluginManager) DefaultKvStoreService() (kv.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.StoreKvService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	service, ok := defaultService.(kv.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.StoreKvService)
	}
	return service, nil
}

// DefaultKvCacheStoreService returns the default KV cache store service.
func (mgr *PluginManager) DefaultKvCacheStoreService() (kvcache.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.StoreKvCacheService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	service, ok := defaultService.(kvcache.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.StoreKvCacheService)
	}
	return service, nil
}

// DefaultStoreService returns the default store service.
func (mgr *PluginManager) DefaultStoreService() (store.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.StoreDocumentService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	service, ok := defaultService.(store.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.StoreDocumentService)
	}
	return service, nil
}

// DefaultTracingService returns the default tracing service.
func (mgr *PluginManager) DefaultTracingService() (tracer.Service, error) {
	defaultService, err := mgr.DefaultService(plugins.TracingService)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	service, ok := defaultService.(tracer.Service)
	if !ok {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.TracingService)
	}
	return service, nil
}

// DefaultAuthenticatorService returns the default authenticator service.
func (mgr *PluginManager) DefaultAuthenticatorService() (auth.Service, error) {
	services := []auth.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.AuthenticatorService) {
		if s, ok := service.(auth.Service); ok {
			services = append(services, s)
		}
	}
	if len(services) == 0 {
		return nil, plugins.NewErrDefaultServiceNotFound(plugins.AuthenticatorService)
	}
	return services[len(services)-1], nil
}

// EnabledKeyCoderServices returns enabled key coder services.
func (mgr *PluginManager) EnabledKeyCoderServices() []key.Service {
	services := []key.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.CoderKeyService) {
		if s, ok := service.(key.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

// EnabledDocumentCoderServices returns enabled document coder services.
func (mgr *PluginManager) EnabledDocumentCoderServices() []document.Service {
	services := []document.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.CoderDocumentService) {
		if s, ok := service.(document.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

// EnabledCoordinatorServices returns enabled coordinator services.
func (mgr *PluginManager) EnabledCoordinatorServices() []coordinator.Service {
	services := []coordinator.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.CoordinatorService) {
		if s, ok := service.(coordinator.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

// EnabledDocumentStoreServices returns enabled document store services.
func (mgr *PluginManager) EnabledDocumentStoreServices() []store.Service {
	services := []store.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.StoreDocumentService) {
		if s, ok := service.(store.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

// EnabledKvStoreServices returns enabled KV store services.
func (mgr *PluginManager) EnabledKvStoreServices() []kv.Service {
	services := []kv.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.StoreKvService) {
		if s, ok := service.(kv.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

// EnabledKvCacheStoreServices returns enabled KV cache store services.
func (mgr *PluginManager) EnabledKvCacheStoreServices() []kvcache.Service {
	services := []kvcache.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.StoreKvCacheService) {
		if s, ok := service.(kvcache.Service); ok {
			services = append(services, s)
		}
	}
	return services
}

// EnabledTracingServices returns enabled tracing services.
func (mgr *PluginManager) EnabledTracingServices() []tracer.Service {
	services := []tracer.Service{}
	for _, service := range mgr.EnabledServicesByType(plugins.TracingService) {
		if s, ok := service.(tracer.Service); ok {
			services = append(services, s)
		}
	}
	return services
}
