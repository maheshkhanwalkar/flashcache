package server

type CommandType int

const (
	PUT CommandType = iota
	GET
)

type Command struct {
	tp CommandType
	key string
	value interface{}
}

