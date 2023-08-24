// Copyright (C) 2020 The PuzzleDB Authors.
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
	"testing"

	"github.com/cybergarage/go-sqltest/sqltest"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestPostgreSQLTestSuite(t *testing.T) {
	server := puzzledbtest.NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	client := sqltest.NewPostgresClient()

	testNames := []string{
		"YCSBWorkload",
		"SimpInsertSelect",
	}

	if err := sqltest.RunEmbedSuites(t, client, testNames...); err != nil {
		t.Logf("\n%s", server.Store().String())
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
		return
	}
}
