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

	docs := q.GetDocuments()
	for _, doc := range docs {
		// See : The _id Field - Documents (https://docs.mongodb.com/manual/core/document/)
		docValue, err := doc.LookupErr("_id")
		if err != nil {
			continue
		}

		isInserted := false

		for _, serverDoc := range service.documents {
			serverValue, err := serverDoc.LookupErr("_id")
			if err != nil {
				continue
			}
			if serverValue.Equal(docValue) {
				isInserted = true
				break
			}
		}

		if !isInserted {
			service.documents = append(service.documents, doc)
		}

		nInserted++
	}

	err = tx.Commit()
	if err != nil {
		return 0, mongo.NewQueryError(q)
	}

	if len(docs) != int(nInserted) {
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

	tx, err := db.Transact(true)
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	foundDoc := make([]bson.Document, 0)

	for _, doc := range service.documents {
		isMatched := true
		for _, cond := range q.GetConditions() {
			condElems, err := cond.Elements()
			if err != nil {
				return nil, mongo.NewQueryError(q)
			}
			for _, condElem := range condElems {
				docValue, err := doc.LookupErr(condElem.Key())
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

		foundDoc = append(foundDoc, doc)
	}

	err = tx.Commit()
	if err != nil {
		return nil, mongo.NewQueryError(q)
	}

	return foundDoc, nil
}

// Update hadles OP_UPDATE and 'update' query of OP_MSG or OP_QUERY.
func (server *Service) Update(q *mongo.Query) (int32, error) {
	nUpdated := 0

	queryDocs := q.GetDocuments()
	queryConds := q.GetConditions()
	if len(queryConds) == 0 {
		return 0, nil
	}

	for n := (len(server.documents) - 1); 0 <= n; n-- {
		serverDoc := server.documents[n]
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

		server.documents = append(server.documents[:n], server.documents[n+1:]...)

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

		server.documents = append(server.documents, updateDoc)

		nUpdated++
	}

	return int32(nUpdated), nil
}

// Delete hadles OP_DELETE and 'delete' query of OP_MSG or OP_QUERY.
func (server *Service) Delete(q *mongo.Query) (int32, error) {
	nDeleted := 0

	queryConds := q.GetConditions()
	if len(queryConds) == 0 {
		nDeleted := len(server.documents)
		server.documents = make([]bson.Document, 0)
		return int32(nDeleted), nil
	}

	for n := (len(server.documents) - 1); 0 <= n; n-- {
		serverDoc := server.documents[n]
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

		server.documents = append(server.documents[:n], server.documents[n+1:]...)
		nDeleted++
	}

	return int32(nDeleted), nil
}
