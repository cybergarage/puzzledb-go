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

package password

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/auth"
	plugins "github.com/cybergarage/puzzledb-go/puzzledb/plugins/auth"
)

// Service is a password authentication service.
type ServiceImpl struct {
	*plugins.BaseService
}

// NewService creates a new service.
func NewService() Service {
	service := &ServiceImpl{
		BaseService: plugins.NewBaseService(),
	}
	return service
}

// ServiceName returns the service name.
func (service *ServiceImpl) ServiceName() string {
	return "password"
}

// CreatePasswordAuthenticatorWithConfig creates a new password authenticator with a configuration.
func (service *ServiceImpl) CreatePasswordAuthenticatorWithConfig(config auth.Config) (auth.PasswordAuthenticator, error) {
	return NewPasswordAuthenticatorWithConfig(config)
}

// Start starts the service.
func (service *ServiceImpl) Start() error {
	return nil
}

// Stop stops the service.
func (service *ServiceImpl) Stop() error {
	return nil
}
