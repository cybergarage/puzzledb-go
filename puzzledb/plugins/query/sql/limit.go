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

// NewLimitWith returns a new store option with the specified limit option.
func NewLimitWith(limit *query.Limit) []store.Option {
	// TODO: Convert query.Limit to store.Option
	// https://github.com/vitessio/vitess/blob/v0.12.6/go/vt/sqlparser/ast.go
	// // Limit represents a LIMIT clause.
	// type Limit struct {
	// 	Offset, Rowcount Expr
	// }
	opts := []store.Option{}
	return opts
}
