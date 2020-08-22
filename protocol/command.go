package protocol

type CommandType int

const (
	PUT CommandType = iota
	GET
	ERR
)

type Command struct {
	tp    CommandType
	key   string
	value *Operand
}

// Create a new GET command
func NewGet(key string) *Command {
	var cmd = new(Command)

	cmd.tp = GET
	cmd.key = key
	cmd.value = nil

	return cmd
}

// Create a new PUT command
func NewPut(key string, value *Operand) *Command {
	var cmd = new(Command)

	cmd.tp = PUT
	cmd.key = key
	cmd.value = value

	return cmd
}

// Create a new ERR command
func NewError(msg string) *Command {
	var cmd = new(Command)

	cmd.tp = ERR
	cmd.key = msg
	cmd.value = nil

	return cmd
}


func (c *Command) Type() CommandType {
	return c.tp
}

func (c *Command) Key() string {
	return c.key
}

func (c *Command) Value() *Operand {
	return c.value
}
