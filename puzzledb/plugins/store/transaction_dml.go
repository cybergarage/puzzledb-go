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

func (txn *transaction) createSchemaKey(schema string) store.Key {
	colKey := document.NewKeyWith(txn.Database().Name(), schema)
	return kv.NewKeyWith(kv.CollectionKeyHeader, colKey)
}

// CreateCollection creates a new schema.
func (txn *transaction) CreateCollection(ctx context.Context, col store.Collection) error {
	ctx.StartSpan("CreateCollection")
	defer ctx.FinishSpan()

	kvSchemaKey := txn.createSchemaKey(col.Name())
	var encSchema bytes.Buffer
	err := txn.EncodeDocument(&encSchema, col.Data())
	if err != nil {
		return err
	}
	kvObj := kv.Object{
		Key:   kvSchemaKey,
		Value: encSchema.Bytes(),
	}
	return txn.kv.Set(&kvObj)
}

// UpdateCollection updates the specified collection.
func (txn *transaction) UpdateCollection(ctx context.Context, col store.Collection) error {
	ctx.StartSpan("UpdateCollection")
	defer ctx.FinishSpan()

	kvSchemaKey := txn.createSchemaKey(col.Name())
	var encSchema bytes.Buffer
	err := txn.EncodeDocument(&encSchema, col.Data())
	if err != nil {
		return err
	}
	kvObj := kv.Object{
		Key:   kvSchemaKey,
		Value: encSchema.Bytes(),
	}
	return txn.kv.Set(&kvObj)
}

// GetCollection returns the specified schema.
func (txn *transaction) GetCollection(ctx context.Context, name string) (store.Collection, error) {
	ctx.StartSpan("GetCollection")
	defer ctx.FinishSpan()

	kvSchemaKey := txn.createSchemaKey(name)
	kvRs, err := txn.kv.GetRange(kvSchemaKey)
	if err != nil {
		return nil, err
	}
	if !kvRs.Next() {
		return nil, store.NewSchemaNotExistError(name)
	}
	kvObj := kvRs.Object()
	obj, err := txn.DecodeDocument(bytes.NewReader(kvObj.Value))
	if err != nil {
		return nil, err
	}
	return document.NewCollectionWith(obj)
}

// RemoveCollection removes the specified schema.
func (txn *transaction) RemoveCollection(ctx context.Context, name string) error {
	ctx.StartSpan("RemoveCollection")
	defer ctx.FinishSpan()

	kvSchemaKey := txn.createSchemaKey(name)
	return txn.kv.Remove(kvSchemaKey)
}

// TruncateCollections removes all schemas.
func (txn *transaction) TruncateCollections(ctx context.Context) error {
	ctx.StartSpan("TruncateCollections")
	defer ctx.FinishSpan()

	kvSchemaKey := kv.NewKeyWith(kv.CollectionKeyHeader, document.NewKeyWith(txn.Database().Name()))
	return txn.kv.RemoveRange(kvSchemaKey)
}
