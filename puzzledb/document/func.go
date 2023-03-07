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

func schemaMapFrom(obj any) (map[uint8]any, bool) {
	smap, ok := obj.(map[uint8]any)
	if ok {
		return smap, true
	}
	amap, ok := obj.(map[any]any)
	if !ok {
		return nil, false
	}
	smap = map[uint8]any{}
	for ak, av := range amap {
		switch k := ak.(type) {
		case int8:
			smap[uint8(k)] = av
		case uint8:
			smap[uint8(k)] = av
		default:
			return nil, false
		}
	}
	return smap, true
}

func schemaMapsFrom(obj any) ([]map[uint8]any, bool) {
	smaps, ok := obj.([]map[uint8]any)
	if ok {
		return smaps, true
	}
	amaps, ok := obj.([]any)
	if !ok {
		return nil, false
	}
	smaps = []map[uint8]any{}
	for _, amap := range amaps {
		smap, ok := schemaMapFrom(amap)
		if !ok {
			return nil, false
		}
		smaps = append(smaps, smap)
	}
	return smaps, true
}
