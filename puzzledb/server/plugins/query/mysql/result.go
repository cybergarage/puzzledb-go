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
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewResultFrom returns a successful result with the specified parameters.
func NewResultFrom(schema document.Schema, objs []document.Object) (*mysql.Result, error) {
	res := mysql.NewResult()

	resRows := [][]mysql.Value{}
	for _, obj := range objs {
		objMap, ok := obj.(map[string]any)
		if !ok {
			return nil, newObjectInvalidError(obj)
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
