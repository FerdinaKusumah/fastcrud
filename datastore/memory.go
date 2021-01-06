package datastore

import (
	"github.com/tidwall/buntdb"
	"sync"
)

var (
	memSyncOnce sync.Once
	memDb       = new(buntdb.DB)
)

// NewMemoryDatastore initialize memory data store
func NewMemoryDatastore() *buntdb.DB {
	memSyncOnce.Do(func() {
		memDb, _ = buntdb.Open(":memory:")
	})
	return memDb
}

