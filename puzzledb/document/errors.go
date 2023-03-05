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

var ErrNotSupported = errors.New("not supported")
var ErrSchemaInvalid = errors.New("invalid schema")

func newErrNotSupported(target string) error {
	return fmt.Errorf("%v is %w", target, ErrNotSupported)
}

func newErrSchemaInvalid(s any) error {
	return fmt.Errorf("%v(%T) is %w", s, s, ErrNotSupported)
}
