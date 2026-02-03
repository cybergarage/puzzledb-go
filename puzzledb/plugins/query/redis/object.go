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

import "github.com/cybergarage/go-redis/redis"

// HashObject represents a hash object.
type HashObject map[string]string

// Set sets the specified field to the hash object and returns 1 if the field was newly created.
func (obj HashObject) Set(field string, val string, opt redis.HSetOption) int {
	_, hasKey := obj[field]
	if opt.NX && hasKey {
		return 0
	}
	obj[field] = val
	if hasKey {
		return 0
	}
	return 1
}

func (obj HashObject) Del(fields []string) int {
	removedFields := 0
	for _, field := range fields {
		_, ok := obj[field]
		if !ok {
			continue
		}
		delete(obj, field)
		removedFields++
	}
	return removedFields
}
