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

package coordinator

// Option represents a option.
type Option = any

// OffsetOption represents an offset option.
type OffsetOption struct {
	Offset uint
}

// NoOffset represents a no offset option.
var NoOffset = uint(0)

// NewLimitOption returns a new offset option.
func NewOffsetOption(offset uint) *OffsetOption {
	return &OffsetOption{
		Offset: offset,
	}
}

// LimitOption represents a limit option.
type LimitOption struct {
	Limit int
}

// NoLimit represents a no limit option.
var NoLimit = int(0)

// NewLimitOption returns a new limit option.
func NewLimitOption(limit int) *LimitOption {
	return &LimitOption{
		Limit: limit,
	}
}

// Order represents a order.
type Order = int

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
