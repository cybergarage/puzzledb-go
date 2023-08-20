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
	"time"

	"github.com/cybergarage/go-mongo/mongo/bson"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// BSONEncoder represents a BSON encoder.
type BSONEncoder struct {
}

// NewBSONEncoder returns a new BSON encoder instance.
func NewBSONEncoder() *BSONEncoder {
	return &BSONEncoder{}
}

// EncodeBSON encodes the specified BSON object to a document object.
func (s *BSONEncoder) EncodeBSON(bsonDoc bson.Document) (document.Object, error) {
	return EncodeBSONDocument(bsonDoc)
}

func EncodeBSONDocument(bsonDoc bson.Document) (document.Object, error) {
	obj := map[string]any{}
	bsonElems, err := bsonDoc.Elements()
	if err != nil {
		return nil, err
	}

	for _, bsonElem := range bsonElems {
		key := bsonElem.Key()
		bsonVal, err := bsonElem.ValueErr()
		if err != nil {
			return nil, err
		}
		data, err := EncodeBSONValue(bsonVal)
		if err != nil {
			return nil, err
		}
		obj[key] = data
	}
	return obj, nil
}

func EncodeBSONValue(bsonVal bsoncore.Value) (any, error) {
	/* TODO: The following BSON types are not supported yet.
	   Undefined        Type = 0x06
	   DBPointer        Type = 0x0C
	   JavaScript       Type = 0x0D
	   Symbol           Type = 0x0E
	   CodeWithScope    Type = 0x0F
	   Timestamp        Type = 0x11
	   Decimal128       Type = 0x13
	   MinKey           Type = 0xFF
	   MaxKey           Type = 0x7F
	*/

	switch bsonVal.Type { //nolint:all
	case bsontype.Array:
		bsonArray := bsonVal.Array()
		bsonElems, err := bsonArray.Values()
		if err != nil {
			return nil, err
		}
		array := []any{}
		for _, bsonElem := range bsonElems {
			obj, err := EncodeBSONValue(bsonElem)
			if err != nil {
				return nil, err
			}
			array = append(array, obj)
		}
		return array, nil
	case bsontype.EmbeddedDocument:
		bsonDoc := bsonVal.Document()
		bsonElems, err := bsonDoc.Elements()
		if err != nil {
			return nil, err
		}
		dict := map[string]any{}
		for _, bsonElem := range bsonElems {
			key := bsonElem.Key()
			obj, err := EncodeBSONValue(bsonElem.Value())
			if err != nil {
				return nil, err
			}
			dict[key] = obj
		}
		return dict, nil
	case bsontype.Double:
		return bsonVal.Double(), nil
	case bsontype.String:
		return bsonVal.StringValue(), nil
	case bsontype.Binary:
		_, bytes := bsonVal.Binary()
		return bytes, nil
	case bsontype.Boolean:
		return bsonVal.Boolean(), nil
	case bsontype.DateTime:
		ts := bsonVal.DateTime()
		return time.Unix(ts, 0), nil
	case bsontype.Null:
		return nil, nil // nolint: nilnil
	case bsontype.Int32:
		return bsonVal.Int32(), nil
	case bsontype.Int64:
		return bsonVal.Int64(), nil
	case bsontype.ObjectID:
		return bsonVal.ObjectID().Hex(), nil
	}
	return nil, newErrBSONTypeNotSupported(bsonVal.Type)
}
