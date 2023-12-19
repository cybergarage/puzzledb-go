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

package password

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/auth"
)

// PasswordAuthenticator is a password authenticator.
type PasswordAuthenticator struct {
	user   string
	passwd string
}

// NewPasswordAuthenticatorWithConfig creates a new password authenticator with a configuration.
func NewPasswordAuthenticatorWithConfig(config auth.Config) (auth.PasswordAuthenticator, error) {
	authenticator := &PasswordAuthenticator{
		user:   config.User,
		passwd: config.Password,
	}
	return authenticator, nil
}

// User returns the authorized user.
func (authenticator *PasswordAuthenticator) User() string {
	return authenticator.user
}

// Password returns the authorized password.
func (authenticator *PasswordAuthenticator) Password() string {
	return authenticator.passwd
}

// AuthenticatePassword authenticates a user with a password.
func (authenticator *PasswordAuthenticator) AuthenticatePassword(conn auth.Conn, username string, password string) (bool, error) {
	if authenticator.user != username {
		return false, nil
	}
	if authenticator.passwd != password {
		return false, nil
	}
	return true, nil

}
