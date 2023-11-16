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

// AuthHandler is an interface for authenticating users.
type AuthHandler interface {
}

// PasswordAuthHandler is an interface for authenticating users with a username.
type PasswordAuthHandler interface {
	// Authenticate authenticates the user with the given credentials.
	Authenticate(conn Conn, username string, password string) (bool, error)
}
