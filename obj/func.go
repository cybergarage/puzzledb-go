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

import (
	"fmt"
)

// ReadInt8Bytes reads the specified bytes as a byte integer.
func ReadInt8Bytes(src []byte) (int8, []byte, error) {
	srcLen := len(src)
	if srcLen < 1 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return int8(src[0]), src[1:], nil
}

// ReadUint8Bytes reads the specified bytes as a byte integer.
func ReadUint8Bytes(src []byte) (uint8, []byte, error) {
	srcLen := len(src)
	if srcLen < 1 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return uint8(src[0]), src[1:], nil
}

// ReadUint16Bytes reads the specified bytes as a integer.
func ReadUint16Bytes(src []byte) (uint16, []byte, error) {
	srcLen := len(src)
	if srcLen < 2 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (uint16(src[0])<<8 | uint16(src[1])), src[2:], nil
}

// ReadUint32Bytes reads the specified bytes as a integer.
func ReadUint32Bytes(src []byte) (uint32, []byte, error) {
	srcLen := len(src)
	if srcLen < 4 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])), src[4:], nil
}

// ReadUint64Bytes reads the specified bytes as a long integer.
func ReadUint64Bytes(src []byte) (uint64, []byte, error) {
	srcLen := len(src)
	if srcLen < 8 {
		return 0, nil, fmt.Errorf(errorInvalidIntegerBytes, src)
	}
	return (uint64(src[0])<<56 | uint64(src[1])<<48 | uint64(src[2])<<40 | uint64(src[3])<<32 | uint64(src[4])<<24 | uint64(src[5])<<16 | uint64(src[6])<<8 | uint64(src[7])), src[8:], nil
}

// ReadTypeBytes reads the specified bytes as a type value.
func ReadTypeBytes(src []byte) (Type, []byte, error) {
	srcLen := len(src)
	if srcLen < 1 {
		return Type(0), nil, fmt.Errorf(errorInvalidTypeBytes, src)
	}
	return Type(src[0]), src[1:], nil
}

// AppendInt8Bytes appends a value to the specified buffer.
func AppendInt8Bytes(buf []byte, val int8) []byte {
	return append(buf,
		byte(val),
	)
}

// AppendUint8Bytes appends a value to the specified buffer.
func AppendUint8Bytes(buf []byte, val uint8) []byte {
	return append(buf,
		byte(val),
	)
}

// AppendUint16Bytes appends a value to the specified buffer.
func AppendUint16Bytes(buf []byte, val uint16) []byte {
	return append(buf,
		byte(val>>8),
		byte(val),
	)
}

// AppendUint32Bytes appends a value to the specified buffer.
func AppendUint32Bytes(buf []byte, val uint32) []byte {
	return append(buf,
		byte(val>>24),
		byte(val>>16),
		byte(val>>8),
		byte(val),
	)
}

// AppendUint64Bytes appends a value to the specified buffer.
func AppendUint64Bytes(buf []byte, val uint64) []byte {
	return append(buf,
		byte(val>>56),
		byte(val>>48),
		byte(val>>40),
		byte(val>>32),
		byte(val>>24),
		byte(val>>16),
		byte(val>>8),
		byte(val),
	)
}

// AppendTypeBytes appends a type value to the specified buffer.
func AppendTypeBytes(buf []byte, val Type) []byte {
	return append(buf, byte(val))
}
