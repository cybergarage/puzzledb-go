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
	"bytes"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func TestCoder(t *testing.T) {
	now := time.Unix(time.Now().Unix(), 0)
	bsonObj := bson.D{
		{Key: "string", Value: "abc"},
		{Key: "binary", Value: []byte("abc")},
		{Key: "int32", Value: int32(1)},
		{Key: "int64", Value: int64(1)},
		{Key: "double", Value: float64(1)},
		{Key: "time", Value: now},
		{Key: "bool", Value: true},
		{Key: "null", Value: nil},
	}

	bsonBytes, err := bson.Marshal(bsonObj)
	if err != nil {
		t.Error(err)
		return
	}

	bsonDoc, err := bsoncore.NewDocumentFromReader(bytes.NewReader(bsonBytes))
	if err != nil {
		t.Error(err)
		return
	}

	s := NewCoder()

	// BSON -> Go
	goObj, err := s.EncodeBSON(bsonDoc)
	if err != nil {
		t.Error(err)
		return
	}

	// Go -> BSON
	bsonDoc, err = s.DecodeBSON(goObj)
	if err != nil {
		t.Error(err)
		return
	}

	// BSON -> Go
	newGoObj, err := s.EncodeBSON(bsonDoc)
	if err != nil {
		t.Error(err)
		return
	}

	// Compares
	if !reflect.DeepEqual(newGoObj, goObj) {
		t.Errorf("%v != %v", newGoObj, goObj)
	}
}
