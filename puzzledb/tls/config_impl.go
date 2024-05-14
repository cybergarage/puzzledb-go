// Copyright (C) 2022 The PuzzleDB Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package tls

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"github.com/cybergarage/puzzledb-go/puzzledb/config"
)

// tlsConfig represents a TLS configuration.
type tlsConfig struct {
	clientAuthType tls.ClientAuthType
	enabled        bool     `mapstructure:"enabled"`
	certFile       string   `mapstructure:"cert_file"`
	keyFile        string   `mapstructure:"key_file"`
	caFile         []string `mapstructure:"ca_files"`
	tlsConfig      *tls.Config
}

// NewTLSConfig returns a new TLS configuration.
func NewTLSConfig() Config {
	return &tlsConfig{
		enabled:        false,
		clientAuthType: tls.RequireAndVerifyClientCert,
		certFile:       "",
		keyFile:        "",
		caFile:         []string{},
		tlsConfig:      nil,
	}
}

// NewConfigWith returns a new configuration for authenticator with the specified configuration.
func NewConfigWith(config config.Config, path ...string) (Config, error) {
	var tlsConfig tlsConfig
	err := config.UnmarshallConfig(path, &tlsConfig)
	return &tlsConfig, err
}

// SetTLSEnabled sets a TLS enabled flag.
func (config *tlsConfig) SetTLSEnabled(enabled bool) {
	config.enabled = enabled
	config.tlsConfig = nil
}

// SetClientAuthType sets a client authentication type.
func (config *tlsConfig) SetClientAuthType(authType tls.ClientAuthType) {
	config.clientAuthType = authType
}

// SetTLSKeyFile sets a TLS key file.
func (config *tlsConfig) SetTLSKeyFile(file string) {
	config.keyFile = file
	config.tlsConfig = nil
}

// SetTLSCertFile sets a TLS certificate file.
func (config *tlsConfig) SetTLSCertFile(file string) {
	config.certFile = file
	config.tlsConfig = nil
}

// SetRootCertFile sets a TLS root certificates.
func (config *tlsConfig) SetTLSCAFiles(files ...string) {
	config.caFile = files
	config.tlsConfig = nil
}

// TLSEnabled returns a TLS enabled flag.
func (config *tlsConfig) TLSEnabled() bool {
	return config.enabled
}

// ClientAuthType returns a client authentication type.
func (config *tlsConfig) ClientAuthType() tls.ClientAuthType {
	return config.clientAuthType
}

// TLSKeyFile returns a TLS key file.
func (config *tlsConfig) TLSKeyFile() string {
	return config.keyFile
}

// TLSCertFile returns a TLS certificate file.
func (config *tlsConfig) TLSCertFile() string {
	return config.certFile
}

// TLSCAFiles returns a TLS root certificates.
func (config *tlsConfig) TLSCAFiles() []string {
	return config.caFile
}

// TLSConfig returns a TLS configuration from the configuration.
func (config *tlsConfig) TLSConfig() (*tls.Config, error) {
	if config.tlsConfig != nil {
		return config.tlsConfig, nil
	}
	if len(config.certFile) == 0 || len(config.keyFile) == 0 {
		return nil, nil
	}
	serverCert, err := tls.LoadX509KeyPair(config.certFile, config.keyFile)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	for _, rootCertFile := range config.caFile {
		rootCert, err := os.ReadFile(rootCertFile)
		if err != nil {
			return nil, err
		}
		certPool.AppendCertsFromPEM(rootCert)
	}
	config.tlsConfig = &tls.Config{ // nolint: exhaustruct
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
		RootCAs:      certPool,
		ClientAuth:   config.clientAuthType,
	}
	return config.tlsConfig, nil
}
