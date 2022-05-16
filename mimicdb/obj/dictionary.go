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

import "math"

const (
	// DictionaryMaxSize is the maximum size of Dictionary
	DictionaryMaxSize = math.MaxUint16
)

// Dictionary represents a dictionary instance.
type Dictionary map[string]Object

// NewDictionary returns a dictionary instance.
func NewDictionary() Dictionary {
	return Dictionary{}
}

// NewDictionaryWithBytes returns a dictionary instance with the specified bytes.
func NewDictionaryWithBytes(src []byte) (Dictionary, []byte, error) {
	return ReadDictionaryBytes(src)
}

// AddElement adds an element to the dictionary.
func (dict Dictionary) AddElement(key string, data Object) {
	dict[key] = data
}

// GetElementKeys returns all keys.
func (dict Dictionary) GetElementKeys() []string {
	keys := make([]string, 0, len(dict))
	for key := range dict {
		keys = append(keys, key)
	}
	return keys
}

// GetElementData returns a element of the specified key.
func (dict Dictionary) GetElementData(key string) (Object, bool) {
	data, ok := dict[key]
	return data, ok
}

// GetData returns the value.
func (dict Dictionary) GetData() interface{} {
	return dict
}

// Bytes returns the binary description.
func (dict Dictionary) Bytes() []byte {
	return AppendDictionaryBytes(nil, dict)
}

// Equals returns true when the specified dictionary is the same as this dictionary, otherwise false.
func (dict Dictionary) Equals(other Dictionary) bool {
	for key, val := range dict {
		otherVal, ok := other[key]
		if !ok {
			return false
		}
		if !val.Equals(otherVal) {
			return false
		}
	}
	return true
}

// AppendDictionaryBytes appends a value to the specified byte buffer.
func AppendDictionaryBytes(buf []byte, dict Dictionary) []byte {
	nMap := len(dict)
	if DictionaryMaxSize < nMap {
		nMap = DictionaryMaxSize
	}
	buf = AppendUint16Bytes(buf, uint16(nMap))
	for key, val := range dict {
		buf = AppendStringBytes(buf, key)
		buf = AppendDataBytes(buf, val)
	}
	return buf
}

// ReadDictionaryBytes reads the specified bytes as a dictionary.
func ReadDictionaryBytes(src []byte) (Dictionary, []byte, error) {
	var err error

	var nMap uint16
	nMap, src, err = ReadUint16Bytes(src)
	if err != nil {
		return nil, nil, err
	}

	dict := Dictionary{}

	var key string
	var val Object
	for n := 0; n < int(nMap); n++ {
		key, src, err = ReadStringBytes(src)
		if err != nil {
			return nil, nil, err
		}
		val, src, err = NewDataWithBytes(src)
		if err != nil {
			return nil, nil, err
		}
		dict[key] = val
	}

	return dict, src, nil
}
