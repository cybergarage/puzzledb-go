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

// authManagerImpl represent an authenticator manager.
type authManagerImpl struct {
	authenticators []Authenticator
}

// NewAuthManager returns a new authenticator manager.
func NewAuthManager() AuthManager {
	manager := &authManagerImpl{
		authenticators: make([]Authenticator, 0),
	}
	return manager
}

// Authenticators returns all authenticators.
func (mgr *authManagerImpl) Authenticators() []Authenticator {
	return mgr.authenticators
}

// AddAuthenticator adds a new authenticator.
func (mgr *authManagerImpl) AddAuthenticator(authenticator Authenticator) {
	mgr.authenticators = append(mgr.authenticators, authenticator)
}

// ClearAuthenticators clears all authenticators.
func (mgr *authManagerImpl) ClearAuthenticators() {
	mgr.authenticators = make([]Authenticator, 0)
}

// AuthenticatePassword authenticates a user with a password.
func (mgr *authManagerImpl) AuthenticatePassword(conn Conn, username string, password string) (bool, error) {
	for _, authenticator := range mgr.authenticators {
		if passwordAuthenticator, ok := authenticator.(PasswordAuthenticator); ok {
			ok, err := passwordAuthenticator.AuthenticatePassword(conn, username, password)
			if err != nil {
				return false, err
			}
			if ok {
				return true, nil
			}
		}
	}
	return false, nil
}
