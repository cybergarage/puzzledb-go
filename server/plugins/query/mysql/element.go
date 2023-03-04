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

package mysql

import (
	"github.com/cybergarage/go-mysql/mysql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewIndexWith creates an index from the specified index object.
func NewElementWith(col *query.ColumnDefinition) (document.Element, error) {
	e := document.NewElement()
	e.SetName(col.Name.String())
	switch col.Type.SQLType() {
	case query.Int8:
		e.SetType(document.Int8)
	case query.Int16:
		e.SetType(document.Int16)
	case query.Int32:
		e.SetType(document.Int32)
	case query.Int64:
		e.SetType(document.Int64)
	case query.Float32:
		e.SetType(document.Float32)
	case query.Float64:
		e.SetType(document.Float64)
	case query.Text, query.VarChar:
		e.SetType(document.String)
	case query.Blob:
		e.SetType(document.Binary)
	default:
		return nil, newErrNotSupported(col.Type.SQLType().String())
	}
	return e, nil
}
