package ttlstore

import (
	"sync"
	"time"
)

type entry struct {
	value      string
	expiration time.Time
}

type TTLStore struct {
	data map[string]entry
	mu   sync.Mutex
	done chan struct{}
}

func NewStore() *TTLStore {
	store := &TTLStore{
		data: make(map[string]entry),
		mu:   sync.Mutex{},
	}

	go func() {
		for k, e := range store.data {
			if e.expiration.After(time.Now()) {
				delete(store.data, k)
			}
		}
	}()

	return store
}
