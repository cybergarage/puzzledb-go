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

package vbde

import (
	"github.com/cybergarage/mimicdb/mimicdb/query"
)

// Executor represents a virtual machine executor.
type Executor struct {
	query.Executor
}

// Execute execute the specified compiled query object.
func (m *Executor) Execute(ctx *query.DBContext, estmt query.Statement) error {
	_, ok := estmt.(*Statement)
	if !ok {
		return nil
	}
	return nil
}
