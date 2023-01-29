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

package document

import (
	"testing"

	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

func TestSchema(t *testing.T) {
	schema := NewSchema()

	if schema.Version() != SchemaVersion {
		t.Errorf("%v != %v", schema.Version(), SchemaVersion)
	}

	name := "s_name"
	schema.SetName(name)
	if schema.Name() != name {
		t.Errorf("%v != %v", schema.Name(), name)
	}

	name = "e_name"
	elem := NewElement()
	elem.SetName(name)
	if elem.Name() != name {
		t.Errorf("%v != %v", elem.Name(), name)
	}
	elem.SetType(store.Int)
	if elem.Type() != store.Int {
		t.Errorf("%v != %v", elem.Type(), store.Int)
	}

	name = "i_name"
	idx := NewIndex()
	idx.SetName(name)
	if idx.Name() != name {
		t.Errorf("%v != %v", idx.Name(), name)
	}
	idx.SetType(store.Primary)
	if idx.Type() != store.Primary {
		t.Errorf("%v != %v", idx.Type(), store.Primary)
	}
}
