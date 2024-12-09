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

// Transact returns a transaction object.
func (service *Service) Transact(conn Conn, db store.Database, write bool) (store.Transaction, error) {
	// Checks the transaction is already started.
	txn, err := service.GetTransaction(conn, db)
	if err == nil {
		return txn, nil
	}
	if !errors.Is(err, ErrNotExist) {
		return nil, err
	}
	// Starts a new transaction.
	txn, err = db.Transact(write)
	if err != nil {
		return nil, err
	}
	txn.SetAutoCommit(true)
	return txn, nil
}

// CommitTransaction commits the specified transaction.
func (service *Service) CommitTransaction(ctx context.Context, conn Conn, db store.Database, txn store.Transaction) error {
	if txErr := txn.Commit(ctx); txErr != nil {
		return txErr
	}

	err := service.RemoveTransaction(conn, db)
	if err != nil {
		return err
	}

	return nil
}

// CancelTransaction cancels the specified transaction.
func (service *Service) CancelTransaction(ctx context.Context, conn Conn, db store.Database, txn store.Transaction) error {
	if txErr := txn.Cancel(ctx); txErr != nil {
		return txErr
	}

	err := service.RemoveTransaction(conn, db)
	if err != nil {
		return err
	}

	return nil
}

// CancelTransactionWithError cancels the specified transaction with the specified error.
func (service *Service) CancelTransactionWithError(ctx context.Context, conn Conn, db store.Database, txn store.Transaction, err error) error {
	if txErr := txn.Cancel(ctx); txErr != nil {
		return errors.Join(err, txErr)
	}

	rmErr := service.RemoveTransaction(conn, db)
	if rmErr != nil {
		return errors.Join(err, rmErr)
	}

	return err
}

// SelectDatabaseObjects returns a result set of the specified database objects.
func (service *Service) SelectDocumentObjects(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, cond query.Condition, orderby query.OrderBy, limit query.Limit) (store.ResultSet, error) {
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
		return txn.FindObjects(ctx, docKey, opts...)
	case document.SecondaryIndex:
		return txn.FindObjectsByIndex(ctx, docKey, opts...)
	}
	return nil, newErrIndexTypeNotSupported(docKeyType)
}

// InsertSecondaryIndexes inserts secondary indexes for the specified object.
func (service *Service) InsertSecondaryIndexes(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj document.MapObject, prKey document.Key) error {
	insertSecondaryIndex := func(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj document.MapObject, secIdx document.Index, prIdx document.Index, prKey document.Key) error {
		dbName := conn.Database()
		secKey, err := NewDocumentKeyFromIndexes(dbName, schema.Name(), obj, secIdx, prIdx)
		if err != nil {
			return err
		}
		return txn.InsertIndex(ctx, secKey, prKey)
	}

	prIdx, err := schema.PrimaryIndex()
	if err != nil {
		return err
	}
	secIdxes, err := schema.SecondaryIndexes()
	if err != nil {
		return err
	}
	for _, secIdx := range secIdxes {
		err := insertSecondaryIndex(ctx, conn, txn, schema, obj, secIdx, prIdx, prKey)
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
		secKey, err := NewDocumentKeyFromIndexes(dbName, schema.Name(), obj, idx)
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
func (service *Service) UpdateObject(ctx context.Context, conn Conn, txn store.Transaction, schema document.Schema, obj any, updateCols query.Columns) error {
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
		if fn, ok := updateCol.IsFunction(); ok {
			v, err := fn.Execute(updateCol, docObj)
			if err != nil {
				return err
			}
			updateVal = v
		} else {
			if !updateCol.HasValue() {
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

	err = txn.UpdateObject(ctx, docKey, docObj)
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
	err := txn.RemoveObject(ctx, docKey)
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
	rs, err := txn.FindObjects(ctx, docKey)
	if err != nil {
		return err
	}
	for rs.Next() {
		rsDoc, err := rs.Document()
		if err != nil {
			return err
		}
		obj, err := document.NewMapObjectFrom(rsDoc.Object())
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
