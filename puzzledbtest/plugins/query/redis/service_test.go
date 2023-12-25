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
	"strings"
	"testing"

	"github.com/cybergarage/go-redis/redistest"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestRedisService(t *testing.T) {

	server := puzzledbtest.NewServer()
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

	t.Run("Hash", func(t *testing.T) {
		// redistest.HashCommandTest(t, client)
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

	t.Run("DEL", func(t *testing.T) {
		records := []struct {
			keys     []string
			expected int64
		}{
			{[]string{"key1_del", "key2_del"}, 2},
			{[]string{"key1_del"}, 0},
			{[]string{"key1_del", "key2_del", "key3_del"}, 1},
			{[]string{"key2_del"}, 0},
		}
		for _, r := range records {
			for _, key := range r.keys {
				err = client.Set(key, key, 0).Err()
				if err != nil {
					t.Error(err)
					return
				}
			}
		}
		for _, r := range records {
			t.Run(strings.Join(r.keys, ","), func(t *testing.T) {
				res, err := client.Del(r.keys...).Result()
				if err != nil {
					t.Error(err)
					return
				}
				if res != r.expected {
					t.Errorf("%d != %d", res, r.expected)
					return
				}
			})
		}
	})

	t.Run("EXISTS", func(t *testing.T) {
		if err := client.Set("key1_exists", "val", 0).Err(); err != nil {
			t.Error(err)
			return
		}
		if err := client.Set("key2_exists", "val", 0).Err(); err != nil {
			t.Error(err)
			return
		}
		records := []struct {
			keys     []string
			expected int64
		}{
			{[]string{"nosuchkey"}, 0},
			{[]string{"key1_exists"}, 1},
			{[]string{"key2_exists"}, 1},
			{[]string{"key1_exists", "key2_exists"}, 2},
			{[]string{"key1_exists", "key2_exists", "nosuchkey"}, 2},
		}
		for _, r := range records {
			t.Run(strings.Join(r.keys, ","), func(t *testing.T) {
				res, err := client.Exists(r.keys...).Result()
				if err != nil {
					t.Error(err)
					return
				}
				if res != r.expected {
					t.Errorf("%v : %d != %d", r.keys, res, r.expected)
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
