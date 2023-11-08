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

const (
	ConfigPlugins = "plugins"
	ConfigDefault = "default"
	ConfigEnabled = "enabled"
	ConfigPort    = "port"
)

type configImpl struct {
	config.Config
}

func NewConfig() Config {
	return NewConfigWith(nil)
}

func NewConfigWith(conf config.Config) Config {
	return &configImpl{
		Config: conf,
	}
}

// SetConfig sets a manager configuration.
func (conf *configImpl) SetConfig(c config.Config) {
	conf.Config = c
}

// Object returns a raw configuration object.
func (conf *configImpl) Object() config.Config {
	return conf.Config
}

func newServiceConfigPath(service Service, paths ...string) []string {
	servicePaths := []string{ConfigPlugins, service.ServiceType().String(), service.ServiceName()}
	servicePaths = append(servicePaths, paths...)
	return servicePaths
}

// GetServiceConfig returns a value for the specified name in the service.
func (conf *configImpl) GetServiceConfig(service Service, paths ...string) (any, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return nil, NewErrCounfigNotFound(path)
	}
	return conf.GetConfig(path...)
}

// GetServiceConfigString returns a string value for the specified name in the service.
func (conf *configImpl) GetServiceConfigString(service Service, paths ...string) (string, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return "", NewErrCounfigNotFound(path)
	}
	return conf.GetConfigString(path...)
}

// GetServiceConfigInt returns an integer value for the specified name in the service.
func (conf *configImpl) GetServiceConfigInt(service Service, paths ...string) (int, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return 0, NewErrCounfigNotFound(path)
	}
	return conf.GetConfigInt(path...)
}

// GetServiceConfigBool returns a boolean value for the specified name in the service.
func (conf *configImpl) GetServiceConfigBool(service Service, paths ...string) (bool, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return false, NewErrCounfigNotFound(path)
	}
	return conf.GetConfigBool(path...)
}

// IsServiceConfigEnabled returns true if the service is enabled.
func (conf *configImpl) IsServiceConfigEnabled(service Service) bool {
	enabled, err := conf.GetServiceConfigBool(service, ConfigEnabled)
	if err != nil {
		return true
	}
	return enabled
}

// GetServiceConfigPort returns a port number for the service.
func (conf *configImpl) GetServiceConfigPort(service Service) (int, error) {
	port, err := conf.GetServiceConfigInt(service, ConfigPort)
	if err != nil {
		return 0, err
	}
	return port, nil
}
