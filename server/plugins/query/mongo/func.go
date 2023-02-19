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

	updatedDoc := bson.StartDocument()
	for _, docElem := range docElems {
		docKey := docElem.Key()
		docVal := docElem.Value()
		for _, updateDoc := range updateDocs {
			updateVal, err := updateDoc.LookupErr(docKey)
			if err == nil {
				docVal = updateVal
				break
			}
		}
		updatedDoc, err = bson.AppendValueElement(updatedDoc, docKey, docVal)
		if err != nil {
			return nil, err
		}
	}

	updatedDoc, err = bson.EndDocument(updatedDoc)
	if err != nil {
		return nil, err
	}

	err = updatedDoc.Validate()
	if err != nil {
		return nil, err
	}

	return updatedDoc, nil
}
