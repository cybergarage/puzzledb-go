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

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"github.com/cybergarage/puzzledb-go/puzzledb/store/kv"
)

func (txn *transaction) createSchemaKey(schema string) store.Key {
	colKey := document.NewKeyWith(txn.Database().Name(), schema)
	return kv.NewKeyWith(kv.SchemaKeyHeader, colKey)
}

// CreateSchema creates a new schema.
func (txn *transaction) CreateSchema(schema store.Schema) error {
	kvSchemaKey := txn.createSchemaKey(schema.Name())
	var encSchema bytes.Buffer
	err := txn.Encode(&encSchema, schema.Data())
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
func (txn *transaction) GetSchema(name string) (store.Schema, error) {
	kvSchemaKey := txn.createSchemaKey(name)
	kvRs, err := txn.kv.Get(kvSchemaKey)
	if err != nil {
		return nil, err
	}
	if !kvRs.Next() {
		return nil, store.NewSchemaNotExistError(name)
	}
	kvObj := kvRs.Object()
	obj, err := txn.Decode(bytes.NewReader(kvObj.Value))
	if err != nil {
		return nil, err
	}
	return document.NewSchemaWith(obj)
}

// RemoveSchema removes the specified schema.
func (txn *transaction) RemoveSchema(name string) error {
	kvSchemaKey := txn.createSchemaKey(name)
	return txn.kv.Remove(kvSchemaKey)
}
