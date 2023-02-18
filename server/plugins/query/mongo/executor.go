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
)

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
		// See : The _id Field - Documents (https://docs.mongodb.com/manual/core/document/)
		queryObjID, err := queryDoc.LookupErr(ObjectID)
		if err != nil {
			continue
		}

		// Insert the document with the primary key

		storeDoc, err := service.Encode(queryDoc)
		if err != nil {
			return 0, err
		}

		storeObjID, err := EncodeBSON(queryObjID)
		if err != nil {
			return 0, err
		}

		storeKey := document.NewKeyWith(q.Database, q.Collection, ObjectID, storeObjID)
		err = tx.InsertDocument(storeKey, storeDoc)
		if err != nil {
			return 0, err
		}

		// Insert the secondary indexes for the all elements

		switch v := storeDoc.(type) {
		case map[string]any:
			for key, val := range v {
				indexKey := document.NewKeyWith(q.Database, q.Collection, key, val)
				err = tx.InsertIndex(indexKey, storeKey)
				if err != nil {
					return 0, err
				}
			}
		}

		nInserted++
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	if len(queryDocs) != int(nInserted) {
		return nInserted, mongo.NewQueryError(q)
	}

	return nInserted, nil
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

	foundDoc := make([]bson.Document, 0)

	for _, cond := range q.GetConditions() {
		condElems, err := cond.Elements()
		if err != nil {
			return nil, mongo.NewQueryError(q)
		}
		for _, condElem := range condElems {
			key := condElem.Key()
			bsonVal := condElem.Value()
			val, err := EncodeBSON(bsonVal)
			if err != nil {
				return nil, err
			}
			idxKey := document.NewKeyWith(q.Database, q.Collection, key, val)
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
				bsonDoc, err := DecodeBSON(obj)
				if err != nil {
					return nil, err
				}
				foundDoc = append(foundDoc, bsonDoc)
			}

		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	return foundDoc, nil
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

	nUpdated := 0

	queryDocs := q.GetDocuments()
	queryConds := q.GetConditions()
	if len(queryConds) == 0 {
		return 0, nil
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
				serverFoundElem, err := serverDoc.LookupErr(condElem.Key())
				if err != nil {
					isMatched = false
					break
				}
				if !condElem.Value().Equal(serverFoundElem) {
					isMatched = false
					break
				}
			}
		}

		if !isMatched {
			continue
		}

		service.documents = append(service.documents[:n], service.documents[n+1:]...)

		serverDocElems, err := serverDoc.Elements()
		if err != nil {
			return int32(nUpdated), mongo.NewQueryError(q)
		}

		updateDoc := bson.StartDocument()
		for _, serverDocElem := range serverDocElems {
			elemKey := serverDocElem.Key()
			elemValue := serverDocElem.Value()
			for _, queryDoc := range queryDocs {
				queryValue, err := queryDoc.LookupErr(elemKey)
				if err == nil {
					elemValue = queryValue
					break
				}
			}
			updateDoc, err = bson.AppendValueElement(updateDoc, elemKey, elemValue)
			if err != nil {
				return int32(nUpdated), mongo.NewQueryError(q)
			}
		}
		updateDoc, err = bson.EndDocument(updateDoc)
		if err != nil {
			return int32(nUpdated), mongo.NewQueryError(q)
		}

		err = updateDoc.Validate()
		if err != nil {
			return int32(nUpdated), mongo.NewQueryError(q)
		}

		service.documents = append(service.documents, updateDoc)

		nUpdated++
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
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
