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

package store

type IndexType int
type ElementType int

const (
	Primary IndexType = 0
)

const (
	Int ElementType = 0
)

type Index interface {
	// Name returns the unique name.
	Name() string
	// Type returns the index type.
	Type() IndexType
	// Elements returns the schema elements.
	Elements() []Element
}

type Element interface {
	// Name returns the unique name.
	Name() string
	// Type returns the index type.
	Type() ElementType
}

type Schema interface {
	// Version returns the schema version.
	Version() int
	// SetName sets the specified name to the schema.
	SetName(name string)
	// Name returns the schema name.
	Name() string
	// AddElement adds the specified element to the schema.
	AddElement(elem Element)
	// Elements returns the schema elements.
	Elements() []Element
	// AddIndex adds the specified index to the schema.
	AddIndex(idx Index)
	// Elements returns the schema elements.
	Indexes() []Index
}
