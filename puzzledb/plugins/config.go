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
	"github.com/cybergarage/puzzledb-go/puzzledb/tls"
)

const (
	ConfigTLS     = "tls"
	ConfigPlugins = "plugins"
	ConfigDefault = "default"
	ConfigEnabled = "enabled"
	ConfigPort    = "port"
)

// ConfigBase represents a basic configuration interface.
type ConfigBase interface {
	config.Config
	// String returns a string representation of the configuration.
	SetConfig(c config.Config)
	// Object returns a raw configuration object.
	Object() config.Config
}

// RootConfig represents a root configuration interface.
type RootConfig interface {
	// LookupTLSConfig returns a TLS configuration.
	LookupTLSConfig() (tls.Config, error)
}

// ServiceTypeConfig represents a configuration interface for service type.
type ServiceTypeConfig interface {
	// LookupServiceTypeConfig returns a value for the specified name in the service type.
	LookupServiceTypeConfig(serviceType ServiceType, item string) (any, error)
	// LookupServiceTypeConfigString returns a string value for the specified name in the service type.
	LookupServiceTypeConfigString(serviceType ServiceType, item string) (string, error)
	// LookupServiceTypeConfigInt returns an integer value for the specified name in the service type.
	LookupServiceTypeConfigInt(serviceType ServiceType, item string) (int, error)
	// LookupServiceTypeConfigBool returns a boolean value for the specified name in the service type.
	LookupServiceTypeConfigBool(serviceType ServiceType, item string) (bool, error)
}

// ServiceTypeExtConfig represents an extension configuration interface for service type.
type ServiceTypeExtConfig interface {
	// IsServiceTypeConfigEnabled returns true if the service type is enabled.
	IsServiceTypeConfigEnabled(serviceType ServiceType) bool
	// LookupServiceTypeConfigPort returns a port number for the service type.
	LookupServiceTypeDefault(serviceType ServiceType) (string, error)
}

// ServiceConfig represents a configuration interface for service.
type ServiceConfig interface {
	// LookupServiceConfig returns a value for the specified name in the service.
	LookupServiceConfig(service Service, paths ...string) (any, error)
	// LookupServiceConfigString returns a string value for the specified name in the service.
	LookupServiceConfigString(service Service, paths ...string) (string, error)
	// LookupServiceConfigInt returns an integer value for the specified name in the service.
	LookupServiceConfigInt(service Service, paths ...string) (int, error)
	// LookupServiceConfigBool returns a boolean value for the specified name in the service.
	LookupServiceConfigBool(service Service, paths ...string) (bool, error)
}

// ServiceExtConfig represents an extension configuration interface for service.
type ServiceExtConfig interface {
	// IsServiceConfigEnabled returns true if the service is enabled.
	IsServiceConfigEnabled(service Service) bool
	// LookupServiceConfigPort returns a port number for the service.
	LookupServiceConfigPort(service Service) (int, error)
}

// Config represents a plug-in configuration interface.
type Config interface {
	ConfigBase
	RootConfig
	ServiceTypeConfig
	ServiceTypeExtConfig
	ServiceConfig
	ServiceExtConfig
}
