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

const (
	errorUnknownType         = "Unknown data type : %X"
	errorInvalidTypeBytes    = "Invalid type bytes : %v"
	errorInvalidValueType    = "Invalid value : %v (%T)"
	errorInvalidArray        = "Invalid array : %v"
	errorInvalidDictionary   = "Invalid dictionary : %v"
	errorInvalidBooleanBytes = "Invalid boolean bytes : %v"
	errorInvalidNullBytes    = "Invalid null bytes : %v"
	errorInvalidStringBytes  = "Invalid string bytes : %d < %d %v"
	errorInvalidBinaryBytes  = "Invalid binary bytes : %d < %d %v"
	errorInvalidIntegerBytes = "Invalid integer bytes : %v"
)
