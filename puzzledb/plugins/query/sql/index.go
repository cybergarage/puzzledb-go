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

package sql

import (
	"github.com/cybergarage/go-sqlparser/sql/query"
	"github.com/cybergarage/puzzledb-go/puzzledb/document"
)

// NewDocumentIndexTypeFrom creates an index type from the specified element.
func NewDocumentIndexTypeFrom(idxType query.IndexType) (document.IndexType, error) {
	switch idxType {
	case query.PrimaryIndex:
		return document.PrimaryIndex, nil
	case query.SecondaryIndex:
		return document.SecondaryIndex, nil
	case query.UnknownIndex:
		return 0, newErrIndexNotSupported(idxType.String())
	}
	return 0, newErrIndexNotSupported(idxType.String())
}

// NewDocumentPrimaryIndexFrom creates an index from the specified element.
func NewDocumentPrimaryIndexFrom(elem document.Element) (document.Index, error) {
	idx := document.NewIndex()
	idx.SetName(elem.Name())
	idx.SetType(document.PrimaryIndex)
	idx.AddElement(elem)
	return idx, nil
}

// NewDocumentIndexFrom creates an index from the specified coulumn definition.
func NewDocumentIndexFrom(s document.Schema, def query.Index) (document.Index, error) {
	idxType, err := NewDocumentIndexTypeFrom(def.Type())
	if err != nil {
		return nil, err
	}

	idx := document.NewIndex()
	idx.SetName(def.Name())
	idx.SetType(idxType)

	for _, col := range def.Columns() {
		elem, err := s.FindElement(col.Name())
		if err != nil {
			return nil, err
		}
		idx.AddElement(elem)
	}

	return idx, nil
}
