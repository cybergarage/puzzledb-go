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
	"github.com/cybergarage/puzzledb-go/puzzledb/auth/tls"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
)

type service struct {
	plugins.Config
}

// NewService returns a new query base service.
func NewService() Service {
	server := &service{
		Config: plugins.NewConfig(),
	}
	return server
}

// ServiceType returns the plug-in service type.
func (service *service) ServiceType() plugins.ServiceType {
	return plugins.AuthenticatorService
}

// ServiceName returns the plug-in service name.
func (service *service) ServiceName() string {
	return "auth"
}

// VerifyCredential verifies the client credential.
func (service *service) VerifyCredential(conn auth.Conn, q auth.Query) (bool, error) {
	return true, nil
}

// VerifyCertificate verifies the client certificate.
func (service *service) VerifyCertificate(conn tls.Conn) (bool, error) {
	return true, nil
}

func (service *service) Start() error {
	return nil
}

func (service *service) Stop() error {
	return nil
}
