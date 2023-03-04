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
	"testing"

	"github.com/cybergarage/go-redis/redistest"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestService(t *testing.T) {

	server := test.NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	client := redistest.NewClient()
	err = client.Open(server.Host)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("Connection", func(t *testing.T) {
		redistest.ConnectionCommandTest(t, client)
	})

	t.Run("SET", func(t *testing.T) {
		records := []struct {
			key      string
			val      string
			expected string
		}{
			{"key_set", "value0", "value0"},
			{"key_set", "value1", "value1"},
			{"key_set", "value2", "value2"},
		}

		for _, r := range records {
			t.Run(r.key+":"+r.val, func(t *testing.T) {
				err := client.Set(r.key, r.val, 0).Err()
				if err != nil {
					t.Error(err)
					return
				}
				res, err := client.Get(r.key).Result()
				if err != nil {
					t.Error(err)
					return
				}
				if res != r.expected {
					t.Errorf("%s != %s", res, r.expected)
					return
				}
			})
		}
	})

	err = client.Close()
	if err != nil {
		t.Error(err)
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
	}
}
