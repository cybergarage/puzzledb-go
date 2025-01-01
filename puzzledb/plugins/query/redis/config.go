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

const (
	tlsPort     = "tls_port"
	requirepass = "requirepass"
)

// LookupServiceConfigRequirepass returns the requirepass value of the service.
func (service *Service) LookupServiceConfigRequirepass() (string, error) {
	passwd, err := service.LookupServiceConfigString(service, requirepass)
	if err != nil {
		return "", err
	}
	return passwd, nil
}

// LookupServiceConfigTLSPort returns the TLS port value of the service.
func (service *Service) LookupServiceConfigTLSPort() (int, error) {
	port, err := service.LookupServiceConfigInt(service, tlsPort)
	if err != nil {
		return 0, err
	}
	return port, nil
}
