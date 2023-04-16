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
	*ManagerConfig
	services []Service
}

// NewManager returns a plug-in manager instance.
func NewManager() *Manager {
	return &Manager{
		ManagerConfig: nil,
		services:      []Service{},
	}
}

// SetConfig sets the manager configuration.
func (mgr *Manager) SetConfig(config Config) {
	mgr.ManagerConfig = NewManagerConfigWith((config))
}

// RegisterService adds a plug-in service.
func (mgr *Manager) RegisterService(srv Service) {
	mgr.services = append(mgr.services, srv)
}

// ReloadServices reloads all plug-in services.
func (mgr *Manager) ReloadServices(srvs []Service) {
	mgr.services = srvs
}

func (mgr *Manager) IsEnabled(s Service) bool {
	configPath := strings.Join([]string{configPlugins, s.ServiceType().String(), s.ServiceName(), configEnabled}, ".")
	enabled, err := mgr.Config.GetBool(configPath)
	if err != nil {
		return true
	}
	return enabled
}

// Services returns all registered plug-in services.
func (mgr *Manager) Services() []Service {
	return mgr.services
}

// ServicesByType returns all registered plug-in services with the specified type.
func (mgr *Manager) ServicesByType(t ServiceType) []Service {
	services := []Service{}
	for _, srv := range mgr.services {
		if srv.ServiceType() == t {
			services = append(services, srv)
		}
	}
	return services
}

// EnabledServicesByType returns all enabled plug-in services with the specified type.
func (mgr *Manager) EnabledServicesByType(t ServiceType) []Service {
	services := []Service{}
	for _, srv := range mgr.ServicesByType(t) {
		if !mgr.IsEnabled(srv) {
			continue
		}
		services = append(services, srv)
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
	configPath := strings.Join([]string{configPlugins, t.String(), configDefault}, ".")
	configName, err := mgr.Config.GetString(configPath)
	if err != nil {
		return services[lastIdx], nil //nolint:nilerr
	}
	for _, srv := range services {
		if srv.ServiceName() != configName {
			continue
		}
		if !mgr.IsEnabled(srv) {
			return nil, NewErrDisabledService(srv)
		}
		return srv, nil
	}
	return nil, NewErrNotFound(fmt.Sprintf("%s (%s)", configName, t.String()))
}

// Start starts all plug-in services.
func (mgr *Manager) Start() error {
	log.Infof("plug-ins loading...")

	for _, srv := range mgr.services {
		if err := srv.Start(); err != nil {
			if !mgr.IsEnabled(srv) {
				log.Infof("%s (%s) skipped", srv.ServiceName(), srv.ServiceType().String())
				continue
			}
			if err := mgr.Stop(); err != nil {
				return err
			}
			return err
		}
		log.Infof("%s (%s) started", srv.ServiceName(), srv.ServiceType().String())
	}

	log.Infof("plug-ins loaded")

	return nil
}

// Stop stops all plug-in services.
func (mgr Manager) Stop() error {
	log.Infof("plug-ins terminating...")
	var lastErr error
	for _, srv := range mgr.services {
		if !mgr.IsEnabled(srv) {
			log.Infof("%s (%s) skipped", srv.ServiceName(), srv.ServiceType().String())
			continue
		}
		if err := srv.Stop(); err != nil {
			lastErr = err
		}
		log.Infof("%s (%s) terminated", srv.ServiceName(), srv.ServiceType().String())
	}
	log.Infof("plug-ins terminated")
	return lastErr
}

// String returns a string representation of the plug-in manager.
func (mgr *Manager) String() string {
	var s string
	for _, servieType := range ServiceTypes() {
		defaultService, err := mgr.DefaultService(servieType)
		names := []string{}
		for _, service := range mgr.services {
			if service.ServiceType() != servieType {
				continue
			}
			name := service.ServiceName()
			if err == nil {
				if name == defaultService.ServiceName() {
					name = fmt.Sprintf("*%s", name)
				}
			}
			if !mgr.IsEnabled(service) {
				name = fmt.Sprintf("-%s", name)
			}
			names = append(names, name)
		}
		if len(names) == 0 {
			continue
		}
		s += fmt.Sprintf("- %s (%s)\n", servieType.String(), strings.Join(names, ", "))
	}
	return strings.TrimSuffix(s, "\n")
}
