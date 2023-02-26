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

// CreateSchema creates a new schema.
func (txn *transaction) CreateSchema(schema store.Schema) error {
	kcColKey := txn.createSchemaKey(schema)
	var encSchema bytes.Buffer
	err := txn.Encode(&encSchema, schema.Data())
	if err != nil {
		return err
	}
	kvObj := kv.Object{
		Key:   kcColKey,
		Value: encSchema.Bytes(),
	}
	return txn.kv.Set(&kvObj)
}

// GetSchema returns the specified schema.
func (txn *transaction) GetSchema(name string) (store.Schema, error) {
	return nil, nil
}

func (txn *transaction) createSchemaKey(schema store.Schema) store.Key {
	colKey := document.NewKeyWith(txn.Database().Name(), schema.Name())
	return kv.NewKeyWith(kv.SchemaKeyHeader, colKey)
}
