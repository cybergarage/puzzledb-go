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
	"fmt"
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
	Key            []byte
	Cert           []byte
	CASCerts       [][]byte
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
		Key:            []byte{},
		Cert:           []byte{},
		CASCerts:       [][]byte{},
		tlsConfig:      nil,
	}
}

// NewConfigWith returns a new configuration for authenticator with the specified configuration.
func NewConfigWith(config config.Config, path ...string) (Config, error) {
	var tlsConfig tlsConfig
	err := config.UnmarshallConfig(path, &tlsConfig)
	if err != nil {
		return nil, err
	}
	tlsConfig.clientAuthType = tls.RequireAndVerifyClientCert
	if !tlsConfig.Enabled {
		return &tlsConfig, nil
	}
	if 0 < len(tlsConfig.CertFile) {
		if err := tlsConfig.SetTLSCertFile(tlsConfig.CertFile); err != nil {
			return nil, err
		}
	}
	if 0 < len(tlsConfig.KeyFile) {
		if err := tlsConfig.SetTLSKeyFile(tlsConfig.KeyFile); err != nil {
			return nil, err
		}
	}
	if 0 < len(tlsConfig.CAFiles) {
		if err := tlsConfig.SetTLSCACertFiles(tlsConfig.CAFiles...); err != nil {
			return nil, err
		}
	}
	return &tlsConfig, nil
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
func (config *tlsConfig) SetTLSKeyFile(file string) error {
	key, err := os.ReadFile(file)
	if err != nil {
		pwd, _ := os.Getwd()
		return fmt.Errorf("%w (%s)", err, pwd)
	}
	config.KeyFile = file
	config.SetTLSKey(key)
	return nil
}

// SetTLSCertFile sets a TLS certificate file.
func (config *tlsConfig) SetTLSCertFile(file string) error {
	cert, err := os.ReadFile(file)
	if err != nil {
		pwd, _ := os.Getwd()
		return fmt.Errorf("%w (%s)", err, pwd)
	}
	config.CertFile = file
	config.SetTLSCert(cert)
	return nil
}

// SetRootCertFile sets a TLS root certificates.
func (config *tlsConfig) SetTLSCACertFiles(files ...string) error {
	certs := make([][]byte, len(files))
	for n, file := range files {
		cert, err := os.ReadFile(file)
		if err != nil {
			pwd, _ := os.Getwd()
			return fmt.Errorf("%w (%s)", err, pwd)
		}
		certs[n] = cert
	}
	config.CAFiles = files
	config.SetTLSCACerts(certs...)
	return nil
}

// SetTLSKey sets a TLS key file.
func (config *tlsConfig) SetTLSKey(b []byte) {
	config.Key = b
	config.tlsConfig = nil
}

// SetTLSCert sets a TLS certificate binaries.
func (config *tlsConfig) SetTLSCert(b []byte) {
	config.Cert = b
	config.tlsConfig = nil
}

// SetTLSCACerts sets a TLS root certificate binaries.
func (config *tlsConfig) SetTLSCACerts(b ...[]byte) {
	config.CASCerts = b
	config.tlsConfig = nil
}

// TLSEnabled returns a TLS enabled flag.
func (config *tlsConfig) TLSEnabled() bool {
	return config.Enabled
}

// TLSKey returns a TLS key file.
func (config *tlsConfig) TLSKey() []byte {
	return config.Key
}

// TLSCert returns a TLS certificate file.
func (config *tlsConfig) TLSCert() []byte {
	return config.Cert
}

// TLSCACerts returns a TLS root certificate bytes.
func (config *tlsConfig) TLSCACerts() [][]byte {
	return config.CASCerts
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
	if len(config.Cert) == 0 || len(config.Key) == 0 {
		return nil, nil
	}
	serverCert, err := tls.X509KeyPair(config.Cert, config.Key)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	for _, caCert := range config.CASCerts {
		certPool.AppendCertsFromPEM(caCert)
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
