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
	"errors"
	"fmt"
)

var (
	ErrExist        = errors.New("exist")
	ErrNotExist     = errors.New("not exist")
	ErrNotSupported = errors.New("not supported")
)

// Common error functions

func newNotSupportedError(target string) error {
	return fmt.Errorf("%v is %w", target, ErrNotSupported)
}

func newExistError(target string) error {
	return fmt.Errorf("%v is %w", target, ErrExist)
}

func newNotExistError(target string) error {
	return fmt.Errorf("%v is %w", target, ErrNotExist)
}

// Detail error functions

func newDatabaseExistError(target string) error {
	return newExistError(fmt.Sprintf("database (%s)", target))
}

func newTableExistError(target string) error {
	return newExistError(fmt.Sprintf("table (%s)", target))
}

func newSchemaExistError(target string) error {
	return newExistError(fmt.Sprintf("schema (%s)", target))
}

func newDatabaseNotExistError(target string) error {
	return newNotExistError(fmt.Sprintf("database (%s)", target))
}

func newTableNotExistError(target string) error {
	return newNotExistError(fmt.Sprintf("table (%s)", target))
}

func newSchemaNotExistError(target string) error {
	return newNotExistError(fmt.Sprintf("schema (%s)", target))
}

func newIndexNotSupportedError(target string) error {
	return newNotSupportedError(fmt.Sprintf("index (%s)", target))
}

func newQueryNotSupportedError(target string) error {
	return newNotSupportedError(fmt.Sprintf("query (%s)", target))
}
