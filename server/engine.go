package server

type ExecutionEngine struct {
	mem *MemoryStore
}

// TODO -- need to define a return value for this function and
//  its helper functions
func (e *ExecutionEngine) Execute(cmd *Command) {
	switch cmd.tp {
	case GET:
		e.executeGet(cmd)
	case PUT:
		e.executePut(cmd)
	}
}

func (e *ExecutionEngine) executeGet(cmd *Command) {
	e.mem.Get(cmd.key)
}

func (e *ExecutionEngine) executePut(cmd *Command) {
	e.mem.Put(cmd.key, cmd.value)
}
