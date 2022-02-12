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

// Service represents a plugin services.
type Services struct {
	services []Service
}

// NewService returns a new MySQL service instance.
func NewServices() *Services {
	return &Services{
		services: []Service{},
	}

}

// Add adds a service.
func (srvs *Services) Add(srv Service) {
	srvs.services = append(srvs.services, srv)
}

// Start starts all services.
func (srvs *Services) Start() error {
	for _, srv := range srvs.services {
		if err := srv.Start(); err != nil {
			srvs.Stop()
			return err
		}
	}
	return nil
}

// Stop stops all services.
func (srvs Services) Stop() error {
	for _, srv := range srvs.services {
		if err := srv.Stop(); err != nil {
			return err
		}
	}
	return nil
}
