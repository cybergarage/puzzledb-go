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
	"github.com/cybergarage/go-mongo/mongo/bson"
)

func UpdateBSONDocument(doc bson.Document, updateDocs []bson.Document) (bson.Document, error) {
	docElems, err := doc.Elements()
	if err != nil {
		return nil, err
	}

	updatedDoc := bson.DocumentStart()
	for _, docElem := range docElems {
		elemKey := docElem.Key()
		elemVal := docElem.Value()
		for _, updateDoc := range updateDocs {
			updateVal, err := updateDoc.LookupErr(elemKey)
			if err == nil {
				elemVal = updateVal
				break
			}
		}
		updatedDoc, err = bson.AppendValueElement(updatedDoc, elemKey, elemVal)
		if err != nil {
			return nil, err
		}
	}

	updatedDoc, err = bson.DocumentEnd(updatedDoc)
	if err != nil {
		return nil, err
	}

	err = updatedDoc.Validate()
	if err != nil {
		return nil, err
	}

	return updatedDoc, nil
}

func LookupBSONDocumentObjectID(bsonDoc bson.Document) (any, error) {
	// See : The _id Field - Documents (https://docs.mongodb.com/manual/core/document/)

	bsonObjID, err := bsonDoc.LookupErr(ObjectID)
	if err != nil {
		return nil, err
	}

	return EncodeBSONValue(bsonObjID)
}
