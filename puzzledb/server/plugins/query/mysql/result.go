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

// NewResultWith returns a successful result with the specified parameters.
func NewResultWith(schema document.Schema, objs []document.Object) (*mysql.Result, error) {
	res := mysql.NewResult()

	resRows := [][]mysql.Value{}
	res.Rows = resRows

	return res, nil
}
