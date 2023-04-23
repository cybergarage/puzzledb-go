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
	"errors"
	"fmt"
)

var (
	ErrNotExist = errors.New("not exist")
	ErrExist    = errors.New("exist")
	ErrInvalid  = errors.New("invalid")
)

func NewDatabaseNotExistError(name string) error {
	return fmt.Errorf("database (%s) is %w", name, ErrNotExist)
}

func NewDatabaseExistError(name string) error {
	return fmt.Errorf("database (%s) is %w", name, ErrExist)
}

func NewSchemaNotExistError(name string) error {
	return fmt.Errorf("schema (%s) is %w ", name, ErrNotExist)
}

func NewObjectNotExistError(key Key) error {
	return fmt.Errorf("object (%s) is %w ", key, ErrNotExist)
}

func NewDatabaseOptionsInvalidError(opts any) error {
	return fmt.Errorf("database options (%v) is %w", opts, ErrInvalid)
}
