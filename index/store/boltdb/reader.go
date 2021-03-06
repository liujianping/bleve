//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package boltdb

import (
	"github.com/blevesearch/bleve/index/store"
	"github.com/boltdb/bolt"
)

type Reader struct {
	store *Store
	tx    *bolt.Tx
}

func newReader(store *Store) *Reader {
	tx, _ := store.db.Begin(false)
	return &Reader{
		store: store,
		tx:    tx,
	}
}

func (r *Reader) Get(key []byte) ([]byte, error) {
	rv := r.tx.Bucket([]byte(r.store.bucket)).Get(key)
	return rv, nil
}

func (r *Reader) Iterator(key []byte) store.KVIterator {
	rv := newIteratorExistingTx(r.store, r.tx)
	rv.Seek(key)
	return rv
}

func (r *Reader) Close() error {
	return r.tx.Rollback()
}
