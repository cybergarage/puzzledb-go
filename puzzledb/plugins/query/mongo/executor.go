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

package mongo

import (
	"errors"
	"time"

	"github.com/cybergarage/go-mongo/mongo"
	"github.com/cybergarage/go-mongo/mongo/bson"
	"github.com/cybergarage/puzzledb-go/puzzledb/context"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func (service *Service) createDocumentAllKey(txn store.Transaction, database string, collection string) document.Key {
	return document.NewKeyWith(database, collection)
}

func (service *Service) createDocumentKey(txn store.Transaction, database string, collection string, key string, val any) document.Key {
	return document.NewKeyWith(database, collection, key, val)
}

func (service *Service) createObjectKey(txn store.Transaction, database string, collection string, objID any) document.Key {
	return service.createDocumentKey(txn, database, collection, ObjectID, objID)
}

func (service *Service) createidxKey(txn store.Transaction, database string, collection string, key string, val any) document.Key {
	return service.createDocumentKey(txn, database, collection, key, val)
}

// Insert hadles OP_INSERT and 'insert' query of OP_MSG or OP_QUERY.
func (service *Service) Insert(conn *mongo.Conn, q *mongo.Query) (int32, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Insert")
	defer ctx.FinishSpan()
	now := time.Now()

	db, err := service.GetDatabase(ctx, q.Database())
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	txn, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	nInserted := int32(0)

	queryDocs := q.Documents()
	for _, queryDoc := range queryDocs {
		err = service.insertDocument(ctx, txn, q, queryDoc)
		if err != nil {
			if err := txn.Cancel(ctx); err != nil {
				return 0, err
			}
			return 0, err
		}
		nInserted++
	}

	err = txn.Commit(ctx)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	mInsertLatency.Observe(float64(time.Since(now).Milliseconds()))

	return nInserted, nil
}

func (service *Service) insertDocument(ctx context.Context, txn store.Transaction, q *mongo.Query, bsonDoc mongo.Document) error {
	// Inserts the document with the primary key

	doc, err := service.EncodeBSON(bsonDoc)
	if err != nil {
		return err
	}

	objID, err := LookupBSONDocumentObjectID(bsonDoc)
	if err != nil {
		return err
	}

	docKey := service.createObjectKey(txn, q.Database(), q.Collection(), objID)
	err = txn.InsertDocument(ctx, docKey, doc)
	if err != nil {
		return err
	}

	// Creates the secondary indexes for the all elements

	err = service.insertDocumentIndexes(ctx, txn, q.Database(), q.Collection(), docKey, doc)
	if err != nil {
		return err
	}

	return err
}

func (service *Service) insertDocumentIndexes(ctx context.Context, txn store.Transaction, db string, col string, docKey document.Key, v any) error {
	switch vmap := v.(type) { //nolint:all
	case map[string]any:
		for secKey, secVal := range vmap {
			if secKey == ObjectID {
				continue
			}
			err := service.insertDocumentIndex(ctx, txn, db, col, secKey, secVal, docKey)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return newErrBSONTypeNotSupported(v)
}

func (service *Service) insertDocumentIndex(ctx context.Context, txn store.Transaction, db string, col string, secKey string, secVal any, docKey document.Key) error {
	idxKey := service.createidxKey(txn, db, col, secKey, secVal)
	return txn.InsertIndex(ctx, idxKey, docKey)
}

// Find hadles 'find' query of OP_MSG or OP_QUERY.
func (service *Service) Find(conn *mongo.Conn, q *mongo.Query) ([]bson.Document, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Find")
	defer ctx.FinishSpan()
	now := time.Now()

	db, err := service.GetDatabase(ctx, q.Database())
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	txn, err := db.Transact(false)
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	foundDoc, err := service.findDocuments(ctx, txn, q)
	if err != nil {
		if err := txn.Cancel(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	mFindLatency.Observe(float64(time.Since(now).Milliseconds()))

	return foundDoc, nil
}

func (service *Service) findDocumentObjects(ctx context.Context, txn store.Transaction, q *mongo.Query) ([]document.Object, error) {
	matchedDocs := []document.Object{}
	if q.HasConditions() {
		// Finds the documents by the conditions
		for _, cond := range q.Conditions() {
			condElems, err := cond.Elements()
			if err != nil {
				return nil, mongo.NewQueryError(q)
			}
			for _, condElem := range condElems {
				key := condElem.Key()
				bsonVal := condElem.Value()
				val, err := EncodeBSONValue(bsonVal)
				if err != nil {
					return nil, err
				}
				idxKey := service.createDocumentKey(txn, q.Database(), q.Collection(), key, val)
				var objs []document.Object
				if isPrimaryKey(key) {
					rs, err := txn.FindDocuments(ctx, idxKey)
					if err != nil {
						return nil, err
					}
					objs = rs.Objects()
				} else {
					rs, err := txn.FindDocumentsByIndex(ctx, idxKey)
					if err != nil {
						return nil, err
					}
					objs = rs.Objects()
				}
				matchedDocs = append(matchedDocs, objs...)
			}
		}
	} else {
		// Finds all documents
		idxKey := service.createDocumentAllKey(txn, q.Database(), q.Collection())
		var objs []document.Object
		rs, err := txn.FindDocuments(ctx, idxKey)
		if err != nil {
			return nil, err
		}
		objs = rs.Objects()
		matchedDocs = append(matchedDocs, objs...)
	}
	return matchedDocs, nil
}

func (service *Service) findDocuments(ctx context.Context, txn store.Transaction, q *mongo.Query) ([]bson.Document, error) {
	docs, err := service.findDocumentObjects(ctx, txn, q)
	if err != nil {
		return nil, err
	}
	bsonDocs := []bson.Document{}
	for _, matchedDoc := range docs {
		bsonDoc, err := DecodeBSONDocument(matchedDoc)
		if err != nil {
			return nil, err
		}
		bsonDocs = append(bsonDocs, bsonDoc)
	}
	return bsonDocs, nil
}

// Update hadles OP_UPDATE and 'update' query of OP_MSG or OP_QUERY.
func (service *Service) Update(conn *mongo.Conn, q *mongo.Query) (int32, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Update")
	defer ctx.FinishSpan()
	now := time.Now()

	db, err := service.GetDatabase(ctx, q.Database())
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	txn, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	foundDocs, err := service.Find(conn, q)
	if err != nil {
		if err := txn.Cancel(ctx); err != nil {
			return 0, err
		}
		return 0, err
	}

	nUpdated, err := service.updateDocumentsByQuery(ctx, txn, foundDocs, q)
	if err != nil {
		if err := txn.Cancel(ctx); err != nil {
			return 0, err
		}
		return 0, err
	}

	err = txn.Commit(ctx)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	mUpdateLatency.Observe(float64(time.Since(now).Milliseconds()))

	return int32(nUpdated), nil
}

func (service *Service) updateDocumentsByQuery(ctx context.Context, txn store.Transaction, bsonDocs []bson.Document, q *mongo.Query) (int32, error) {
	nUpdated := 0
	for _, bsonDoc := range bsonDocs {
		err := service.updateDocumentByQuery(ctx, txn, bsonDoc, q)
		if err != nil {
			return 0, err
		}
		nUpdated++
	}
	return int32(nUpdated), nil
}

func (service *Service) updateDocumentByQuery(ctx context.Context, txn store.Transaction, bsonDoc bson.Document, q *mongo.Query) error {
	updateBSONDocs := q.Documents()

	// Removes current secondary indexes for the all elements

	for _, updateBSONDoc := range updateBSONDocs {
		err := service.deleteUpdateDocumentIndexes(ctx, txn, q.Database(), q.Collection(), bsonDoc, updateBSONDoc)
		if err != nil && !errors.Is(err, store.ErrNotExist) {
			return err
		}
	}

	// Updates the matched doucments by the query

	updatedBSONDoc, err := UpdateBSONDocument(bsonDoc, updateBSONDocs)
	if err != nil {
		return err
	}
	objID, err := LookupBSONDocumentObjectID(updatedBSONDoc)
	if err != nil {
		return err
	}
	updatedDoc, err := service.EncodeBSON(updatedBSONDoc)
	if err != nil {
		return err
	}
	docKey := service.createObjectKey(txn, q.Database(), q.Collection(), objID)
	err = txn.UpdateDocument(ctx, docKey, updatedDoc)
	if err != nil {
		return err
	}

	// Updates the secondary indexes for the all elements

	for _, updateBSONDoc := range updateBSONDocs {
		updateDoc, err := service.EncodeBSON(updateBSONDoc)
		if err != nil {
			return err
		}
		err = service.insertDocumentIndexes(ctx, txn, q.Database(), q.Collection(), docKey, updateDoc)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete hadles OP_DELETE and 'delete' query of OP_MSG or OP_QUERY.
func (service *Service) Delete(conn *mongo.Conn, q *mongo.Query) (int32, error) {
	ctx := context.NewContextWith(conn.SpanContext())
	ctx.StartSpan("Delete")
	defer ctx.FinishSpan()
	now := time.Now()

	db, err := service.GetDatabase(ctx, q.Database())
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	txn, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	foundDocs, err := service.Find(conn, q)
	if err != nil {
		if err := txn.Cancel(ctx); err != nil {
			return 0, err
		}
		return 0, err
	}

	nDeleted, err := service.deleteDocuments(ctx, txn, q.Database(), q.Collection(), foundDocs)
	if err != nil {
		if err := txn.Cancel(ctx); err != nil {
			return 0, err
		}
		return 0, err
	}

	err = txn.Commit(ctx)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	mDeleteLatency.Observe(float64(time.Since(now).Milliseconds()))

	return int32(nDeleted), nil
}

func (service *Service) deleteDocuments(ctx context.Context, txn store.Transaction, db string, col string, bsonDocs []bson.Document) (int32, error) {
	nDeleted := 0
	for _, bsonDoc := range bsonDocs {
		err := service.deleteDocument(ctx, txn, db, col, bsonDoc)
		if err != nil {
			return 0, err
		}
		nDeleted++
	}
	return int32(nDeleted), nil
}

func (service *Service) deleteDocument(ctx context.Context, txn store.Transaction, db string, col string, bsonDoc bson.Document) error {
	objID, err := LookupBSONDocumentObjectID(bsonDoc)
	if err != nil {
		return err
	}
	docKey := service.createObjectKey(txn, db, col, objID)
	err = txn.RemoveDocument(ctx, docKey)
	if err != nil {
		return err
	}

	// Removes the secondary indexes for the all elements.

	doc, err := service.EncodeBSON(bsonDoc)
	if err != nil {
		return err
	}

	err = service.deleteDocumentIndexes(ctx, txn, db, col, doc)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) deleteUpdateDocumentIndexes(ctx context.Context, txn store.Transaction, db string, col string, bsonDoc bson.Document, updateBSONDoc bsoncore.Document) error {
	updateBSONElems, err := updateBSONDoc.Elements()
	if err != nil {
		return err
	}
	for _, updateBSONElem := range updateBSONElems {
		updateBSONKey := updateBSONElem.Key()
		updateBSONVal, err := bsonDoc.LookupErr(updateBSONKey)
		if err != nil {
			continue
		}
		idxKey := service.createidxKey(txn, db, col, updateBSONKey, updateBSONVal.Data)
		err = txn.RemoveIndex(ctx, idxKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *Service) deleteDocumentIndexes(ctx context.Context, txn store.Transaction, db string, col string, v any) error {
	switch vmap := v.(type) { //nolint:all
	case map[string]any:
		for key, val := range vmap {
			err := service.deleteDocumentIndex(ctx, txn, db, col, key, val)
			if err != nil && !errors.Is(err, store.ErrNotExist) {
				return err
			}
		}
		return nil
	}
	return newErrBSONTypeNotSupported(v)
}

func (service *Service) deleteDocumentIndex(ctx context.Context, txn store.Transaction, db string, col string, key string, val any) error {
	idxKey := service.createidxKey(txn, db, col, key, val)
	return txn.RemoveIndex(ctx, idxKey)
}
