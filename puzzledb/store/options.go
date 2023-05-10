// DocumentOperation represents a document operation. type DocumentOperation interface { // InsertDocument puts a document object with the specified primary key. InsertDocument(ctx context.Context, docKey Key, obj Object) error // FindDocuments returns a result set matching the specified key. FindDocuments(ctx context.Context, docKey Key) (ResultSet, error) // UpdateDocument updates a document object with the specified primary key. UpdateDocument(ctx context.Context, docKey Key, obj Object) error // RemoveDocument removes a document object with the specified primary key. RemoveDocument(ctx context.Context, docKey Key) error // RemoveDocuments removes document objects with the specified primary key. RemoveDocuments(ctx context.Context, docKey Key) error // TruncateDocuments removes all document objects. TruncateDocuments(ctx context.Context) error }// Copyright (C) 2022 The PuzzleDB Authors.
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

// Option represents a option.
type Option = any

// LimitOption represents a limit option.
type LimitOption struct {
	Limit int
}

// NewLimitOptionWith returns a new limit option.
func NewLimitOptionWith(limit int) *LimitOption {
	return &LimitOption{
		Limit: limit,
	}
}

// Order represents a order.
type Order = int

// NoLimit represents a no limit option.
var NoLimit = int(0)

// Order options.
const (
	OrderNone = Order(0)
	OrderAsc  = Order(1)
	OrderDesc = Order(2)
)

// OrderOption represents a order option.
type OrderOption struct {
	Order Order
}

var orderOptionNone = &OrderOption{
	Order: OrderNone,
}

var orderOptionAsc = &OrderOption{
	Order: OrderAsc,
}

var orderOptionDesc = &OrderOption{
	Order: OrderDesc,
}

// NewOrderOptionWith returns a new order option.
func NewOrderOptionWith(order Order) *OrderOption {
	switch order {
	case OrderAsc:
		return orderOptionAsc
	case OrderDesc:
		return orderOptionDesc
	default:
		return orderOptionNone
	}
}
