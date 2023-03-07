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
	"github.com/cybergarage/go-mongo/mongo"
	"github.com/cybergarage/go-mongo/mongo/bson"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

func (service *Service) createStoreKey(tx store.Transaction, q *mongo.Query, key string, val any) document.Key {
	return document.NewKeyWith(q.Database, q.Collection, key, val)
}

func (service *Service) createDocumentKey(tx store.Transaction, q *mongo.Query, objID any) document.Key {
	return service.createStoreKey(tx, q, ObjectID, objID)
}

func (service *Service) createIndexKey(tx store.Transaction, q *mongo.Query, key string, val any) document.Key {
	return service.createStoreKey(tx, q, key, val)
}

// Insert hadles OP_INSERT and 'insert' query of OP_MSG or OP_QUERY.
func (service *Service) Insert(conn *mongo.Conn, q *mongo.Query) (int32, error) {
	db, err := service.GetDatabase(q.Database)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	tx, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	nInserted := int32(0)

	queryDocs := q.GetDocuments()
	for _, queryDoc := range queryDocs {
		err = service.insertDocument(tx, q, queryDoc)
		if err != nil {
			if err := tx.Cancel(); err != nil {
				return 0, err
			}
			return 0, err
		}
		nInserted++
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	return nInserted, nil
}

func (service *Service) insertDocument(tx store.Transaction, q *mongo.Query, bsonDoc mongo.Document) error {
	// Inserts the document with the primary key

	doc, err := service.EncodeBSON(bsonDoc)
	if err != nil {
		return err
	}

	objID, err := LookupBSONDocumentObjectID(bsonDoc)
	if err != nil {
		return err
	}

	docKey := service.createDocumentKey(tx, q, objID)
	err = tx.InsertDocument(docKey, doc)
	if err != nil {
		return err
	}

	// Creates the secondary indexes for the all elements

	err = service.updateDocumentIndexes(tx, q, docKey, doc)
	if err != nil {
		return err
	}

	return err
}

func (service *Service) updateDocumentIndexes(tx store.Transaction, q *mongo.Query, docKey document.Key, v any) error {
	switch vmap := v.(type) { //nolint:all
	case map[string]any:
		for key, val := range vmap {
			indexKey := service.createIndexKey(tx, q, key, val)
			err := tx.InsertIndex(indexKey, docKey)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return newErrBSONTypeNotSupported(v)
}

// Find hadles 'find' query of OP_MSG or OP_QUERY.
func (service *Service) Find(conn *mongo.Conn, q *mongo.Query) ([]bson.Document, error) {
	db, err := service.GetDatabase(q.Database)
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	tx, err := db.Transact(false)
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	foundDoc, err := service.findDocuments(tx, q)
	if err != nil {
		if err := tx.Cancel(); err != nil {
			return nil, err
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	return foundDoc, nil
}

func (service *Service) findDocumentObjects(tx store.Transaction, q *mongo.Query) ([]document.Object, error) {
	matchedDocs := []document.Object{}
	for _, cond := range q.GetConditions() {
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
			idxKey := service.createStoreKey(tx, q, key, val)
			var objs []document.Object
			if isPrimaryKey(key) {
				rs, err := tx.FindDocuments(idxKey)
				if err != nil {
					return nil, err
				}
				objs = rs.Objects()
			} else {
				rs, err := tx.FindDocumentsByIndex(idxKey)
				if err != nil {
					return nil, err
				}
				objs = rs.Objects()
			}
			matchedDocs = append(matchedDocs, objs...)
		}
	}
	return matchedDocs, nil
}

func (service *Service) findDocuments(tx store.Transaction, q *mongo.Query) ([]bson.Document, error) {
	docs, err := service.findDocumentObjects(tx, q)
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
	db, err := service.GetDatabase(q.Database)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	tx, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	foundDocs, err := service.Find(conn, q)
	if err != nil {
		if err := tx.Cancel(); err != nil {
			return 0, err
		}
		return 0, err
	}

	nUpdated, err := service.updateDocuments(tx, q, foundDocs)
	if err != nil {
		if err := tx.Cancel(); err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	return int32(nUpdated), nil
}

func (service *Service) updateDocuments(tx store.Transaction, q *mongo.Query, bsonDocs []bson.Document) (int32, error) {
	nUpdated := 0
	for _, bsonDoc := range bsonDocs {
		err := service.updateDocument(tx, q, bsonDoc)
		if err != nil {
			return 0, err
		}
		nUpdated++
	}
	return int32(nUpdated), nil
}

func (service *Service) updateDocument(tx store.Transaction, q *mongo.Query, bsonDoc bson.Document) error {
	// Updates the matched doucments by the query

	updateBSONDocs := q.GetDocuments()
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
	docKey := service.createDocumentKey(tx, q, objID)
	err = tx.UpdateDocument(docKey, updatedDoc)
	if err != nil {
		return err
	}

	// Updates the secondary indexes for the all elements

	for _, updateBSONDoc := range updateBSONDocs {
		updateDoc, err := service.EncodeBSON(updateBSONDoc)
		if err != nil {
			return err
		}
		err = service.updateDocumentIndexes(tx, q, docKey, updateDoc)
		if err != nil {
			return err
		}
	}

	return nil
}

// Delete hadles OP_DELETE and 'delete' query of OP_MSG or OP_QUERY.
func (service *Service) Delete(conn *mongo.Conn, q *mongo.Query) (int32, error) {
	db, err := service.GetDatabase(q.Database)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	tx, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	foundDocs, err := service.Find(conn, q)
	if err != nil {
		if err := tx.Cancel(); err != nil {
			return 0, err
		}
		return 0, err
	}

	nDeleted, err := service.deleteDocuments(tx, q, foundDocs)
	if err != nil {
		if err := tx.Cancel(); err != nil {
			return 0, err
		}
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	return int32(nDeleted), nil
}

func (service *Service) deleteDocuments(tx store.Transaction, q *mongo.Query, bsonDocs []bson.Document) (int32, error) {
	nDeleted := 0
	for _, bsonDoc := range bsonDocs {
		err := service.deleteDocument(tx, q, bsonDoc)
		if err != nil {
			return 0, err
		}
		nDeleted++
	}
	return int32(nDeleted), nil
}

func (service *Service) deleteDocument(tx store.Transaction, q *mongo.Query, bsonDoc bson.Document) error {
	objID, err := LookupBSONDocumentObjectID(bsonDoc)
	if err != nil {
		return err
	}
	docKey := service.createDocumentKey(tx, q, objID)
	err = tx.RemoveDocument(docKey)
	if err != nil {
		return err
	}

	// TODO: Removes the secondary indexes for the all elements.

	doc, err := service.EncodeBSON(bsonDoc)
	if err != nil {
		return err
	}

	err = service.deleteDocumentIndexes(tx, q, docKey, doc)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) deleteDocumentIndexes(tx store.Transaction, q *mongo.Query, docKey document.Key, v any) error {
	switch vmap := v.(type) { //nolint:all
	case map[string]any:
		for key, val := range vmap {
			indexKey := service.createIndexKey(tx, q, key, val)
			err := tx.RemoveIndex(indexKey)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return newErrBSONTypeNotSupported(v)
}
