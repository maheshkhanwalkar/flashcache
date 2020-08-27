package protocol

type CommandType int

const (
	PUT CommandType = iota
	GET
	ERR
)

type Command struct {
	tp    CommandType
	key   *Operand
	value *Operand
}

// Create a new GET command
func NewGet(key string) *Command {
	var cmd = new(Command)

	cmd.tp = GET
	cmd.key = &Operand{tp: STRING, data: key}
	cmd.value = nil

	return cmd
}

// Create a new PUT command
func NewPut(key string, value *Operand) *Command {
	var cmd = new(Command)

	cmd.tp = PUT
	cmd.key = &Operand{tp: STRING, data: key}
	cmd.value = value

	return cmd
}

// Create a new ERR command
func NewError(msg string) *Command {
	var cmd = new(Command)

	cmd.tp = ERR
	cmd.key = &Operand{tp: STRING, data: msg}
	cmd.value = nil

	return cmd
}


func (c *Command) Type() CommandType {
	return c.tp
}

func (c *Command) Key() string {
	return c.key.data.(string)
}

func (c *Command) Value() *Operand {
	return c.value
}
