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

package puzzledbtest

import (
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb"
)

func TestConfig(t *testing.T) {
	c, err := puzzledb.NewConfigWithPath(".")
	if err != nil {
		t.Error(err)
		return
	}

	conf := puzzledb.NewServerConfigWith(c)
	ports := []struct {
		name     string
		expected int
	}{
		{
			name:     "mysql",
			expected: 3306,
		},
	}
	for _, port := range ports {
		portNum, err := conf.Port(port.name)
		if err != nil {
			t.Error(err)
			t.Log(conf.String())
			return
		}
		if portNum != port.expected {
			t.Errorf("expected port number is %d but got %d", port.expected, portNum)
			return
		}
	}
}
