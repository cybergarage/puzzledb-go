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

// MapObject represents a map object.
type MapObject map[string]any

// NewMapObjectFrom returns a new map object from the specified object.
func NewMapObjectFrom(anyObj any) (MapObject, error) {
	obj, ok := anyObj.(MapObject)
	if ok {
		return obj, nil
	}
	objMap, ok := anyObj.(map[any]any)
	if ok {
		obj := MapObject{}
		for key, val := range objMap {
			switch k := key.(type) {
			case string:
				obj[k] = val
			case []byte:
				obj[string(k)] = val
			default:
				return nil, newErrObjectInvalid(obj)
			}
		}
		return obj, nil
	}
	return nil, newErrObjectInvalid(obj)
}
