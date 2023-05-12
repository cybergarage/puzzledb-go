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

package actor

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
)

// Status represents an actor status.
type Status int

const (
	StatusUnknown Status = iota
	StatusRunning
	StatusStopped
)

// Service represents a actor service.
type Service struct {
	plugins.Config
	coordinator.Service
	serviceStatus Status
}

// NewService returns a new actor service.
func NewService() *Service {
	return NewServiceWith(nil)
}

// NewServiceWith returns a new actor service with the specified coordinator.
func NewServiceWith(coordinator coordinator.Service) *Service {
	return &Service{
		Config:        plugins.NewConfig(),
		Service:       coordinator,
		serviceStatus: StatusStopped,
	}
}

// SetConfig sets a manager configuration.
func (service *Service) SetConfig(c config.Config) {
	service.Config.SetConfig(c)
}

// ServiceName returns the plug-in service name.
func (service *Service) ServiceName() string {
	return "actor"
}

// ServiceType returns the plug-in service type.
func (service *Service) ServiceType() plugins.ServiceType {
	return plugins.SystemService
}

// SetStatus sets a actor status.
func (service *Service) SetCoordinator(c coordinator.Service) { // nolint: stylecheck
	service.Service = c
}

// SetStatus sets a actor status.
func (service *Service) SetStatus(serviceStatus Status) { // nolint: stylecheck
	service.serviceStatus = serviceStatus
}

// Status returns a actor status.
func (service *Service) Status() Status {
	return service.serviceStatus
}

// Start starts the actor service.
func (service *Service) Start() error {
	return nil
}

// Stop stops the actor server.
func (service *Service) Stop() error {
	return nil
}
