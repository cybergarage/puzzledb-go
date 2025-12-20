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

package kv

// Option represents a option.
type Option = any

// Offset represents an offset option.
type Offset uint

// NoOffset represents a no offset option.
var NoOffset = uint(0)

// NewLimit returns a new offset option.
func NewOffset(offset uint) Offset {
	return Offset(offset)
}

// Limit represents a limit option.
type Limit uint

// NoLimit represents a no limit option.
var NoLimit = uint(0)

// NewLimit returns a new limit option.
func NewLimit(limit int) Limit {
	return Limit(limit)
}

// Order represents a order.
type Order = int

// Order options.
const (
	OrderNone = Order(0)
	OrderAsc  = Order(1)
	OrderDesc = Order(2)
)
