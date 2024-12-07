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
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// InsertIndex puts a secondary index with the primary key.
func (txn *transaction) InsertIndex(ctx context.Context, idxKey store.Key, docKey store.Key) error {
	ctx.StartSpan("InsertIndex")
	defer ctx.FinishSpan()

	kvDocKey := kv.NewKeyWith(kv.DocumentKeyHeader, docKey)
	kvDocKeyBytes, err := txn.EncodeKey(kvDocKey)
	if err != nil {
		return err
	}
	kvIdxKey := kv.NewKeyWith(kv.SecondaryIndexHeader, idxKey)
	kvObj := kv.NewObject(kvIdxKey, kvDocKeyBytes)
	return txn.kv.Set(kvObj)
}

// RemoveIndex removes the specified secondary index.
func (txn *transaction) RemoveIndex(ctx context.Context, idxKey store.Key) error {
	ctx.StartSpan("RemoveIndex")
	defer ctx.FinishSpan()

	kvIdxKey := kv.NewKeyWith(kv.SecondaryIndexHeader, idxKey)
	return wrapKeyNotExistError(idxKey, txn.kv.Remove(kvIdxKey))
}

// FindDocumentsByIndex gets document objects matching the specified index key.
func (txn *transaction) FindObjectsByIndex(ctx context.Context, idxKey store.Key, opts ...store.Option) (store.ResultSet, error) {
	ctx.StartSpan("FindDocumentsByIndex")
	defer ctx.FinishSpan()

	kvIdxKey := kv.NewKeyWith(kv.SecondaryIndexHeader, idxKey)
	kvOpts := NewKvOptionsWith(opts...)
	kvIdxRs, err := txn.kv.GetRange(kvIdxKey, kvOpts...)
	if err != nil {
		return nil, err
	}
	return newIndexResultSet(txn, txn.KeyCoder, txn.Coder, kvIdxRs), nil
}

// TruncateIndexes removes all secondary indexes.
func (txn *transaction) TruncateIndexes(ctx context.Context) error {
	ctx.StartSpan("TruncateIndexes")
	defer ctx.FinishSpan()

	kvSchemaKey := kv.NewKeyWith(kv.SecondaryIndexHeader, document.NewKeyWith(txn.Database().Name()))
	return txn.kv.RemoveRange(kvSchemaKey)
}
