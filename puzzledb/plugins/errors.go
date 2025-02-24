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

package plugins

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNotFound = errors.New("not found")
var ErrInvalid = errors.New("invalid")
var ErrDisabled = errors.New("disabled")

func NewErrNotFound(target string) error {
	return fmt.Errorf("%v is %w", target, ErrNotFound)
}

func NewErrInvalid(target string) error {
	return fmt.Errorf("%v is %w", target, ErrInvalid)
}

func NewErrDisabled(target string) error {
	return fmt.Errorf("%v is %w", target, ErrDisabled)
}

func NewErrServiceNotFound(t ServiceType) error {
	return NewErrInvalid(t.String())
}

func NewErrCounfigNotFound(s any) error {
	switch v := s.(type) {
	case []string:
		return NewErrNotFound(fmt.Sprintf("config (%s)", strings.Join(v, ".")))
	default:
		return NewErrNotFound(fmt.Sprintf("config (%v)", v))
	}
}

func NewErrDefaultServiceNotFound(t ServiceType) error {
	return NewErrNotFound(fmt.Sprintf("default %s service", t.String()))
}

func NewErrInvalidService(s Service) error {
	return NewErrInvalid(fmt.Sprintf("%s (%s)", s.ServiceName(), s.ServiceType().String()))
}

func NewErrDisabledService(s Service) error {
	return NewErrDisabled(fmt.Sprintf("%s (%s)", s.ServiceName(), s.ServiceType().String()))
}
