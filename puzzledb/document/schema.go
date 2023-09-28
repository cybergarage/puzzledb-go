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

// Schema represents a schema.
type Schema interface {
	// Version returns the schema version.
	Version() int
	// SetName sets the specified name to the schema.
	SetName(name string)
	// Name returns the schema name.
	Name() string
	// AddElement adds the specified element to the schema.
	AddElement(elem Element)
	// DropElement drops the specified element from the schema.
	DropElement(name string) error
	// Elements returns the schema elements.
	Elements() Elements
	// FindElement returns the schema elements by the specified name.
	FindElement(name string) (Element, error)
	// AddIndex adds the specified index to the schema.
	AddIndex(idx Index)
	// DropIndex drops the specified index from the schema.
	DropIndex(name string) error
	// Indexes returns the schema indexes.
	Indexes() Indexes
	// FindIndex returns the schema index by the spacified name.
	FindIndex(name string) (Index, error)
	// PrimaryIndex returns the schema primary index.
	PrimaryIndex() (Index, error)
	// SecondaryIndexes returns the schema secondary indexes.
	SecondaryIndexes() (Indexes, error)
	// Data returns the raw representation data in memory.
	Data() any
}
