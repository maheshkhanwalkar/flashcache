package protocol

type CommandType int

const (
	PUT CommandType = iota
	GET
)

type Command struct {
	tp    CommandType
	key   string
	value interface{}
}

// Create a new GET command
func NewGet(key string) *Command {
	var cmd Command

	cmd.tp = GET
	cmd.key = key
	cmd.value = nil

	return &cmd
}

// Create a new PUT command
func NewPut(key string, value interface{}) *Command {
	var cmd Command

	cmd.tp = PUT
	cmd.key = key
	cmd.value = value

	return &cmd
}

func (c *Command) Type() CommandType {
	return c.tp
}

func (c *Command) Key() string {
	return c.key
}

func (c *Command) Value() interface{} {
	return c.value
}
