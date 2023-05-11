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
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
)

// ServiceStatus represents a server status.
type ServiceStatus int

const (
	ServiceStatusUnknown ServiceStatus = iota
	ServiceStatusRunning
	ServiceStatusStopped
)

// StatusService represents a server status service.
type StatusService struct {
	coordinator.Service
	serviceStatus ServiceStatus
}

// NewStatuServiceWith returns a new server status service.
func NewStatuServiceWith(coordinator coordinator.Service) *StatusService {
	return &StatusService{
		Service:       coordinator,
		serviceStatus: ServiceStatusStopped,
	}
}

// SetStatus sets a server status.
func (status *StatusService) SetStatus(serviceStatus ServiceStatus) { // nolint: stylecheck
	status.serviceStatus = serviceStatus
}

// Status returns a server status.
func (status *StatusService) Status() ServiceStatus {
	return status.serviceStatus
}

// Start starts the service.
func (service *StatusService) Start() error {
	return nil
}

// Stop stops the Grpc server.
func (service *StatusService) Stop() error {
	return nil
}
