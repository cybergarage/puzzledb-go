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

	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

var (
	ErrExist        = errors.New("exist")
	ErrNotExist     = errors.New("not exist")
	ErrNotSupported = errors.New("not supported")
	ErrInvalid      = errors.New("invalid")
	ErrNotEqual     = errors.New("not equal")
)

// Common error functions

func newNotSupportedError(obj any) error {
	return fmt.Errorf("%v is %w", obj, ErrNotSupported)
}

func newExistError(obj string) error {
	return fmt.Errorf("%v is %w", obj, ErrExist)
}

func newNotExistError(obj string) error {
	return fmt.Errorf("%v is %w", obj, ErrNotExist)
}

func newInvalidError(obj any) error {
	return fmt.Errorf("%v is %w", obj, ErrInvalid)
}

func newNotEqualError(obj1 any, obj2 any) error {
	return fmt.Errorf("%v and %v are %w", obj1, obj2, ErrNotEqual)
}

// Detail error functions

func newDatabaseExistError(obj string) error {
	return newExistError(fmt.Sprintf("database (%s)", obj))
}

func newTableExistError(obj string) error {
	return newExistError(fmt.Sprintf("table (%s)", obj))
}

func newSchemaExistError(obj string) error {
	return newExistError(fmt.Sprintf("schema (%s)", obj))
}

func newDatabaseNotExistError(obj string) error {
	return newNotExistError(fmt.Sprintf("database (%s)", obj))
}

func newTableNotExistError(obj string) error {
	return newNotExistError(fmt.Sprintf("table (%s)", obj))
}

func newSchemaNotExistError(obj string) error {
	return newNotExistError(fmt.Sprintf("schema (%s)", obj))
}

func newIndexNotSupportedError(obj string) error {
	return newNotSupportedError(fmt.Sprintf("index (%s)", obj))
}

func newIndexTypeNotSupportedError(t document.IndexType) error {
	return newNotSupportedError(fmt.Sprintf("index type (%02X)", t))
}

func newQueryNotSupportedError(obj string) error {
	return newNotSupportedError(fmt.Sprintf("query (%s)", obj))
}

func newPrimaryKeyDataNotExistError(keyName string, obj any) error {
	return newNotExistError(fmt.Sprintf("primary key data (%s:%v)", keyName, obj))
}

func newObjectInvalidError(obj any) error {
	return newInvalidError(fmt.Sprintf("object (%s:%v)", obj, obj))
}

// Not implemented error functions

func newJoinQueryNotSupportedError(obj any) error {
	return newNotSupportedError(fmt.Sprintf("JOIN query (%v)", obj))
}

func newQueryConditionNotSupportedError(obj any) error {
	return newNotSupportedError(fmt.Sprintf("query condition (%v)", obj))
}

func newDataTypeNotEqualError(obj any, et document.ElementType) error {
	return newNotEqualError(fmt.Sprintf("%v(%T)", obj, obj), et.String())
}
