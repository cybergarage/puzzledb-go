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

// Service represents a document store service interface.
type Service interface {
	plugins.Service
	// VerifyCredential verifies the client credential.
	VerifyCredential(conn auth.Conn, q auth.Query) (bool, error)
	// VerifyCertificate verifies the client certificate.
	VerifyCertificate(conn tls.Conn) (bool, error)
}
