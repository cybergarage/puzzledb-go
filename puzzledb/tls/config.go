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

package tls

import (
	"crypto/tls"
)

// Config represents a TLS configuration.
type Config interface {
	// SetTLSEnabled sets a TLS enabled flag.
	SetTLSEnabled(enabled bool)
	// SetClientAuthType sets a client authentication type.
	SetClientAuthType(authType tls.ClientAuthType)
	// SetTLSKeyFile sets a TLS key file.
	SetTLSKeyFile(file string)
	// SetTLSCertFile sets a TLS certificate file.
	SetTLSCertFile(file string)
	// SetRootCertFile sets a TLS root certificates.
	SetTLSCAFiles(files ...string)
	// TLSEnabled returns a TLS enabled flag.
	TLSEnabled() bool
	// ClientAuthType returns a client authentication type.
	ClientAuthType() tls.ClientAuthType
	// TLSKeyFile returns a TLS key file.
	TLSKeyFile() string
	// TLSCertFile returns a TLS certificate file.
	TLSCertFile() string
	// TLSCAFiles returns a TLS root certificates.
	TLSCAFiles() []string
	// TLSConfig returns a TLS configuration from the configuration.
	TLSConfig() (*tls.Config, error)
}
