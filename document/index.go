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

type IndexType int

const (
	Primary   IndexType = 1
	Secondary IndexType = 2
)

type Index interface {
	// Name returns the unique name.
	Name() string
	// Type returns the index type.
	Type() IndexType
	// Elements returns the schema elements.
	Elements() []Element
	// SetName sets the specified name to the index.
	SetName(name string) Index
	// SetType sets the specified type to the element.
	SetType(t IndexType) Index
	// AddElement returns the schema elements.
	AddElement(elem Element) Index
	// Data returns the raw representation data in memory.
	Data() any
}
