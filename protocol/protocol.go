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
	if len(buffer) < 3 {
		return nil, 0, errors.New("buffer is too small")
	}

	tp, err := getCommandType(buffer[0])

	if err != nil {
		return nil, 0, err
	}

	cmd := new(Command)
	cmd.tp = tp

	keySize := binary.LittleEndian.Uint16(buffer[1:3])

	if int(keySize) > MaxKeySize || keySize < 0 {
		return nil, 0, errors.New("provided key is too long or negative")
	}

	if len(buffer)-3 < int(keySize) {
		return nil, 0, errors.New("buffer does not contain the entire key")
	}

	cmd.key = string(buffer[3 : 3+keySize])
	var count = 1 + 2 + int(keySize)

	if cmd.tp == PUT {
		opBuf := buffer[3+keySize:]

		if len(opBuf) < 3 {
			return nil, 0, errors.New("buffer does not contain operand byte")
		}

		op, opCount, err := getOperand(opBuf)

		if err != nil {
			return nil, 0, err
		}

		cmd.value = op
		count += opCount
	}

	return cmd, count, nil
}

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

// Get the operand from the operand buffer
// Returns an error if the operand byte value is invalid
func getOperand(opBuf []byte) (interface{}, int, error) {
	switch opBuf[0] {
	case 0:
		return int(binary.LittleEndian.Uint32(opBuf[1:5])), 5, nil
	case 1:
		opSize := binary.LittleEndian.Uint16(opBuf[1:3])
		return string(opBuf[3 : 3+opSize]), int(2 + opSize), nil
	default:
		return nil, -1, errors.New("invalid operand byte")
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
