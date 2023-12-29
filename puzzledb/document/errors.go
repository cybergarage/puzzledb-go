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

package document

import (
	"errors"
	"fmt"
)

var (
	ErrNotSupported = errors.New("not supported")
	ErrInvalid      = errors.New("invalid")
	ErrNotExist     = errors.New("not exist")
	ErrNotFound     = errors.New("not found")
)

func newErrElementMapNotExist() error {
	return fmt.Errorf("element map is %w", ErrNotExist)
}

func newErrElementNotExistError(name string) error {
	return fmt.Errorf("element (%s) is %w", name, ErrNotExist)
}

func newErrIndexNotExist(name string) error {
	return fmt.Errorf("index (%s) is %w", name, ErrNotExist)
}

func newErrIndexMapNotExist() error {
	return fmt.Errorf("index map is %w", ErrNotExist)
}

func newErrElementInvalid(obj any) error {
	return fmt.Errorf("element (%T:%v) is %w", obj, obj, ErrInvalid)
}

func newErrSchemaInvalid(obj any) error {
	return fmt.Errorf("schema (%T:%v) is %w", obj, obj, ErrInvalid)
}

func newErrObjectInvalid(obj any) error {
	return fmt.Errorf("object (%T:%v) is %w", obj, obj, ErrInvalid)
}

func newErrIndexInvalid(obj any) error {
	return fmt.Errorf("index (%T:%v) is %w", obj, obj, ErrInvalid)
}

func newErrElementTypeInvalid(v any) error {
	return fmt.Errorf("element type (%s:%v) is %w", v, v, ErrInvalid)
}

func newErrDatabaseKeyNotFound(key Key) error {
	return fmt.Errorf("database ken (%s) is %w", key.String(), ErrNotFound)
}

func newErrCollectionKeyNotFound(key Key) error {
	return fmt.Errorf("collection ken (%s) is %w", key.String(), ErrNotFound)
}

// NewErrPrimaryIndexNotExist returns a new error that the primary index is not exist.
func NewErrPrimaryIndexNotExist() error {
	return fmt.Errorf("primary index is %w", ErrNotExist)
}

// NewErrObjectNotFound returns a new error that the object is not exist.
func NewErrObjectNotFound(key Key) error {
	return fmt.Errorf("object (%s) is %w ", key, ErrNotFound)
}
