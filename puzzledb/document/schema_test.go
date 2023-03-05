// Copyright (C) 2020 The PuzzleDB Authors.
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

package document

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestSchema(t *testing.T) {
	now := time.Now()

	s1 := NewSchema()

	s1.SetName(now.String())

	for n, et := range elementTypes {
		e := NewElement().SetName(strconv.Itoa(n)).SetType(et)
		s1.AddElement(e)
	}

	for n, e := range s1.Elements() {
		idx := NewIndex().SetName(strconv.Itoa(n))
		if n == 0 {
			idx.SetType(PrimaryIndex)
		} else {
			idx.SetType(SecondaryIndex)
		}
		idx.AddElement(e)
		s1.AddIndex(idx)
	}

	// Compares

	s2, err := NewSchemaWith(s1.Data())
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(s2.Data(), s1.Data()) {
		t.Errorf("%v ! =%v", s2, s1)
	}

	if len(s1.Elements()) != len(s2.Elements()) {
		t.Errorf("%v ! =%v", s2, s1)
	}

	if len(s1.Indexes()) != len(s2.Indexes()) {
		t.Errorf("%v ! =%v", s2, s1)
	}
}
