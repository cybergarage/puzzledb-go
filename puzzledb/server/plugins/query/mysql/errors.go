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

func newErrNotSupported(target string) error {
	return fmt.Errorf("%v is %w", target, ErrNotSupported)
}

func newErrExist(target string) error {
	return fmt.Errorf("%v is %w", target, ErrExist)
}

func newErrNotExist(target string) error {
	return fmt.Errorf("%v is %w", target, ErrNotExist)
}

// Detail error functions

func newErrDatabaseExist(target string) error {
	return newErrExist(fmt.Sprintf("database (%s)", target))
}

func newErrTableExist(target string) error {
	return newErrExist(fmt.Sprintf("table (%s)", target))
}

func newErrSchemaExist(target string) error {
	return newErrExist(fmt.Sprintf("schema (%s)", target))
}

func newErrDatabaseNotExist(target string) error {
	return newErrNotExist(fmt.Sprintf("database (%s)", target))
}

func newErrTableNotExist(target string) error {
	return newErrNotExist(fmt.Sprintf("table (%s)", target))
}

func newErrSchemaNotExist(target string) error {
	return newErrNotExist(fmt.Sprintf("schema (%s)", target))
}

func newErrIndexNotSupported(target string) error {
	return newErrNotSupported(fmt.Sprintf("index (%s)", target))
}

func newErrQueryNotSupported(target string) error {
	return newErrNotSupported(fmt.Sprintf("query (%s)", target))
}
