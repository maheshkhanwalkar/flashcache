package protocol

import (
	"encoding/binary"
	"errors"
)

const (
	MaxKeySize int = 256
)

// Read a command from the input buffer and return the constructed object along with a new slice
// which starts at the next byte to be consumed. Returns an error if the input is invalid or not long enough
func ReadCommand(buffer []byte) (*Command, []byte, error) {
	if len(buffer) == 0 {
		return nil, nil, BufferTooSmallError{}
	}

	// FIXME: probably should just reduce this to a switch statement, rather than having a
	//  separate function dedicated to do some sort of cast + validation
	tp, err := getCommandType(buffer[0])

	if err != nil {
		return nil, nil, err
	}

	cmd := new(Command)
	cmd.tp = tp

	key, buffer, err := ReadString(buffer[1:])

	if err != nil {
		return nil, nil, BufferTooSmallError{}
	}

	cmd.key = key

	if hasOperand(cmd.tp) {
		op, buffer, err := ReadOperand(buffer)

		if err != nil {
			return nil, nil, err
		}

		cmd.value = op.data
		return cmd, buffer, nil
	}

	return cmd, buffer, nil
}

// FIXME: update WriteCommand to use the operand helper methods, instead of what is currently
//  being done

// Convert the provided command into its raw byte form, which can be written over the network
// and parsed back into its original object form
func WriteCommand(cmd *Command) ([]byte, error) {
	cb, err := getCommandByte(cmd.tp)

	if err != nil {
		return nil, err
	}

	var keySize = len(cmd.key)

	if keySize > MaxKeySize {
		return nil, errors.New("provided key is too long or negative")
	}

	var bufSize = 1 + 2 + keySize
	var buffer = make([]byte, bufSize)

	// Command byte, key size, key
	buffer[0] = cb
	binary.LittleEndian.PutUint16(buffer[1:], uint16(bufSize))
	copy(buffer[2:], cmd.key)

	if cmd.tp == PUT {
		op, err := getOpByte(cmd.value)

		if err != nil {
			return nil, err
		}

		buffer = append(buffer, op)
		buffer = appendOp(buffer, cmd.value)
	}

	return buffer, nil
}

// Get the command byte value for the given command type
// Returns an error if the command type is invalid
func getCommandByte(tp CommandType) (byte, error) {
	switch tp {
	case GET:
		return 0, nil
	case PUT:
		return 1, nil
	default:
		return 0, errors.New("invalid command type")
	}
}

// Get the command type associated with the given command byte value
// Returns an error if the command byte is invalid
func getCommandType(cb byte) (CommandType, error) {
	switch cb {
	case 0:
		return GET, nil
	case 1:
		return PUT, nil
	default:
		return 0, errors.New("invalid command byte")
	}
}

// Returns whether the given command type takes an operand or not
func hasOperand(tp CommandType) bool {
	switch tp {
	case GET:
		return false
	case PUT:
		return true
	default:
		return false
	}
}

// Get the operand segment value and size for the given command type
// Returns an error if the operand value type is invalid
func getOpByte(value interface{}) (byte, error) {
	switch value.(type) {
	case int:
		return 0, nil
	case string:
		return 1, nil
	default:
		return 0, errors.New("invalid operand value")
	}
}

// Append the operand to the buffer
func appendOp(buffer []byte, value interface{}) []byte {
	switch value.(type) {
	case int:
		buffer = append(buffer, make([]byte, 4)...)
		binary.LittleEndian.PutUint32(buffer[len(buffer)-4:], value.(uint32))

	case string:
		buffer = append(buffer, make([]byte, 2)...)
		binary.LittleEndian.PutUint16(buffer[len(buffer)-2:], uint16(len(value.(string))))

		buffer = append(buffer, value.(string)...)
	}

	return buffer
}
