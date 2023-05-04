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
	return kv.NewKeyWith(kv.SchemaKeyHeader, colKey)
}

// CreateSchema creates a new schema.
func (txn *transaction) CreateSchema(ctx context.Context, schema store.Schema) error {
	ctx.StartSpan("CreateSchema")
	defer ctx.FinishSpan()

	kvSchemaKey := txn.createSchemaKey(schema.Name())
	var encSchema bytes.Buffer
	err := txn.EncodeDocument(&encSchema, schema.Data())
	if err != nil {
		return err
	}
	kvObj := kv.Object{
		Key:   kvSchemaKey,
		Value: encSchema.Bytes(),
	}
	return txn.kv.Set(&kvObj)
}

// GetSchema returns the specified schema.
func (txn *transaction) GetSchema(ctx context.Context, name string) (store.Schema, error) {
	ctx.StartSpan("GetSchema")
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
	return document.NewSchemaWith(obj)
}

// RemoveSchema removes the specified schema.
func (txn *transaction) RemoveSchema(ctx context.Context, name string) error {
	ctx.StartSpan("RemoveSchema")
	defer ctx.FinishSpan()

	kvSchemaKey := txn.createSchemaKey(name)
	return txn.kv.Remove(kvSchemaKey)
}

// TruncateSchemas removes all schemas.
func (txn *transaction) TruncateSchemas(ctx context.Context) error {
	ctx.StartSpan("TruncateSchemas")
	defer ctx.FinishSpan()

	kvSchemaKey := kv.NewKeyWith(kv.SchemaKeyHeader, document.NewKeyWith(txn.Database().Name()))
	return txn.kv.RemoveRange(kvSchemaKey)
}
