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

func (service *Service) updateDocumentIndexes(tx store.Transaction, q *mongo.Query, docKey document.Key, v map[string]any) error {
	for key, val := range v {
		indexKey := service.createIndexKey(tx, q, key, val)
		err := tx.InsertIndex(indexKey, docKey)
		if err != nil {
			return err
		}
	}
	return nil
}

// Insert hadles OP_INSERT and 'insert' query of OP_MSG or OP_QUERY.
func (service *Service) Insert(q *mongo.Query) (int32, error) {
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
			tx.Cancel()
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
	document.NewKeyWith(q.Database, q.Collection, ObjectID, objID)
	err = tx.InsertDocument(docKey, doc)
	if err != nil {
		return err
	}

	// Makes the secondary indexes for the all elements

	switch v := doc.(type) {
	case map[string]any:
		service.updateDocumentIndexes(tx, q, docKey, v)
	}

	return err
}

// Find hadles 'find' query of OP_MSG or OP_QUERY.
func (service *Service) Find(q *mongo.Query) ([]bson.Document, error) {
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
		tx.Cancel()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	return foundDoc, nil
}

func (service *Service) findDocuments(tx store.Transaction, q *mongo.Query) ([]bson.Document, error) {
	foundDocs := make([]bson.Document, 0)
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
				objs, err = tx.SelectDocuments(idxKey)
				if err != nil {
					return nil, err
				}
			} else {
				objs, err = tx.SelectDocumentsByIndex(idxKey)
				if err != nil {
					return nil, err
				}
			}
			for _, obj := range objs {
				bsonDoc, err := DecodeBSONDocument(obj)
				if err != nil {
					return nil, err
				}
				foundDocs = append(foundDocs, bsonDoc)
			}

		}
	}
	return foundDocs, nil
}

// Update hadles OP_UPDATE and 'update' query of OP_MSG or OP_QUERY.
func (service *Service) Update(q *mongo.Query) (int32, error) {
	db, err := service.GetDatabase(q.Database)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	tx, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	foundDocs, err := service.Find(q)
	if err != nil {
		tx.Cancel()
		return 0, err
	}

	nUpdated, err := service.updateDocument(tx, q, foundDocs)
	if err != nil {
		tx.Cancel()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	return int32(nUpdated), nil
}

func (service *Service) updateDocument(tx store.Transaction, q *mongo.Query, bsonDocs []bson.Document) (int32, error) {
	nUpdated := 0

	updateDocs := q.GetDocuments()
	for _, bsonDoc := range bsonDocs {
		updatedBSONDoc, err := UpdateBSONDocument(bsonDoc, updateDocs)
		if err != nil {
			return 0, err
		}
		objID, err := LookupBSONDocumentObjectID(updatedBSONDoc)
		if err != nil {
			return 0, err
		}
		updatedDoc, err := service.EncodeBSON(updatedBSONDoc)
		if err != nil {
			return 0, err
		}
		docKey := service.createDocumentKey(tx, q, objID)
		err = tx.UpdateDocument(docKey, updatedDoc)
		if err != nil {
			return 0, err
		}
		nUpdated++
	}

	return int32(nUpdated), nil
}

// Delete hadles OP_DELETE and 'delete' query of OP_MSG or OP_QUERY.
func (service *Service) Delete(q *mongo.Query) (int32, error) {
	db, err := service.GetDatabase(q.Database)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	tx, err := db.Transact(true)
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	nDeleted := 0

	queryConds := q.GetConditions()
	if len(queryConds) == 0 {
		nDeleted := len(service.documents)
		service.documents = make([]bson.Document, 0)
		return int32(nDeleted), nil
	}

	for n := (len(service.documents) - 1); 0 <= n; n-- {
		serverDoc := service.documents[n]
		isMatched := true
		for _, cond := range q.GetConditions() {
			condElems, err := cond.Elements()
			if err != nil {
				return 0, mongo.NewQueryError(q)
			}
			for _, condElem := range condElems {
				docValue, err := serverDoc.LookupErr(condElem.Key())
				if err != nil {
					isMatched = false
					break
				}
				condValue := condElem.Value()
				if !condValue.Equal(docValue) {
					isMatched = false
					break
				}
			}
		}

		if !isMatched {
			continue
		}

		service.documents = append(service.documents[:n], service.documents[n+1:]...)
		nDeleted++
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	return int32(nDeleted), nil
}
