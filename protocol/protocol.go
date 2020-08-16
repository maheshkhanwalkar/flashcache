package protocol

import (
	"encoding/binary"
	"errors"
)

const (
	MaxKeySize int = 256
)

// Read a command from the input buffer and return the constructed object along with
// the number of bytes consumed. Returns an error if the input is invalid or not long enough
func ReadCommand(buffer []byte) (*Command, int, error) {
	// TODO
	return nil, 0, nil
}

// Convert the provided command into its raw byte form, which can be written over the network
// and parsed back into its original object form
func WriteCommand(cmd *Command) ([]byte, error) {
	cb, err := getCommandByte(cmd.tp)

	if err != nil {
		return nil, err
	}

	var keySize = len(cmd.key)

	if keySize > MaxKeySize || keySize < 0 {
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

// Get the command segment value for the given command type
// Returns an error if the command type is invalid
func getCommandByte(tp CommandType) (byte, error) {
	switch tp {
	case GET:
		return 0, nil
	case PUT:
		return 1, nil
	default:
		return -1, errors.New("invalid command type")
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
		return -1, errors.New("invalid operand value")
	}
}

// Append the operand to the buffer
func appendOp(buffer []byte, value interface{}) []byte {
	switch value.(type) {
	case int:
		buffer = append(buffer, make([]byte, 4)...)
		binary.LittleEndian.PutUint32(buffer[len(buffer) - 4:], value.(uint32))

	case string:
		buffer = append(buffer, make([]byte, 2)...)
		binary.LittleEndian.PutUint16(buffer[len(buffer) - 2:], uint16(len(value.(string))))

		buffer = append(buffer, value.(string)...)
	}

	return buffer
}
