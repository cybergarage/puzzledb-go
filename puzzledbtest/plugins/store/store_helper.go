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

package store

import (
	"testing"

	plugins "github.com/cybergarage/puzzledb-go/puzzledb/plugins/store"
)

const (
	testDBName    = "testdoc"
	testKeyCount  = 10
	testValBufMax = 8
)

//nolint:gosec,cyclop,gocognit,gocyclo,maintidx
func StoreTest(t *testing.T, service plugins.Service) {
	t.Helper()

	if err := service.Start(); err != nil {
		t.Error(err)
		return
	}
	if err := service.CreateDatabase(testDBName); err != nil {
		t.Error(err)
		return
	}
	_, err := service.GetDatabase(testDBName)
	if err != nil {
		t.Error(err)
		return
	}

	if err := service.Stop(); err != nil {
		t.Error(err)
		return
	}
}
