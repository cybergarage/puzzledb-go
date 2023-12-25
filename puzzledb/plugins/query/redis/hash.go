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
	"github.com/cybergarage/go-redis/redis"
)

func (service *Service) HDel(conn *Conn, key string, fields []string) (*Message, error) {
	return nil, newErrNotSupported("HDel")
}

func (service *Service) HSet(conn *Conn, key string, field string, val string, opt redis.HSetOption) (*Message, error) {
	return nil, newErrNotSupported("HSet")
}

func (service *Service) HGet(conn *Conn, key string, field string) (*Message, error) {
	return nil, newErrNotSupported("HGet")
}

func (service *Service) HGetAll(conn *Conn, key string) (*Message, error) {
	return nil, newErrNotSupported("HGetAll")
}
