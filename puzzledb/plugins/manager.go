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
	"fmt"
	"strings"

	"github.com/cybergarage/go-logger/log"
)

// Manager represents a plug-in manager.
type Manager struct {
	Config
	services []Service
}

// NewManager returns a plug-in manager instance.
func NewManager() *Manager {
	return &Manager{
		Config:   NewConfig(),
		services: []Service{},
	}
}

// RegisterService adds a plug-in service.
func (mgr *Manager) RegisterService(service Service) {
	mgr.services = append(mgr.services, service)
}

// ReloadServices reloads all plug-in services.
func (mgr *Manager) ReloadServices(services []Service) {
	mgr.services = services
}

// Services returns all registered plug-in services.
func (mgr *Manager) Services() []Service {
	return mgr.services
}

// ServicesByType returns all registered plug-in services with the specified type.
func (mgr *Manager) ServicesByType(t ServiceType) []Service {
	services := []Service{}
	for _, service := range mgr.services {
		if service.ServiceType() == t {
			services = append(services, service)
		}
	}
	return services
}

// EnabledServicesByType returns all enabled plug-in services with the specified type.
func (mgr *Manager) EnabledServicesByType(t ServiceType) []Service {
	services := []Service{}
	for _, service := range mgr.ServicesByType(t) {
		if !mgr.IsServiceConfigEnabled(service) {
			continue
		}
		services = append(services, service)
	}
	return services
}

// DefaultService returns the default plug-in service with the specified type.
func (mgr *Manager) DefaultService(t ServiceType) (Service, error) {
	if !t.IsExclusive() {
		return nil, NewErrServiceNotFound(t)
	}
	services := mgr.ServicesByType(t)
	if len(services) == 0 {
		return nil, NewErrServiceNotFound(t)
	}
	lastIdx := len(services) - 1
	if mgr.Config == nil {
		return services[lastIdx], nil
	}
	configName, err := mgr.GetConfigString(ConfigPlugins, t.String(), ConfigDefault)
	if err != nil {
		return services[lastIdx], nil //nolint:nilerr
	}
	for _, service := range services {
		if service.ServiceName() != configName {
			continue
		}
		if !mgr.IsServiceConfigEnabled(service) {
			return nil, NewErrDisabledService(service)
		}
		return service, nil
	}
	return nil, NewErrNotFound(fmt.Sprintf("%s (%s)", configName, t.String()))
}

// Start starts all plug-in services.
func (mgr *Manager) Start() error {
	log.Infof("plug-ins loading...")

	for _, service := range mgr.services {
		service.SetConfig(mgr.Config.Object())
		if !mgr.IsServiceConfigEnabled(service) {
			log.Infof("%s (%s) skipped", service.ServiceName(), service.ServiceType().String())
			continue
		}
		if err := service.Start(); err != nil {
			if err := mgr.Stop(); err != nil {
				return err
			}
			return err
		}
		log.Infof("%s (%s) started", service.ServiceName(), service.ServiceType().String())
	}

	log.Infof("plug-ins loaded")

	return nil
}

// Stop stops all plug-in services.
func (mgr Manager) Stop() error {
	log.Infof("plug-ins terminating...")
	var lastErr error
	for _, service := range mgr.services {
		if !mgr.IsServiceConfigEnabled(service) {
			log.Infof("%s (%s) skipped", service.ServiceName(), service.ServiceType().String())
			continue
		}
		if err := service.Stop(); err != nil {
			lastErr = err
		}
		log.Infof("%s (%s) terminated", service.ServiceName(), service.ServiceType().String())
	}
	log.Infof("plug-ins terminated")
	return lastErr
}

// String returns a string representation of the plug-in manager.
func (mgr *Manager) String() string {
	var s string
	for _, servieType := range ServiceTypes() {
		defaultService, _ := mgr.DefaultService(servieType)
		names := []string{}
		for _, service := range mgr.services {
			if service.ServiceType() != servieType {
				continue
			}
			name := service.ServiceName()
			serviceStatus := "-"
			if mgr.IsServiceConfigEnabled(service) {
				serviceStatus = "+"
				if defaultService != nil {
					if name == defaultService.ServiceName() {
						serviceStatus = "*"
					}
				}
			}
			name = fmt.Sprintf("%s%s", name, serviceStatus)
			names = append(names, name)
		}
		if len(names) == 0 {
			continue
		}
		s += fmt.Sprintf("- %s (%s)\n", servieType.String(), strings.Join(names, ", "))
	}
	return strings.TrimSuffix(s, "\n")
}
