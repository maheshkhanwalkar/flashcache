package server

import "sync"

// In-memory key-value store
type MemoryStore struct {
	store sync.Map
}

// Put (key, value) into the store
func (e *MemoryStore) Put(key string, value interface{}) {
	e.store.Store(key, value)
}

// Get the value associated within the given key
// Returns nil if no such key exists within the store
func (e *MemoryStore) Get(key string) interface{} {
	value, good := e.store.Load(key)

	if !good {
		return nil
	}

	return value
}

// Remove the (key, value) pair from the store
// Does nothing (no-op) if the key does not exist within the store
func (e *MemoryStore) Remove(key string) {
	e.store.Delete(key)
}
