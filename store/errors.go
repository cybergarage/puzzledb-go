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

package store

import (
	"fmt"

	"github.com/cybergarage/puzzledb-go/puzzledb/errors"
)

var (
	ErrDatabaseNotFound = errors.New("database not found")
	ErrSchemaNotFound   = errors.New("schema not found")
	ErrObjectNotFound   = errors.New("object not found")
)

func NewDatabaseNotFoundError(name string) error {
	return fmt.Errorf("%w (%s)", ErrDatabaseNotFound, name)
}

func NewSchemaNotFoundError(name string) error {
	return fmt.Errorf("%w (%s)", ErrSchemaNotFound, name)
}
