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
	return txn.InsertObject(ctx, docKey, val)
}

// GetKeyObjects returns the objects with the specified key.
func (txn *Transaction) GetKeyObjects(ctx context.Context, key string) ([]any, error) {
	docKey := NewDocumentKeyWith(txn.DatabaseID, key)
	rs, err := txn.FindObjects(ctx, docKey)
	if err != nil {
		return nil, err
	}
	return store.ReadAll(rs)
}

// GetKeyObject returns the object with the specified key.
func (txn *Transaction) GetKeyObject(ctx context.Context, key string) (any, error) {
	docKey := NewDocumentKeyWith(txn.DatabaseID, key)
	rs, err := txn.FindObjects(ctx, docKey)
	if err != nil {
		return nil, err
	}

	objs, err := store.ReadAll(rs)
	if err != nil {
		return nil, err
	}
	if len(objs) == 0 {
		return nil, document.NewErrObjectNotExist(docKey)
	}

	if len(objs) != 1 {
		return nil, fmt.Errorf("%w: multiple objects are found (%d)", ErrInvalid, len(objs))
	}

	return objs[0], nil
}

// GetKeyString returns the string with the specified key.
func (txn *Transaction) GetKeyString(ctx context.Context, key string) (string, error) {
	obj, err := txn.GetKeyObject(ctx, key)
	if err != nil {
		return "", err
	}

	switch v := obj.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	}

	return fmt.Sprintf("%v", obj), nil
}

// SetKeyHashObject sets the objects with the specified key.
func (txn *Transaction) SetKeyHashObject(ctx context.Context, key string, val HashObject) error {
	docKey := NewDocumentKeyWith(txn.DatabaseID, key)
	return txn.InsertObject(ctx, docKey, val)
}

// GetKeyHashObject returns the hash objects with the specified key.
func (txn *Transaction) GetKeyHashObject(ctx context.Context, key string) (HashObject, error) {
	keyObj, err := txn.GetKeyObject(ctx, key)
	if err != nil {
		return nil, err
	}

	switch obj := keyObj.(type) {
	case map[string]string:
		return make(HashObject), nil
	case map[any]any:
		hashObj := make(HashObject)
		for k, v := range obj {
			hashObj[fmt.Sprintf("%v", k)] = fmt.Sprintf("%v", v)
		}
		return hashObj, nil
	}

	return nil, fmt.Errorf("%w object type (%T)", ErrInvalid, keyObj)
}

// CancelWithError cancels the transaction with an error.
func (txn *Transaction) CancelWithError(ctx context.Context, err error) error {
	if txErr := txn.Cancel(ctx); txErr != nil {
		return errors.Join(err, txErr)
	}
	return err
}
