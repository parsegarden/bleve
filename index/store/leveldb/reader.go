//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

// +build leveldb full

package leveldb

import (
	"github.com/blevesearch/bleve/index/store"
	"github.com/jmhodges/levigo"
)

type Reader struct {
	store    *Store
	snapshot *levigo.Snapshot
}

func newReader(store *Store) (*Reader, error) {
	return &Reader{
		store:    store,
		snapshot: store.db.NewSnapshot(),
	}, nil
}

func (r *Reader) BytesSafeAfterClose() bool {
	return true
}

func (r *Reader) Get(key []byte) ([]byte, error) {
	return r.store.getWithSnapshot(key, r.snapshot)
}

func (r *Reader) Iterator(key []byte) store.KVIterator {
	rv := newIteratorWithSnapshot(r.store, r.snapshot)
	rv.Seek(key)
	return rv
}

func (r *Reader) Close() error {
	r.store.db.ReleaseSnapshot(r.snapshot)
	return nil
}
