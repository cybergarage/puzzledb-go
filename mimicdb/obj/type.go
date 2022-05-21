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

package obj

type Type uint8

const (
	DICTIONARY = Type(0x01)
	ARRAY      = Type(0x02)
	STRING     = Type(0x11)
	TINY       = Type(0x20)
	SHORT      = Type(0x21)
	INT        = Type(0x22)
	LONG       = Type(0x23)
	FLOAT      = Type(0x30)
	DOUBLE     = Type(0x31)
	TIMESTAMP  = Type(0x40)
	DATETIME   = Type(0x41)
	NULL       = Type(0x80)
	BOOL       = Type(0x81)
	BINARY     = Type(0x82)
)
