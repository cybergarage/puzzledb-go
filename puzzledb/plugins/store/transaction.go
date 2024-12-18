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
	"fmt"
	"time"

	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

type transaction struct {
	kv kv.Transaction
	document.Coder
	document.KeyCoder
	db         *database
	autoCommit bool
}

func newTransaction(db *database, kvTx kv.Transaction, docCoder document.Coder, keyCoder document.KeyCoder) (store.Transaction, error) {
	return &transaction{
		db:         db,
		kv:         kvTx,
		Coder:      docCoder,
		KeyCoder:   keyCoder,
		autoCommit: true,
	}, nil
}

// Database returns the transaction database.
func (txn *transaction) Database() store.Database {
	return txn.db
}

// SetAutoCommit sets the auto commit flag.
func (txn *transaction) SetAutoCommit(v bool) {
	txn.autoCommit = v
}

// IsAutoCommit returns true whether the auto commit flag is set.
func (txn *transaction) IsAutoCommit() bool {
	return txn.autoCommit
}

// Commit commits this transaction.
func (txn *transaction) Commit(ctx context.Context) error {
	ctx.StartSpan("Commit")
	ctx.FinishSpan()
	return txn.kv.Commit()
}

// Cancel cancels this transaction.
func (txn *transaction) Cancel(ctx context.Context) error {
	ctx.StartSpan("Cancel")
	ctx.FinishSpan()
	return txn.kv.Cancel()
}

// CancelWithError cancels this transaction.
func (txn *transaction) CancelWithError(err error) error {
	txnErr := txn.kv.Cancel()
	if txnErr == nil {
		return err
	}
	return fmt.Errorf("%w: %w", err, txnErr)
}

// SetTimeout sets the timeout of this transaction.
func (txn *transaction) SetTimeout(t time.Duration) error {
	return txn.kv.SetTimeout(t)
}
