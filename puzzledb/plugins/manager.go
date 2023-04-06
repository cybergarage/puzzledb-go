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
	services []Service
}

// NewManager returns a plug-in manager instance.
func NewManager() *Manager {
	return &Manager{
		services: []Service{},
	}
}

// Add adds a service.
func (mgr *Manager) Add(srv Service) {
	mgr.services = append(mgr.services, srv)
}

// Reload reloads all services.
func (mgr *Manager) Reload(srvs []Service) {
	mgr.services = srvs
}

// Start starts all services.
func (mgr *Manager) Start() error {
	log.Infof("plug-ins loading...")

	for _, srv := range mgr.services {
		if err := srv.Start(); err != nil {
			if err := mgr.Stop(); err != nil {
				return err
			}
			return err
		}
		log.Infof("%s (%s) loaded", srv.ServiceName(), srv.ServiceType().String())
	}

	log.Infof("plug-ins loaded")

	for _, s := range strings.Split(mgr.String(), "\n") {
		log.Infof("%s", s)
	}

	return nil
}

// Stop stops all services.
func (mgr Manager) Stop() error {
	log.Infof("plug-ins terminating...")
	var lastErr error
	for _, srv := range mgr.services {
		if err := srv.Stop(); err != nil {
			lastErr = err
		}
	}
	log.Infof("plug-ins terminated")
	return lastErr
}

// String returns a string representation of the plug-in manager.
func (mgr *Manager) String() string {
	var s string
	for _, servieType := range ServiceTypes() {
		names := []string{}
		for _, service := range mgr.services {
			if service.ServiceType() == servieType {
				names = append(names, service.ServiceName())
			}
		}
		if len(names) == 0 {
			continue
		}
		s += fmt.Sprintf("- %s (%s)\n", servieType.String(), strings.Join(names, ", "))
	}
	return strings.TrimSuffix(s, "\n")
}
