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

package postgresql

import (
	"github.com/cybergarage/go-postgresql/postgresql"
	pgauth "github.com/cybergarage/go-postgresql/postgresql/auth"
	"github.com/cybergarage/go-postgresql/postgresql/protocol/message"
	"github.com/cybergarage/puzzledb-go/puzzledb/auth"
)

// Authenticate authenticates the connection with the startup message.
func (service *Service) Authenticate(conn *postgresql.Conn, startupMessage *message.Startup) (message.Response, error) {
	authcators := service.Authenticators()
	if len(authcators) == 0 {
		return message.NewAuthenticationOk()
	}
	var lastErr error
	for _, authcator := range authcators {
		if authcator, ok := authcator.(auth.PasswordAuthenticator); ok {
			ok, err := service.authenticateCleartextPassword(conn, startupMessage, authcator)
			if err != nil {
				lastErr = err
				continue
			}
			if ok {
				return message.NewAuthenticationOk()
			}
		}
	}
	return nil, lastErr
}

func (service *Service) authenticateCleartextPassword(conn *postgresql.Conn, startupMessage *message.Startup, authcator auth.PasswordAuthenticator) (bool, error) {
	pgAuthenticator := pgauth.NewCleartextPasswordAuthenticatorWith(
		authcator.User(),
		authcator.Password())
	ok, err := pgAuthenticator.Authenticate(conn, startupMessage)
	if err != nil {
		return false, err
	}
	return ok, nil
}
