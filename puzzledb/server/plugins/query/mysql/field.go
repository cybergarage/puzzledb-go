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

// NewFieldFromSchema creates schema fields from the specified schema object.
func NewFieldFromSchema(dbName string, schema document.Schema) ([]*query.Field, error) {
	fields := make([]*query.Field, 0)
	/*
		tblName := schema.TableName()
		for _, column := range schema.DDL.GetTableSpec().Columns {
			colName := column.Name.String()
			// FIXME: Set more appreciate column length to check official MySQL implementation
			colLen := 65535 // len(name) + 1
			field := &Field{
				Database:     db.Name(),
				Table:        tblName,
				OrgTable:     tblName,
				Name:         colName,
				OrgName:      colName,
				ColumnLength: uint32(colLen),
				Charset:      255, // utf8mb4,
				Type:         column.Type.SQLType(),
			}

			switch field.Type {
			case vitesspq.Type_BLOB, vitesspq.Type_TEXT:
				field.Flags = field.Flags | uint32(vitesspq.MySqlFlag_BLOB_FLAG)
			}
			if column.Type.Options.Null != nil && !*column.Type.Options.Null {
				field.Flags = field.Flags | uint32(vitesspq.MySqlFlag_NOT_NULL_FLAG)
			}
			if column.Type.Options.KeyOpt == ColKeyPrimary {
				field.Flags = field.Flags | uint32(vitesspq.MySqlFlag_PRI_KEY_FLAG)
				field.Flags = field.Flags | uint32(vitesspq.MySqlFlag_NOT_NULL_FLAG)
			}

			fields = append(fields, field)
		}
	*/
	return fields, nil
}
