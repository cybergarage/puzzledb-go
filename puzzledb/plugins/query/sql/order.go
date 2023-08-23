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

package sql

import (
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/store"
)

// NewOrderWith returns a new store option with the specified orderby option.
func NewOrderWith(orderBy *query.OrderBy) []store.Option {
	// TODO: Convert query.OrderBy to store.Option
	// https://github.com/vitessio/vitess/blob/v0.12.6/go/vt/sqlparser/ast.go
	// // Order represents an ordering expression.
	// type Order struct {
	// 	Expr      Exprd
	// 	Direction OrderDirection
	// }
	// // OrderDirection is an enum for the direction in which to order - asc or desc.
	// type OrderDirection int8
	opts := []store.Option{}
	return opts
}
