// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

package nstree

import (
	"sync"

	"github.com/RaduBerinde/btree" // TODO(#144504): switch to the newer btree
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descpb"
)

func (t byIDMap) len() int {
	return t.t.Len()
}

type byIDItem struct {
	id descpb.ID
	v  interface{}
}

func makeByIDItem(d interface{ GetID() descpb.ID }) byIDItem {
	return byIDItem{id: d.GetID(), v: d}
}

var _ btree.Item = (*byIDItem)(nil)

func (b *byIDItem) Less(thanItem btree.Item) bool {
	than := thanItem.(*byIDItem)
	return b.id < than.id
}

var byIDItemPool = sync.Pool{
	New: func() interface{} { return new(byIDItem) },
}

func (b byIDItem) get() *byIDItem {
	alloc := byIDItemPool.Get().(*byIDItem)
	*alloc = b
	return alloc
}

func (b *byIDItem) value() interface{} {
	return b.v
}

func (b *byIDItem) put() {
	*b = byIDItem{}
	byIDItemPool.Put(b)
}
