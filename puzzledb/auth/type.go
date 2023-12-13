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

import "fmt"

// AuthenticatorType represents an authenticator type.
type AuthenticatorType = int

const (
	// AuthenticatorTypeNone represents an authenticator type of none.
	AuthenticatorTypeNone AuthenticatorType = iota
	// AuthenticatorTypePassword represents an authenticator type of password.
	AuthenticatorTypePassword
)

const (
	// AuthenticatorTypePasswordString represents an authenticator type of password as a string.
	AuthenticatorTypePasswordString = "password"
)

// AuthenticatorTypeFromString returns an authenticator type from the specified string.
func AuthenticatorTypeFromString(str string) (AuthenticatorType, error) {
	switch str {
	case AuthenticatorTypePasswordString:
		return AuthenticatorTypePassword, nil
	}
	return AuthenticatorTypeNone, fmt.Errorf("unknown authenticator type: %s", str)
}
