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

// ActorStatus represents an actor status.
type ActorStatus int

const (
	ActorStatusUnknown ActorStatus = iota
	ActorStatusRunning
	ActorStatusStopped
)

// ActorService represents a actor service.
type ActorService struct {
	coordinator.Service
	serviceStatus ActorStatus
}

// NewActorServiceWith returns a new actor service.
func NewActorServiceWith(coordinator coordinator.Service) *ActorService {
	return &ActorService{
		Service:       coordinator,
		serviceStatus: ActorStatusStopped,
	}
}

// SetStatus sets a actor status.
func (status *ActorService) SetStatus(serviceStatus ActorStatus) { // nolint: stylecheck
	status.serviceStatus = serviceStatus
}

// Status returns a actor status.
func (status *ActorService) Status() ActorStatus {
	return status.serviceStatus
}

// Start starts the actor service.
func (service *ActorService) Start() error {
	return nil
}

// Stop stops the actor server.
func (service *ActorService) Stop() error {
	return nil
}
