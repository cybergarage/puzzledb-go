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

package ot

import (
	"github.com/cybergarage/go-tracing/tracer/otel"
	tracer_plugin "github.com/cybergarage/puzzledb-go/puzzledb/plugins/tracer"
)

// Service represents a tracer service.
type Service struct {
	*tracer_plugin.BaseService
}

// NewService returns a new tracer service.
func NewService() *Service {
	return &Service{
		BaseService: tracer_plugin.NewBaseServiceWith(otel.NewTracer()),
	}
}

// ServiceName returns the plug-in service name.
func (service *Service) ServiceName() string {
	return "opentelemetry"
}

// GetServiceEndpoint returns the service endpoint.
func (service *Service) GetServiceEndpoint() (string, error) {
	e, err := service.GetServiceConfigString(service, tracer_plugin.EndpointConfig)
	if err != nil {
		return "", err
	}
	return e, nil
}

// Start starts the service.
func (service *Service) Start() error {
	endpoint, err := service.GetServiceEndpoint()
	if err == nil {
		service.SetEndpoint(endpoint)
	}
	if err := service.Tracer.Start(); err != nil {
		return err
	}
	return nil
}

// Stop stops the service.
func (service *Service) Stop() error {
	if err := service.Tracer.Stop(); err != nil {
		return err
	}
	return nil
}
