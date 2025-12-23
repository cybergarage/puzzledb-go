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
	"bytes"

	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

// InsertDocument puts a document object with the primary key.
func (txn *transaction) InsertObject(ctx context.Context, docKey store.Key, obj store.Object) error {
	ctx.StartSpan("InsertDocument")
	defer ctx.FinishSpan()

	var encObj bytes.Buffer
	err := txn.EncodeDocument(&encObj, obj)
	if err != nil {
		return err
	}
	kvDocKey := kv.NewKeyWith(kv.DocumentKeyHeader, docKey)
	kvObj := kv.NewObject(kvDocKey, encObj.Bytes())
	return txn.kv.Set(kvObj)
}

// FindDocuments returns a result set matching the specified key.
func (txn *transaction) FindObjects(ctx context.Context, docKey store.Key, opts ...store.Option) (store.ResultSet, error) {
	ctx.StartSpan("FindDocuments")
	defer ctx.FinishSpan()

	kvIdxKey := kv.NewKeyWith(kv.DocumentKeyHeader, docKey)
	kvOpts := NewKvOptionsWith(opts...)
	kvRs, err := txn.kv.Scan(kvIdxKey, kvOpts...)
	if err != nil {
		return nil, err
	}
	return newResultSet(txn.Coder, kvRs), nil
}

// UpdateDocument updates a document object with the specified primary key.
func (txn *transaction) UpdateObject(ctx context.Context, docKey store.Key, obj store.Object) error {
	ctx.StartSpan("UpdateDocument")
	defer ctx.FinishSpan()

	var encObj bytes.Buffer
	err := txn.EncodeDocument(&encObj, obj)
	if err != nil {
		return err
	}
	kvDocKey := kv.NewKeyWith(kv.DocumentKeyHeader, docKey)
	kvObj := kv.NewObject(kvDocKey, encObj.Bytes())
	return txn.kv.Set(kvObj)
}

// RemoveDocument removes a document object with the specified primary key.
func (txn *transaction) RemoveObject(ctx context.Context, docKey store.Key) error {
	ctx.StartSpan("RemoveDocument")
	defer ctx.FinishSpan()

	kvDocKey := kv.NewKeyWith(kv.DocumentKeyHeader, docKey)
	return wrapKeyNotExistError(docKey, txn.kv.Remove(kvDocKey))
}

// RemoveDocument removes document objects with the specified primary key.
func (txn *transaction) RemoveObjects(ctx context.Context, docKey store.Key) error {
	ctx.StartSpan("RemoveDocuments")
	defer ctx.FinishSpan()

	kvDocKey := kv.NewKeyWith(kv.DocumentKeyHeader, docKey)
	return wrapKeyNotExistError(docKey, txn.kv.RemoveRange(kvDocKey))
}

// TruncateDocuments removes all document objects.
func (txn *transaction) TruncateObjects(ctx context.Context) error {
	ctx.StartSpan("TruncateDocuments")
	defer ctx.FinishSpan()

	return txn.RemoveObjects(ctx, document.NewKeyWith(txn.Database().Name()))
}
