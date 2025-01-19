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

package auth

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/auth"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
)

type service struct {
	plugins.Config
	auth.AuthManager
}

// NewService returns a new query base service.
func NewService() Service {
	service := &service{
		AuthManager: auth.NewAuthManager(),
		Config:      plugins.NewConfig(),
	}
	return service
}

// ServiceType returns the plug-in service type.
func (service *service) ServiceType() plugins.ServiceType {
	return plugins.AuthService
}

// ServiceName returns the plug-in service name.
func (service *service) ServiceName() string {
	return "auth"
}

// Start starts the service.
func (service *service) Start() error {
	ok := service.IsServiceTypeConfigEnabled(plugins.AuthService)
	if !ok {
		return nil
	}

	plainConfigs, err := auth.NewPlainConfigFrom(
		service,
		plugins.ConfigPlugins,
		service.ServiceType().String(),
		auth.AuthenticatorTypePlainString,
	)
	if err != nil {
		return err
	}

	creds := []auth.Credential{}
	for _, plainConfig := range plainConfigs {
		cred := auth.NewCredential(
			auth.WithCredentialUsername(plainConfig.Username),
			auth.WithCredentialPassword(plainConfig.Password),
		)
		creds = append(creds, cred)
	}
	service.SetCredentials(creds...)

	return nil
}

// Stop stops the service.
func (service *service) Stop() error {
	return nil
}
