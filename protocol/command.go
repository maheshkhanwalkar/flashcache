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

// Read a command from the input buffer and return the constructed object along with a new slice
// which starts at the next byte to be consumed. Returns an error if the input is invalid or not long enough
func ReadCommand(buffer []byte) (*Command, []byte, error) {
	if len(buffer) == 0 {
		return nil, nil, BufferTooSmallError{}
	}

	tp, err := getCommandType(buffer[0])

	if err != nil {
		return nil, nil, err
	}

	cmd := new(Command)
	cmd.tp = tp

	key, buffer, err := ReadOperand(buffer[1:])

	if err != nil {
		return nil, nil, err
	}

	cmd.key = &key

	if hasValueOperand(cmd.tp) {
		op, buffer, err := ReadOperand(buffer)

		if err != nil {
			return nil, nil, err
		}

		cmd.value = &op
		return cmd, buffer, nil
	}

	return cmd, buffer, nil
}

// Convert the provided command into its raw byte form, which can be written over the network
// and parsed back into its original object form
func WriteCommand(cmd *Command) ([]byte, error) {
	var cb = byte(cmd.tp)

	// FIXME: the size computation for the key should really be delegated to the string handling
	//  code rather than doing it here -- abstraction leak!

	var keySize = ComputeOperandSize(cmd.key)
	var bufSize = 1 + keySize

	// Add operand size to the total buffer space
	if hasValueOperand(cmd.tp) {
		bufSize += ComputeOperandSize(cmd.value)
	}

	var buffer = make([]byte, bufSize)

	// Command byte, key size, key
	buffer[0] = cb
	next, _ := WriteOperand(cmd.key, buffer[1:])

	// Write the value operand if it exists
	if hasValueOperand(cmd.tp) {
		_, _ = WriteOperand(cmd.value, next)
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
