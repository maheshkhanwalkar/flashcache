package protocol

import (
	"errors"
)

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

	// TODO: perhaps the key should be generalised to not necessarily assume a string, but
	//  could be *anything*, i.e. represent it using an operand just like value...
	key, buffer, err := ReadString(buffer[1:])

	if err != nil {
		return nil, nil, err
	}

	cmd.key = key

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

	var keySize = len(cmd.key)

	if keySize > MaxStringSize || keySize < 0 {
		return nil, errors.New("provided key is too long or negative")
	}

	var bufSize = 1 + 2 + keySize

	// Add operand size to the total buffer space
	if hasValueOperand(cmd.tp) {
		bufSize += ComputeOperandSize(cmd.value)
	}

	var buffer = make([]byte, bufSize)

	// Command byte, key size, key
	buffer[0] = cb
	next, _ := WriteString(cmd.key, buffer[1:])

	// Write the value operand if it exists
	if hasValueOperand(cmd.tp) {
		_, _ = WriteOperand(cmd.value, next)
	}

	return buffer, nil
}

// Get the command byte value for the given command type
// Returns an error if the command type is invalid
func getCommandByte(tp CommandType) (byte, error) {
	switch tp {
	case GET, PUT:
		return byte(tp), nil
	default:
		return 0, errors.New("invalid command type")
	}
}

// Get the command type associated with the given command byte value
// Returns an error if the command byte is invalid
func getCommandType(cb byte) (CommandType, error) {
	equiv := CommandType(cb)

	switch equiv {
	case GET, PUT:
		return equiv, nil
	default:
		return 0, errors.New("invalid command byte")
	}
}

// Returns whether the given command type takes an value operand or not
func hasValueOperand(tp CommandType) bool {
	switch tp {
	case GET:
		return false
	case PUT:
		return true
	default:
		return false
	}
}
