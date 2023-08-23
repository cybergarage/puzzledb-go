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

package sql

import (
	"errors"

	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	plugins "github.com/cybergarage/puzzledb-go/puzzledb/plugins/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// Service represents a new MySQL service instance.
type Service struct {
	*plugins.BaseService
}

// NewService returns a new MySQL service.
func NewService() *Service {
	service := &Service{
		BaseService: plugins.NewBaseService(),
	}
	return service
}

// cancelTransactionWithError cancels the specified transaction with the specified error.
func (service *Service) cancelTransactionWithError(ctx context.Context, txn store.Transaction, err error) error {
	if txErr := txn.Cancel(ctx); txErr != nil {
		return txErr
	}
	return err
}

func (service *Service) insertSecondaryIndexes(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj Object, prKey document.Key) error {
	idxes, err := schema.SecondaryIndexes()
	if err != nil {
		return err
	}
	for _, idx := range idxes {
		err := service.insertSecondaryIndex(ctx, conn, txn, schema, obj, idx, prKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *Service) insertSecondaryIndex(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj Object, idx document.Index, prKey document.Key) error {
	dbName := conn.Database()
	secKey, err := NewKeyFromIndex(dbName, schema, idx, obj)
	if err != nil {
		return err
	}
	return txn.InsertIndex(ctx, secKey, prKey)
}

func (service *Service) selectDocumentObjects(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, cond *query.Condition, orderby *query.OrderBy, limit *query.Limit) (store.ResultSet, error) {
	docKey, docKeyType, err := NewKeyFromCond(conn.Database(), schema, cond)
	if err != nil {
		return nil, err
	}

	opts := []store.Option{}
	opts = append(opts, NewLimitWith(limit)...)
	opts = append(opts, NewOrderWith(orderby)...)

	switch docKeyType {
	case document.PrimaryIndex:
		return txn.FindDocuments(ctx, docKey, opts...)
	case document.SecondaryIndex:
		return txn.FindDocumentsByIndex(ctx, docKey, opts...)
	}
	return nil, newErrIndexTypeNotSupported(docKeyType)
}

func (service *Service) removeSecondaryIndexes(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj Object) error {
	idxes, err := schema.SecondaryIndexes()
	if err != nil {
		return err
	}
	var lastErr error
	for _, idx := range idxes {
		err := service.removeSecondaryIndex(ctx, conn, txn, schema, obj, idx)
		if err != nil && !errors.Is(err, store.ErrNotExist) {
			lastErr = err
		}
	}
	return lastErr
}

func (service *Service) removeSecondaryIndex(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj Object, idx document.Index) error {
	dbName := conn.Database()
	secKey, err := NewKeyFromIndex(dbName, schema, idx, obj)
	if err != nil {
		return err
	}
	return txn.RemoveIndex(ctx, secKey)
}