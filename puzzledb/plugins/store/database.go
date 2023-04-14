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
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type database struct {
	kv kv.Database
	document.Coder
	document.KeyCoder
}

// Name returns the unique name.
func (db *database) Name() string {
	return db.kv.Name()
}

// Transact begin a new transaction.
func (db *database) Transact(write bool) (store.Transaction, error) {
	kvTx, err := db.kv.Transact(write)
	if err != nil {
		return nil, err
	}
	return newTransaction(db, kvTx, db.Coder, db.KeyCoder)
}
