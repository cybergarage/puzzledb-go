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

package serializer

import (
	"github.com/cybergarage/puzzledb-go/puzzledb/obj"
	"github.com/cybergarage/puzzledb-go/puzzledb/server/plugins"
)

// Serializer represents a serializer interface.
type Serializer interface {
	plugins.Service
	// Encode dumps a specified object to the byte array.
	Encode(obj obj.Object) ([]byte, error)
	// Decode creates an object from the specified byte array.
	Decode([]byte) (obj.Object, error)
}
