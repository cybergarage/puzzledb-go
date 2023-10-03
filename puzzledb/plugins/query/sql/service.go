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

// Service represents a new SQL service instance.
type Service struct {
	*plugins.BaseService
	ConnectionMap
}

// NewService returns a new SQL service.
func NewService() *Service {
	service := &Service{
		BaseService:   plugins.NewBaseService(),
		ConnectionMap: NewConnectionMap(),
	}
	return service
}

// CancelTransactionWithError cancels the specified transaction with the specified error.
func (service *Service) CancelTransactionWithError(ctx context.Context, txn store.Transaction, err error) error {
	if txErr := txn.Cancel(ctx); txErr != nil {
		return txErr
	}
	return err
}

// SelectDatabaseObjects returns a result set of the specified database objects.
func (service *Service) SelectDocumentObjects(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, cond *query.Condition, orderby *query.OrderBy, limit *query.Limit) (store.ResultSet, error) {
	docKey, docKeyType, err := NewDocumentKeyFromCond(conn.Database(), schema, cond)
	if err != nil {
		return nil, err
	}

	opts := []store.Option{}
	if limit != nil {
		opts = append(opts, NewLimitWith(limit)...)
	}
	if orderby != nil {
		opts = append(opts, NewOrderWith(orderby)...)
	}

	switch docKeyType {
	case document.PrimaryIndex:
		return txn.FindDocuments(ctx, docKey, opts...)
	case document.SecondaryIndex:
		return txn.FindDocumentsByIndex(ctx, docKey, opts...)
	}
	return nil, newErrIndexTypeNotSupported(docKeyType)
}

// InsertSecondaryIndexes inserts secondary indexes for the specified object.
func (service *Service) InsertSecondaryIndexes(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj document.MapObject, prKey document.Key) error {
	insertSecondaryIndex := func(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj document.MapObject, idx document.Index, prKey document.Key) error {
		dbName := conn.Database()
		secKey, err := NewDocumentKeyFromIndex(dbName, schema, idx, obj)
		if err != nil {
			return err
		}
		return txn.InsertIndex(ctx, secKey, prKey)
	}

	idxes, err := schema.SecondaryIndexes()
	if err != nil {
		return err
	}
	for _, idx := range idxes {
		err := insertSecondaryIndex(ctx, conn, txn, schema, obj, idx, prKey)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveSecondaryIndexes removes secondary indexes for the specified object.
func (service *Service) RemoveSecondaryIndexes(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj document.MapObject) error {
	removeSecondaryIndex := func(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj document.MapObject, idx document.Index) error {
		dbName := conn.Database()
		secKey, err := NewDocumentKeyFromIndex(dbName, schema, idx, obj)
		if err != nil {
			return err
		}
		return txn.RemoveIndex(ctx, secKey)
	}

	idxes, err := schema.SecondaryIndexes()
	if err != nil {
		return err
	}
	var lastErr error
	for _, idx := range idxes {
		err := removeSecondaryIndex(ctx, conn, txn, schema, obj, idx)
		if err != nil && !errors.Is(err, store.ErrNotExist) {
			lastErr = err
		}
	}
	return lastErr
}

// UpdateDocument updates the specified object.
func (service *Service) UpdateDocument(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj any, updateCols query.ColumnList) error {
	docObj, err := document.NewMapObjectFrom(obj)
	if err != nil {
		return err
	}

	// Removes current secondary indexes
	err = service.RemoveSecondaryIndexes(ctx, conn, txn, schema, docObj)
	if err != nil {
		return err
	}

	// Updates object
	dbName := conn.Database()
	for _, updateCol := range updateCols.Columns() {
		updateColName := updateCol.Name()

		var updateVal any
		if exe := updateCol.Executor(); exe != nil {
			v, err := updateCol.ExecuteUpdator(docObj)
			if err != nil {
				return err
			}
			updateVal = v
		} else {
			if !updateCol.HasLiteral() {
				continue
			}
			updateVal = updateCol.Value()
		}

		v, err := document.NewValueForSchema(schema, updateColName, updateVal)
		if err != nil {
			return err
		}
		docObj[updateColName] = v
	}
	docKey, err := NewDocumentKeyFromObject(dbName, schema, docObj)
	if err != nil {
		return err
	}

	err = txn.UpdateDocument(ctx, docKey, docObj)
	if err != nil {
		return err
	}

	// Inserts new secondary indexes.
	err = service.InsertSecondaryIndexes(ctx, conn, txn, schema, docObj, docKey)
	if err != nil {
		return err
	}
	return nil
}

// DeleteDocument deletes the specified object.
func (service *Service) DeleteDocument(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, docKey document.Key) error {
	err := txn.RemoveDocument(ctx, docKey)
	if err != nil {
		return err
	}
	// Removes secondary indexes
	idxes, err := schema.SecondaryIndexes()
	if err != nil {
		return err
	}
	if len(idxes) == 0 {
		return nil
	}
	rs, err := txn.FindDocuments(ctx, docKey)
	if err != nil {
		return err
	}
	for rs.Next() {
		docObj := rs.Object()
		obj, err := document.NewMapObjectFrom(docObj)
		if err != nil {
			return err
		}
		err = service.RemoveSecondaryIndexes(ctx, conn, txn, schema, obj)
		if err != nil {
			return err
		}
	}
	return nil
}
