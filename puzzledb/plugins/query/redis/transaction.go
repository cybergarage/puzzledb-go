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

package redis

import (
	"errors"
	"fmt"

	"github.com/cybergarage/go-redis/redis"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

type Transaction struct {
	store.Transaction
	redis.DatabaseID
}

// SetKeyObject sets the objects with the specified key.
func (txn *Transaction) SetKeyObject(ctx context.Context, key string, val any) error {
	docKey := NewDocumentKeyWith(txn.DatabaseID, key)
	return txn.InsertDocument(ctx, docKey, val)
}

// GetKeyObjects returns the objects with the specified key.
func (txn *Transaction) GetKeyObjects(ctx context.Context, key string) ([]any, error) {
	docKey := NewDocumentKeyWith(txn.DatabaseID, key)
	rs, err := txn.FindDocuments(ctx, docKey)
	if err != nil {
		return nil, err
	}
	return rs.Objects(), nil
}

// GetKeyObject returns the object with the specified key.
func (txn *Transaction) GetKeyObject(ctx context.Context, key string) (any, error) {
	docKey := NewDocumentKeyWith(txn.DatabaseID, key)
	rs, err := txn.FindDocuments(ctx, docKey)
	if err != nil {
		return nil, err
	}

	objs := rs.Objects()
	if len(objs) == 0 {
		return nil, document.NewErrObjectNotFound(docKey)
	}

	if len(objs) != 1 {
		return nil, fmt.Errorf("%w: multiple objects are found (%d)", ErrInvalid, len(objs))
	}

	return objs[0], nil
}

// SetKeyHashObject sets the objects with the specified key.
func (txn *Transaction) SetKeyHashObject(ctx context.Context, key string, val HashObject) error {
	docKey := NewDocumentKeyWith(txn.DatabaseID, key)
	return txn.InsertDocument(ctx, docKey, val)
}

// GetKeyHashObject returns the hash objects with the specified key.
func (txn *Transaction) GetKeyHashObject(ctx context.Context, key string) (HashObject, error) {
	obj, err := txn.GetKeyObject(ctx, key)
	if err != nil {
		return nil, err
	}
	hobj, ok := obj.(HashObject)
	if !ok {
		return nil, fmt.Errorf("%w object type (%T)", ErrInvalid, obj)
	}
	return hobj, nil
}

// CancelWithError cancels the transaction with an error.
func (txn *Transaction) CancelWithError(ctx context.Context, err error) error {
	if txErr := txn.Cancel(ctx); txErr != nil {
		return errors.Join(err, txErr)
	}
	return err
}
