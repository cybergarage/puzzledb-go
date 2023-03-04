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
	"fmt"
	"strconv"
	"time"

	"github.com/cybergarage/go-mongo/mongo/bson"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// BSONDecoder represents a decorder.
type BSONDecoder struct {
}

// NewBSONDecoder returns a new CBOR erializer instance.
func NewBSONDecoder() *BSONDecoder {
	return &BSONDecoder{}

}

// DecodeBSON returns the decorded BSON object from the specified object.
func (s *BSONDecoder) DecodeBSON(obj document.Object) (bson.Document, error) {
	return DecodeBSONDocument(obj)
}

// DecodeBSONDocument returns the decorded BSON object from the specified object.
func DecodeBSONDocument(obj document.Object) (bson.Document, error) {
	idx, bsonDoc := bsoncore.AppendDocumentStart(nil)
	var err error
	switch v := obj.(type) {
	case map[string]any:
		for key, val := range v {
			bsonDoc, err = bsonDocumentAddObject(bsonDoc, key, val)
			if err != nil {
				return nil, err
			}
		}
	case map[any]any:
		for key, val := range v {
			switch vkey := key.(type) {
			case string:
				bsonDoc, err = bsonDocumentAddObject(bsonDoc, vkey, val)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	bsonDoc, err = bsoncore.AppendDocumentEnd(bsonDoc, idx)
	if err != nil {
		return nil, err
	}
	return bsonDoc, nil
}

func DecodeBSONValue(obj any) (*bsoncore.Value, error) {
	var err error
	switch v := obj.(type) {
	case map[string]any:
		idx, bsonDoc := bsoncore.AppendDocumentStart(nil)
		for key, val := range v {
			bsonDoc, err = bsonDocumentAddObject(bsonDoc, key, val)
			if err != nil {
				return nil, err
			}
		}
		bsonDoc, err = bsoncore.AppendDocumentEnd(bsonDoc, idx)
		if err != nil {
			return nil, err
		}
		bsonVal := &bsoncore.Value{
			Type: bsontype.EmbeddedDocument,
			Data: bsonDoc,
		}
		return bsonVal, nil
	case []any:
		idx, bsonDoc := bsoncore.AppendArrayStart(nil)
		for n, val := range v {
			key := strconv.Itoa(n)
			bsonDoc, err = bsonDocumentAddObject(bsonDoc, key, val)
			if err != nil {
				return nil, err
			}
		}
		bsonDoc, err = bsoncore.AppendArrayEnd(bsonDoc, idx)
		if err != nil {
			return nil, err
		}
		bsonVal := &bsoncore.Value{
			Type: bsontype.Array,
			Data: bsonDoc,
		}
		return bsonVal, nil
	case string:
		bsonVal := &bsoncore.Value{
			Type: bsontype.String,
			Data: bsoncore.AppendString(nil, v),
		}
		return bsonVal, nil
	case int:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Int32,
			Data: bsoncore.AppendInt32(nil, int32(v)),
		}
		return bsonVal, nil
	case int32:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Int32,
			Data: bsoncore.AppendInt32(nil, v),
		}
		return bsonVal, nil
	case int64:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Int64,
			Data: bsoncore.AppendInt64(nil, v),
		}
		return bsonVal, nil
	case float32:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Double,
			Data: bsoncore.AppendDouble(nil, float64(v)),
		}
		return bsonVal, nil
	case float64:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Double,
			Data: bsoncore.AppendDouble(nil, v),
		}
		return bsonVal, nil
	case time.Time:
		bsonVal := &bsoncore.Value{
			Type: bsontype.DateTime,
			Data: bsoncore.AppendDateTime(nil, v.Unix()),
		}
		return bsonVal, nil
	case nil:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Null,
			Data: make([]byte, 0),
		}
		return bsonVal, nil
	case bool:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Boolean,
			Data: bsoncore.AppendBoolean(nil, v),
		}
		return bsonVal, nil
	case []byte:
		bsonVal := &bsoncore.Value{
			Type: bsontype.Binary,
			Data: bsoncore.AppendBinary(nil, 0x00, v),
		}
		return bsonVal, nil
	}

	return nil, fmt.Errorf("unknown object type : %T", obj)
}

func bsonDocumentAddObject(bsonDoc []byte, key string, obj any) ([]byte, error) {
	bsonVal, err := DecodeBSONValue(obj)
	if err != nil {
		return nil, err
	}

	switch bsonVal.Type {
	case bsontype.Array:
		return bsoncore.AppendArrayElement(bsonDoc, key, bsonVal.Array()), nil
	case bsontype.Boolean:
		return bsoncore.AppendBooleanElement(bsonDoc, key, bsonVal.Boolean()), nil
	case bsontype.Int32:
		return bsoncore.AppendInt32Element(bsonDoc, key, bsonVal.Int32()), nil
	case bsontype.Int64:
		return bsoncore.AppendInt64Element(bsonDoc, key, bsonVal.Int64()), nil
	case bsontype.Double:
		return bsoncore.AppendDoubleElement(bsonDoc, key, bsonVal.Double()), nil
	case bsontype.String:
		return bsoncore.AppendStringElement(bsonDoc, key, bsonVal.StringValue()), nil
	case bsontype.EmbeddedDocument:
		return bsoncore.AppendDocumentElement(bsonDoc, key, bsonVal.Document()), nil
	case bsontype.ObjectID:
		return bsoncore.AppendObjectIDElement(bsonDoc, key, bsonVal.ObjectID()), nil
	case bsontype.DateTime:
		return bsoncore.AppendDateTimeElement(bsonDoc, key, bsonVal.DateTime()), nil
	case bsontype.Binary:
		subType, binData := bsonVal.Binary()
		return bsoncore.AppendBinaryElement(bsonDoc, key, subType, binData), nil
	case bsontype.Null:
		return bsoncore.AppendNullElement(bsonDoc, key), nil
	}

	return bsonDoc, fmt.Errorf("unknown element type : %v", bsonVal)
}
