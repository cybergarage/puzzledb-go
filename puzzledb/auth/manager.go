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

// AuthManager represent an authenticator manager interface.
type AuthManager interface {
	// AddAuthenticator adds a new authenticator.
	AddAuthenticator(authenticator Authenticator)
	// ClearAuthenticators clears all authenticators.
	ClearAuthenticators()
	// AuthenticatePassword authenticates a user with a password.
	AuthenticatePassword(conn Conn, username string, password string) (bool, error)
}
