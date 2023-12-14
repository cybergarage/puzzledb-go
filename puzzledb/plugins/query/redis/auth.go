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

package redis

import (
	"errors"

	"github.com/cybergarage/go-redis/redis"
)

// Auth authenticates a client.
func (service *Service) Auth(conn *Conn, username string, password string) (*Message, error) {
	ok, err := service.AuthenticatePassword(conn, username, password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid username or password")
	}
	return redis.NewOKMessage(), nil
}
