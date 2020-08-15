package server

type ExecutionEngine struct {
	store map[string]interface{}
}

// Put (key, value) into the store
func (e *ExecutionEngine) Put(key string, value interface{}) {
	e.store[key] = value
}

// Get the value associated within the given key
// Returns nil if no such key exists within the store
func (e *ExecutionEngine) Get(key string) interface{} {
	value, good := e.store[key]

	if !good {
		return nil
	}

	return value
}

// Remove the (key, value) pair from the store
// Does nothing (no-op) if the key does not exist within the store
func (e *ExecutionEngine) Remove(key string) {
	delete(e.store, key)
}
