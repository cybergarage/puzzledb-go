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
	"github.com/cybergarage/go-mysql/mysql"
	"github.com/cybergarage/go-mysql/mysql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewResultFrom returns a successful result with the specified parameters.
func NewResultFrom(schema document.Schema, objs []document.Object) (*mysql.Result, error) {
	res := mysql.NewResult()

	resRows := [][]mysql.Value{}
	for _, obj := range objs {
		objMap, err := NewObjectWith(obj)
		if err != nil {
			return nil, err
		}
		resValues := []mysql.Value{}
		for colName, colVal := range objMap {
			colElem, err := schema.FindElement(colName)
			if err != nil {
				return nil, err
			}
			resValue, err := NewValueFrom(colElem, colVal)
			if err != nil {
				return nil, err
			}
			resValues = append(resValues, resValue)
		}
		resRows = append(resRows, resValues)
	}

	res.Rows = resRows

	return res, nil
}

// NewFieldFromSchema creates schema fields from the specified schema object.
func NewFieldFromSchema(dbName string, schema document.Schema) ([]*query.Field, error) {
	fields := make([]*query.Field, 0)
	tblName := schema.Name()
	for _, elem := range schema.Elements() {
		colName := elem.Name()
		colType, err := sqlTypeFromElementType(elem.Type())
		if err != nil {
			return nil, err
		}
		// FIXME: Set more appreciate column length to check official MySQL implementation
		colLen := 65535 // len(name) + 1
		field := &query.Field{
			// name of the field as returned by mysql C API
			Name: colName,
			// Remaining fields from mysql C API.
			Database: dbName,
			Table:    tblName,
			OrgTable: tblName,
			OrgName:  colName,
			// column_length is really a uint32. All 32 bits can be used.
			ColumnLength: uint32(colLen),
			// charset is actually a uint16. Only the lower 16 bits are used.
			Charset: 255, // utf8mb4,
			// vitess-defined type. Conversion function is in sqltypes package.
			Type: colType,
			// decimals is actually a uint8. Only the lower 8 bits are used.
			Decimals: 0,
			// flags is actually a uint16. Only the lower 16 bits are used.
			Flags: 0,
			// column_type is optionally populated from information_schema.columns
			ColumnType: "",
		}
		/*
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
		*/
		fields = append(fields, field)
	}
	return fields, nil
}
