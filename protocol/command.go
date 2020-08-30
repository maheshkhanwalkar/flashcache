package protocol

import "errors"

type CommandType int

const (
	PUT CommandType = iota
	GET
	ERR
)

type Command struct {
	tp    CommandType
	ops []Operand
}

// Create a new GET command
func NewGet(key string) *Command {
	var cmd = new(Command)

	cmd.tp = GET
	cmd.ops = make([]Operand, 1)

	cmd.ops[0] = Operand{tp: STRING, data: key}
	return cmd
}

// Create a new PUT command
func NewPut(key string, value *Operand) *Command {
	var cmd = new(Command)

	cmd.tp = PUT
	cmd.ops = make([]Operand, 2)

	cmd.ops[0] = Operand{tp: STRING, data: key}
	cmd.ops[1] = *value

	return cmd
}

// Create a new ERR command
func NewError(msg string) *Command {
	var cmd = new(Command)

	cmd.tp = ERR
	cmd.ops = make([]Operand, 1)

	cmd.ops[0] = Operand{tp: STRING, data: msg}
	return cmd
}


func (c *Command) Type() CommandType {
	return c.tp
}

func (c *Command) Key() string {
	return c.ops[0].data.(string)
}

func (c *Command) Value() *Operand {
	if len(c.ops) == 1 {
		return nil
	}

	return &c.ops[1]
}

// Read a command from the input buffer and return the constructed object along with a new slice
// which starts at the next byte to be consumed. Returns an error if the input is invalid or not long enough
func ReadCommand(buffer []byte) (*Command, []byte, error) {
	if len(buffer) == 0 {
		return nil, nil, BufferTooSmallError{}
	}

	// Read command byte
	tp, err := getCommandType(buffer[0])

	if err != nil {
		return nil, nil, err
	}

	cmd := new(Command)
	cmd.tp = tp

	// Read number of operands
	numOps, buffer, err := ReadInt(buffer[1:])

	if err != nil {
		return nil, nil, err
	}

	cmd.ops = make([]Operand, numOps)

	// Read each of the operands
	for i := 0; i < numOps; i++ {
		var op Operand
		op, buffer, err = ReadOperand(buffer)

		if err != nil {
			return nil, nil, err
		}

		cmd.ops[i] = op
	}

	return cmd, buffer, nil
}

// Convert the provided command into its raw byte form, which can be written over the network
// and parsed back into its original object form
func WriteCommand(cmd *Command) ([]byte, error) {
	var cb = byte(cmd.tp)
	var numOps = len(cmd.ops)

	var bufSize = 1 + 4

	for i := 0; i < numOps; i++ {
		bufSize += ComputeOperandSize(&cmd.ops[i])
	}

	var buffer = make([]byte, bufSize)
	buffer[0] = cb

	next, _ := WriteInt(numOps, buffer[1:])

	for i := 0; i < numOps; i++ {
		next, _ = WriteOperand(&cmd.ops[i], next)
	}

	return buffer, nil
}

// Get the command type associated with the given command byte value
// Returns an error if the command byte is invalid
func getCommandType(cb byte) (CommandType, error) {
	equiv := CommandType(cb)

	switch equiv {
	case GET, PUT, ERR:
		return equiv, nil
	default:
		return 0, errors.New("invalid command byte")
	}
}

// Returns whether the given command type takes an value operand or not
func hasValueOperand(tp CommandType) bool {
	switch tp {
	case GET, ERR:
		return false
	case PUT:
		return true
	default:
		return false
	}
}
