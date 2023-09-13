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
	"github.com/cybergarage/go-postgresql/postgresql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewRowFrom creates a new row from the specified object.
func NewRowFrom(obj document.Object) (query.Row, error) {
	objMap, err := document.NewMapObjectFrom(obj)
	if err != nil {
		return nil, err
	}
	return objMap, nil
}
