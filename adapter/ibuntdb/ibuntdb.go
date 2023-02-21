package ibuntdb

import (
	"fmt"
	"github.com/cute-angelia/go-utils/components/caches"
	cache "github.com/victorspringer/http-cache"
	"time"
)

// Adapter is the memory adapter data structure.
type Adapter struct {
	store caches.Cache
}

// NewAdapter initializes Redis adapter.
func NewAdapter(store caches.Cache) cache.Adapter {
	return &Adapter{
		store: store,
	}
}

func (a Adapter) Get(key uint64) ([]byte, bool) {
	v, error := a.store.Get(fmt.Sprintf("%d", key))
	return []byte(v), error == nil
}

func (a Adapter) Set(key uint64, response []byte, expiration time.Time) {
	a.store.Set(fmt.Sprintf("%d", key), string(response), expiration.Sub(time.Now()))
}

func (a Adapter) Release(key uint64) {
	a.store.Delete(fmt.Sprintf("%d", key))
}
