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

package sqltest

import (
	"testing"

	"github.com/cybergarage/go-sqltest/sqltest"
	"github.com/cybergarage/puzzledb-go/puzzledbtest"
)

func TestMySQLTestSuite(t *testing.T) {
	server := puzzledbtest.NewServer()
	err := server.Start()
	if err != nil {
		t.Error(err)
		return
	}

	client := sqltest.NewMySQLClient()
	client.SetPreparedStatementEnabled(false)

	testRegexes := []string{
		"SmplTxn.*",
		"SmplCrud.*",
		"SmplIndex*",
		"YcsbWorkload",
	}

	var databaseDump string
	dumpDatabase := func(*sqltest.Suite, *sqltest.ScenarioTester, error) {
		databaseDump = server.Store().String()
	}

	stepHander := func(scenario *sqltest.Scenario, n int, query string, err error) {
		// t.Logf("[%d]: %s", n, query)
		// t.Logf("\n%s", server.Store().String())
	}

	suite, err := sqltest.NewSuiteWith(
		sqltest.WithSuiteEmbeds(),
		sqltest.WithSuiteRegexes(testRegexes...),
		sqltest.WithSuiteClient(client),
		sqltest.WithSuiteErrorHandler(dumpDatabase),
		sqltest.WithSuiteStepHandler(stepHander),
	)
	if err != nil {
		t.Error(err)
	}

	err = suite.Test(t)
	if err != nil {
		t.Logf("\n%s", databaseDump)
	}

	err = server.Stop()
	if err != nil {
		t.Error(err)
		return
	}
}
