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
	"fmt"
	"testing"
	"time"
)

const (
	testDBPrefix = "grpc"
)

func TestClient(t *testing.T) {
	server := NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		err = server.Stop()
		if err != nil {
			t.Error(err)
			return
		}
	}()

	client := NewClient()
	err = client.Open()
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		err = client.Close()
		if err != nil {
			t.Error(err)
			return
		}
	}()

	// Checks config services

	_, err = client.ListConfig()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = client.GetConfig("logging.level")
	if err != nil {
		t.Error(err)
		return
	}

	// Checks store services

	testDBName := fmt.Sprintf("%s%d", testDBPrefix, time.Now().UnixNano())

	err = client.CreateDatabase(testDBName)
	if err != nil {
		t.Error(err)
		return
	}

	dbs, err := client.ListDatabases()
	if err != nil {
		t.Error(err)
		return
	}

	if len(dbs) < 1 || dbs[0] != testDBName {
		t.Errorf("Unexpected database list : %v", dbs)
		return
	}

	err = client.RemoveDatabase(testDBName)
	if err != nil {
		t.Error(err)
		return
	}

	dbs, err = client.ListDatabases()
	if err != nil {
		t.Error(err)
		return
	}

	if len(dbs) != 0 {
		t.Errorf("Unexpected database list : %v", dbs)
		return
	}
}
