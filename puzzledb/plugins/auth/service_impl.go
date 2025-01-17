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
	"regexp"

	"github.com/cybergarage/puzzledb-go/puzzledb/auth"
	"github.com/cybergarage/puzzledb-go/puzzledb/auth/tls"
	"github.com/cybergarage/puzzledb-go/puzzledb/plugins"
)

type service struct {
	plugins.Config
	auth.AuthManager
	credStore        map[string]auth.Credential
	commonNameRegexp []*regexp.Regexp
}

// NewService returns a new query base service.
func NewService() Service {
	service := &service{
		AuthManager:      auth.NewAuthManager(),
		Config:           plugins.NewConfig(),
		credStore:        map[string]auth.Credential{},
		commonNameRegexp: []*regexp.Regexp{},
	}
	service.AuthManager.SetCredentialStore(service)
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

// SetCommonNameRegexps sets common name regular expressions.
func (service *service) SetCommonNameRegexps(regexps ...string) error {
	for _, re := range regexps {
		r, err := regexp.Compile(re)
		if err != nil {
			return err
		}
		service.commonNameRegexp = append(service.commonNameRegexp, r)
	}
	return nil
}

// SetCredential sets a credential.
func (service *service) SetCredentials(creds ...auth.Credential) error {
	for _, cred := range creds {
		service.credStore[cred.Username()] = cred
	}
	return nil
}

// LookupCredential looks up a credential.
func (service *service) LookupCredential(q auth.Query) (auth.Credential, bool, error) {
	user := q.Username()
	cred, ok := service.credStore[user]
	return cred, ok, nil
}

// VerifyCertificate verifies the client certificate.
func (service *service) VerifyCertificate(conn tls.Conn) (bool, error) {
	if len(service.commonNameRegexp) == 0 {
		return true, nil
	}
	for _, cert := range conn.ConnectionState().PeerCertificates {
		for _, re := range service.commonNameRegexp {
			if re.MatchString(cert.Subject.CommonName) {
				return true, nil
			}
		}
	}
	return false, nil
}

// Start starts the service.
func (service *service) Start() error {
	service.credStore = map[string]auth.Credential{}
	service.commonNameRegexp = []*regexp.Regexp{}

	plainConfigs, err := auth.NewPlainConfigFrom(
		service,
		plugins.ConfigPlugins,
		service.ServiceType().String(),
		auth.AuthenticatorTypePlainString,
	)
	if err != nil {
		return err
	}

	for _, plainConfig := range plainConfigs {
		if !plainConfig.Enabled {
			continue
		}
		cred := auth.NewCredential(
			auth.WithCredentialUsername(plainConfig.Username),
			auth.WithCredentialPassword(plainConfig.Password),
		)
		service.SetCredentials(cred)
	}

	return nil
}

// Stop stops the service.
func (service *service) Stop() error {
	return nil
}
