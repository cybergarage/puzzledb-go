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

func newServiceTypeConfigPath(serviceType ServiceType, item string) []string {
	return []string{ConfigPlugins, serviceType.String(), item}
}

// LookupServiceTypeConfig returns a value for the specified name in the service type.
func (conf *configImpl) LookupServiceTypeConfig(serviceType ServiceType, item string) (any, error) {
	path := newServiceTypeConfigPath(serviceType, item)
	if conf.Config == nil {
		return nil, NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigObject(path...)
}

// LookupServiceTypeConfigString returns a string value for the specified name in the service type.
func (conf *configImpl) LookupServiceTypeConfigString(serviceType ServiceType, item string) (string, error) {
	path := newServiceTypeConfigPath(serviceType, item)
	if conf.Config == nil {
		return "", NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigString(path...)
}

// LookupServiceTypeConfigInt returns an integer value for the specified name in the service type.
func (conf *configImpl) LookupServiceTypeConfigInt(serviceType ServiceType, item string) (int, error) {
	path := newServiceTypeConfigPath(serviceType, item)
	if conf.Config == nil {
		return 0, NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigInt(path...)
}

// LookupServiceTypeConfigBool returns a boolean value for the specified name in the service type.
func (conf *configImpl) LookupServiceTypeConfigBool(serviceType ServiceType, item string) (bool, error) {
	path := newServiceTypeConfigPath(serviceType, item)
	if conf.Config == nil {
		return false, NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigBool(path...)
}

// IsServiceTypeConfigEnabled returns true if the service type is enabled.
func (conf *configImpl) IsServiceTypeConfigEnabled(serviceType ServiceType) bool {
	enabled, err := conf.LookupServiceTypeConfigBool(serviceType, ConfigEnabled)
	if err != nil {
		return true
	}
	return enabled
}

// LookupServiceTypeConfigPort returns a port number for the service type.
func (conf *configImpl) LookupServiceTypeDefault(serviceType ServiceType) (string, error) {
	def, err := conf.LookupServiceTypeConfigString(serviceType, ConfigDefault)
	if err != nil {
		return "", err
	}
	return def, nil
}

func newServiceConfigPath(service Service, paths ...string) []string {
	servicePaths := make([]string, 0, 3+len(paths))
	servicePaths = append(servicePaths, ConfigPlugins, service.ServiceType().String(), service.ServiceName())
	servicePaths = append(servicePaths, paths...)
	return servicePaths
}

// LookupServiceConfig returns a value for the specified name in the service.
func (conf *configImpl) LookupServiceConfig(service Service, paths ...string) (any, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return nil, NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigObject(path...)
}

// LookupServiceConfigString returns a string value for the specified name in the service.
func (conf *configImpl) LookupServiceConfigString(service Service, paths ...string) (string, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return "", NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigString(path...)
}

// LookupServiceConfigInt returns an integer value for the specified name in the service.
func (conf *configImpl) LookupServiceConfigInt(service Service, paths ...string) (int, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return 0, NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigInt(path...)
}

// LookupServiceConfigBool returns a boolean value for the specified name in the service.
func (conf *configImpl) LookupServiceConfigBool(service Service, paths ...string) (bool, error) {
	path := newServiceConfigPath(service, paths...)
	if conf.Config == nil {
		return false, NewErrCounfigNotFound(path)
	}
	return conf.LookupConfigBool(path...)
}

// IsServiceConfigEnabled returns true if the service is enabled.
func (conf *configImpl) IsServiceConfigEnabled(service Service) bool {
	enabled, err := conf.LookupServiceConfigBool(service, ConfigEnabled)
	if err != nil {
		return true
	}
	return enabled
}

// LookupServiceConfigPort returns a port number for the service.
func (conf *configImpl) LookupServiceConfigPort(service Service) (int, error) {
	port, err := conf.LookupServiceConfigInt(service, ConfigPort)
	if err != nil {
		return 0, err
	}
	return port, nil
}
