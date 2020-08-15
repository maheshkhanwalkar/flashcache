package server

import "flashcache/protocol"

type ExecutionEngine struct {
	mem *MemoryStore
}

// TODO -- need to define a return value for this function and
//  its helper functions
func (e *ExecutionEngine) Execute(cmd *protocol.Command) {
	switch cmd.Type() {
	case protocol.GET:
		e.executeGet(cmd)
	case protocol.PUT:
		e.executePut(cmd)
	}
}

func (e *ExecutionEngine) executeGet(cmd *protocol.Command) {
	e.mem.Get(cmd.key)
}

func (e *ExecutionEngine) executePut(cmd *protocol.Command) {
	e.mem.Put(cmd.key, cmd.value)
}
