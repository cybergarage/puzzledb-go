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
	Enabled        bool     `mapstructure:"enabled"`
	CertFile       string   `mapstructure:"cert_file"`
	KeyFile        string   `mapstructure:"key_file"`
	CAFiles        []string `mapstructure:"ca_files"`
	tlsConfig      *tls.Config
}

// NewConfig returns a new TLS configuration.
func NewConfig() Config {
	return &tlsConfig{
		Enabled:        false,
		clientAuthType: tls.RequireAndVerifyClientCert,
		CertFile:       "",
		KeyFile:        "",
		CAFiles:        []string{},
		tlsConfig:      nil,
	}
}

// NewConfigWith returns a new configuration for authenticator with the specified configuration.
func NewConfigWith(config config.Config, path ...string) (Config, error) {
	var tlsConfig tlsConfig
	tlsConfig.clientAuthType = tls.RequireAndVerifyClientCert
	err := config.UnmarshallConfig(path, &tlsConfig)
	return &tlsConfig, err

}

// SetTLSEnabled sets a TLS enabled flag.
func (config *tlsConfig) SetTLSEnabled(enabled bool) {
	config.Enabled = enabled
	config.tlsConfig = nil
}

// SetClientAuthType sets a client authentication type.
func (config *tlsConfig) SetClientAuthType(authType tls.ClientAuthType) {
	config.clientAuthType = authType
}

// SetTLSKeyFile sets a TLS key file.
func (config *tlsConfig) SetTLSKeyFile(file string) {
	config.KeyFile = file
	config.tlsConfig = nil
}

// SetTLSCertFile sets a TLS certificate file.
func (config *tlsConfig) SetTLSCertFile(file string) {
	config.CertFile = file
	config.tlsConfig = nil
}

// SetRootCertFile sets a TLS root certificates.
func (config *tlsConfig) SetTLSCAFiles(files ...string) {
	config.CAFiles = files
	config.tlsConfig = nil
}

// TLSEnabled returns a TLS enabled flag.
func (config *tlsConfig) TLSEnabled() bool {
	return config.Enabled
}

// ClientAuthType returns a client authentication type.
func (config *tlsConfig) ClientAuthType() tls.ClientAuthType {
	return config.clientAuthType
}

// TLSKeyFile returns a TLS key file.
func (config *tlsConfig) TLSKeyFile() string {
	return config.KeyFile
}

// TLSCertFile returns a TLS certificate file.
func (config *tlsConfig) TLSCertFile() string {
	return config.CertFile
}

// TLSCAFiles returns a TLS root certificates.
func (config *tlsConfig) TLSCAFiles() []string {
	return config.CAFiles
}

// TLSConfig returns a TLS configuration from the configuration.
func (config *tlsConfig) TLSConfig() (*tls.Config, error) {
	if config.tlsConfig != nil {
		return config.tlsConfig, nil
	}
	if len(config.CertFile) == 0 || len(config.KeyFile) == 0 {
		return nil, nil
	}
	serverCert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	for _, rootCertFile := range config.CAFiles {
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
