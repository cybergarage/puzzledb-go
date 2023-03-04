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

type ElementType int

const (
	// 0x0x (Reserved).
	// 0x1x (Reserved - Collection).
	Array ElementType = 0x10
	Map   ElementType = 0x11
	// 0x20 Integer.
	Int8  ElementType = 0x20
	Int16 ElementType = 0x21
	Int32 ElementType = 0x22
	Int64 ElementType = 0x23
	// 0x30 String.
	String ElementType = 0x30
	Blob   ElementType = 0x31
	// 0x40 Floating-point.
	Float  ElementType = 0x40
	Double ElementType = 0x41
	// 0x70 Special.
	DateTime ElementType = 0x70
	Bool     ElementType = 0x71
)

type Element interface {
	// Name returns the unique name.
	Name() string
	// Type returns the index type.
	Type() ElementType
	// Data returns the raw representation data in memory.
	Data() any
	// SetName sets the specified name to the element.
	SetName(name string) Element
	// SetType sets the specified type to the element.
	SetType(t ElementType) Element
}
