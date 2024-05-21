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

package puzzledbtest

import (
	"crypto/tls"
	_ "embed"

	puzzledb_tls "github.com/cybergarage/puzzledb-go/puzzledb/tls"
)

//go:embed certs/key.pem
var severKey []byte

//go:embed certs/cert.pem
var serverCert []byte

//go:embed certs/ca.pem
var caCert []byte

func NewTLSConfig() (*tls.Config, error) {
	cfg := puzzledb_tls.NewConfig()
	cfg.SetTLSKey(severKey)
	cfg.SetTLSCert(serverCert)
	cfg.SetTLSCACerts(caCert)
	return cfg.TLSConfig()
}
