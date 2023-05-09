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
	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

// ConfigBase represents a basic configuration interface.
type ConfigBase interface {
	// GetConfig returns a value for the specified name.
	GetConfig(paths ...string) (any, error)
	// GetConfigString returns a string value for the specified name.
	GetConfigString(paths ...string) (string, error)
	// GetConfigInt returns an integer value for the specified name.
	GetConfigInt(paths ...string) (int, error)
	// GetConfigBool returns a boolean value for the specified name.
	GetConfigBool(paths ...string) (bool, error)
	// String returns a string representation of the configuration.
	SetConfig(c config.Config)
	// Object returns a raw configuration object.
	Object() config.Config
}

// ServiceConfig represents a configuration interface for service.
type ServiceConfig interface {
	// GetServiceConfig returns a value for the specified name in the service.
	GetServiceConfig(service Service, paths ...string) (any, error)
	// GetServiceConfigString returns a string value for the specified name in the service.
	GetServiceConfigString(service Service, paths ...string) (string, error)
	// GetServiceConfigInt returns an integer value for the specified name in the service.
	GetServiceConfigInt(service Service, paths ...string) (int, error)
	// GetServiceConfigBool returns a boolean value for the specified name in the service.
	GetServiceConfigBool(service Service, paths ...string) (bool, error)
}

// ServiceExtConfig represents an extension configuration interface for service.
type ServiceExtConfig interface {
	// IsServiceEnabled returns true if the service is enabled.
	IsServiceEnabled(service Service) bool
	// GetServicePort returns a port number for the service.
	GetServicePort(service Service) (int, error)
}

// Config represents a plug-in configuration interface.
type Config interface {
	ConfigBase
	ServiceConfig
	ServiceExtConfig
}
