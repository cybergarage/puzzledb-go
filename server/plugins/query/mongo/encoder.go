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
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// Encoder represents a BSON encoder.
type Encoder struct {
}

// NewEncoder returns a new BSON encoder instance.
func NewEncoder() *Encoder {
	return &Encoder{}

}

// Encode encodes the specified BSON object to a document object.
func (s *Encoder) Encode(obj bson.Document) (document.Object, error) {
	return nil, nil
}
