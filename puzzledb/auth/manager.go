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
	"github.com/cybergarage/go-authenticator/auth"
	"github.com/cybergarage/go-authenticator/auth/tls"
)

// CredentialAuthenticator represent a credential authenticator.
type CredentialAuthenticator = auth.CredentialAuthenticator

// CredentialStore represent a credential store.
type CredentialStore = auth.CredentialStore

// CertificateAuthenticator represent a certificate authenticator.
type CertificateAuthenticator = auth.CertificateAuthenticator

// AuthManager represent an authenticator manager.
type AuthManager interface {
	// SetCommonNameRegexps sets common name regular expressions.
	SetCommonNameRegexps(regexps ...string) error
	// SetCredential sets a credential.
	SetCredentials(creds ...auth.Credential) error
	// LookupCredential looks up a credential.
	LookupCredential(q auth.Query) (Credential, bool, error)
	// VerifyCredential verifies the client credential.
	VerifyCredential(conn Conn, q Query) (bool, error)
	// VerifyCertificate verifies the client certificate.
	VerifyCertificate(conn tls.Conn) (bool, error)
}
