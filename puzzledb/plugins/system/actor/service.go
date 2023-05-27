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
	"github.com/cybergarage/puzzledb-go/puzzledb/cluster"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
	coordinator_plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/coordinator"
)

// Service represents a actor service.
type Service struct {
	coordinator coordinator_plugin.Service
	plugins.Config
}

// NewService returns a new actor service.
func NewService() *Service {
	return &Service{
		coordinator: nil,
		Config:      plugins.NewConfig(),
	}
}

// SetConfig sets a manager configuration.
func (service *Service) SetConfig(c config.Config) {
	service.Config.SetConfig(c)
}

// SetCoordinator sets a coordinator service.
func (service *Service) SetCoordinator(coord coordinator_plugin.Service) {
	service.coordinator = coord
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
func (service *Service) SetStatus(status cluster.NodeStatus) error {
	err := service.coordinator.SetNodeState(service.coordinator)
	if err != nil {
		return err
	}
	service.coordinator.SetStatus(status)
	return nil
}

// Status returns a actor status.
func (service *Service) Status() cluster.NodeStatus {
	return service.coordinator.Status()
}

// Join joins the actor service to the cluster.
func (service *Service) Join() error {
	if err := service.SetStatus(cluster.NodeJoining); err != nil {
		return err
	}
	return nil
}

// Leave leaves the actor service from the cluster.
func (service *Service) Leave() error {
	if err := service.SetStatus(cluster.NodeLeaving); err != nil {
		return err
	}
	return nil
}

// Start starts the actor service.
func (service *Service) Start() error {
	if err := service.Join(); err != nil {
		return err
	}
	if err := service.SetStatus(cluster.NodeUp); err != nil {
		return err
	}
	return nil
}

// Stop stops the actor server.
func (service *Service) Stop() error {
	if err := service.Leave(); err != nil {
		return err
	}
	if err := service.SetStatus(cluster.NodeDown); err != nil {
		return err
	}
	return nil
}
