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

package redis

import (
	"github.com/cybergarage/go-redis/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
)

// Service represents a new Redis service instance.
type Service struct {
	*redis.Server
	*query.BaseService
}

// NewService returns a new Redis service instance.
func NewService() query.Service {
	service := &Service{
		Server:      redis.NewServer(),
		BaseService: query.NewBaseService(),
	}
	service.Server.SetCommandHandler(service)
	service.Server.SetAuthCommandHandler(service)
	return service
}

// SetConfig overrides redis.Server::SetConfig() to sets the specified plug-in configuration.
func (service *Service) SetConfig(conf config.Config) {
	service.BaseService.SetConfig(conf)
}

// ServiceName returns the plug-in service name.
func (service *Service) ServiceName() string {
	return "redis"
}

// Start starts the service.
func (service *Service) Start() error {
	// Set configurations

	port, err := service.GetServiceConfigPort(service)
	if err == nil {
		service.SetPort(port)
	}

	tlsPort, err := service.GetServiceConfigTLSPort()
	if err == nil && (0 < tlsPort) {
		service.SetTLSPort(tlsPort)
		tlsConfig, ok := service.TLSConfig()
		if ok {
			service.Server.SetTLSConfig(tlsConfig)
		}
	}

	passwd, err := service.GetServiceConfigRequirepass()
	if err == nil {
		service.SetRequirePass(passwd)
	}

	if err := service.Server.Start(); err != nil {
		return err
	}

	return nil
}

// Stop stops the service.
func (service *Service) Stop() error {
	if err := service.Server.Stop(); err != nil {
		return err
	}
	return nil
}
