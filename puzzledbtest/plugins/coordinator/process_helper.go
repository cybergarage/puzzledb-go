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

package coordinator

import (
	"errors"
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/coordinator"
)

func truncateCoordinatorStore(coord coordinator.Coordinator) error {
	txn, err := coord.Transact()
	if err != nil {
		return err
	}
	err = txn.Truncate()
	if err != nil {
		return errors.Join(err, txn.Cancel())
	}
	return txn.Commit()
}

// nolint:goerr113, gocognit, gci, gocyclo, gosec, maintidx
func CoordinatorProcessTest(t *testing.T, coord coordinator.Coordinator) {
	t.Helper()

	if err := truncateCoordinatorStore(coord); err != nil {
		t.Error(err)
		return
	}
}
