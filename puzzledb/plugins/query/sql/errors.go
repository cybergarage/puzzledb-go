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
	ErrNotFound     = errors.New("not found")
)

// Common error functions

func newErrNotSupported(obj any) error {
	return fmt.Errorf("%v is %w", obj, ErrNotSupported)
}

func newErrExist(obj string) error {
	return fmt.Errorf("%v is %w", obj, ErrExist)
}

func newErrNotExist(obj string) error {
	return fmt.Errorf("%v is %w", obj, ErrNotExist)
}

func newErrInvalid(obj any) error {
	return fmt.Errorf("%v is %w", obj, ErrInvalid)
}

func newErrNotEqual(obj1 any, obj2 any) error {
	return fmt.Errorf("%v and %v are %w", obj1, obj2, ErrNotEqual)
}

// Detail error functions

func newErrDatabaseExist(obj string) error {
	return newErrExist(fmt.Sprintf("database (%s)", obj))
}

func newErrTableExist(obj string) error {
	return newErrExist(fmt.Sprintf("table (%s)", obj))
}

func newErrSchemaExist(obj string) error {
	return newErrExist(fmt.Sprintf("schema (%s)", obj))
}

func newErrDatabaseNotExist(obj string) error {
	return newErrNotExist(fmt.Sprintf("database (%s)", obj))
}

func newErrTableNotExist(obj string) error {
	return newErrNotExist(fmt.Sprintf("table (%s)", obj))
}

func newErrSchemaNotExist(obj string) error {
	return newErrNotExist(fmt.Sprintf("schema (%s)", obj))
}

func newErrIndexNotSupported(obj string) error {
	return newErrNotSupported(fmt.Sprintf("index (%s)", obj))
}

func newErrIndexTypeNotSupported(t document.IndexType) error {
	return newErrNotSupported(fmt.Sprintf("index type (%02X)", t))
}

func newErrQueryNotSupported(obj string) error {
	return newErrNotSupported(fmt.Sprintf("query (%s)", obj))
}

func newErrPrimaryKeyDataNotExist(keyName string, obj any) error {
	return newErrNotExist(fmt.Sprintf("primary key data (%s:%v)", keyName, obj))
}

func newErrObjectInvalid(obj any) error {
	return newErrInvalid(fmt.Sprintf("object (%s:%v)", obj, obj))
}

func newErrCoulumNotExist(obj any) error {
	return newErrNotExist(fmt.Sprintf("coulum (%s)", obj))
}

func newErrConnectionExist(obj string) error {
	return newErrExist(fmt.Sprintf("connection (%s)", obj))
}

// Not implemented error functions

func newErrJoinQueryNotSupported(obj any) error {
	return newErrNotSupported(fmt.Sprintf("JOIN query (%v)", obj))
}

func newErrQueryConditionNotSupported(obj any) error {
	return newErrNotSupported(fmt.Sprintf("query condition (%v)", obj))
}

func newErrDataTypeNotEqual(obj any, et document.ElementType) error {
	return newErrNotEqual(fmt.Sprintf("%v(%T)", obj, obj), et.String())
}
