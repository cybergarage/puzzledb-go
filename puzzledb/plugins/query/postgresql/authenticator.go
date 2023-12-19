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
	"github.com/cybergarage/go-postgresql/postgresql/protocol/message"
)

// Authenticate authenticates the connection with the startup message.
func (service *Service) Authenticate(conn *postgresql.Conn, startupMessage *message.Startup) (message.Response, error) {
	auths := service.Authenticators()
	if len(auths) == 0 {
		return message.NewAuthenticationOk()
	}
	for _, auth := range auths {
		if passwordAuth, ok := auth.(postgresql.PasswordAuthenticator); ok {
			return passwordAuth.Authenticate()
		}
	}
	return message.NewAuthenticationOk()
}

func (service *Service) authenticateCleartextPassword(conn *postgresql.Conn, startupMessage *message.Startup) (bool, error) {
	/*clientUsername*/ _, ok := startupMessage.User()
	if !ok {
		return false, nil
	}
	/*
		if clientUsername != authenticator.username {
			return false, nil
		}
		authMsg, err := message.NewAuthenticationCleartextPassword()
		if err != nil {
			return false, err
		}
		err = conn.ResponseMessage(authMsg)
		if err != nil {
			return false, err
		}
		msg, err := message.NewPasswordWithReader(conn.MessageReader())
		if err != nil {
			return false, err
		}
		if msg.Password != authenticator.password {
			return false, nil
		}
	*/
	return true, nil
}
