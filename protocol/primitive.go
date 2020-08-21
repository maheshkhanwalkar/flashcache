package protocol

import (
	"encoding/binary"
	"errors"
)

type Operand int

const (
	INTEGER Operand = iota
	STRING
)

const (
	MaxStringSize = 512
)

// FIXME: these methods should be optimised to reduce unnecessary slice creation
//  and copying -- since these operands are slow...

// Read an integer from the given slice
// Returns an error if the provided slice is smaller than 4 bytes long
func ReadInt(buffer []byte) (int, error) {
	if len(buffer) < 4 {
		return 0, errors.New("buffer is not large enough")
	}

	return int(binary.LittleEndian.Uint32(buffer)), nil
}

// Read a string from the given slice
// Returns an error if the provided slice is smaller than the specified string length
func ReadString(buffer []byte) (string, error) {
	sz, err := ReadInt(buffer)

	if err != nil {
		return "", err
	}

	if sz > MaxStringSize {
		return "", errors.New("specified string length is too long")
	}

	if len(buffer) < 4 + sz {
		return "", errors.New("buffer is not large enough")
	}

	strBuf := buffer[4 : 4 + sz]
	return string(strBuf), nil
}

// Read an operand from the given slice
// Returns an error if an invalid operand is specified or the slice is too small
func ReadOperand(buffer []byte) (Operand, interface{}, error) {
	tp := buffer[0]
	actual := buffer[1:]

	switch Operand(tp) {
	case INTEGER:
		data, err := ReadInt(actual)
		return INTEGER, data, err
	case STRING:
		data, err := ReadString(actual)
		return STRING, data, err
	default:
		return 0, nil, errors.New("unknown operand specified")
	}
}

// Write the integer into a slice and return the result
func WriteInt(num int) []byte {
	buffer := make([]byte, 4)

	binary.LittleEndian.PutUint32(buffer, uint32(num))
	return buffer
}

// Write the string into a slice and return the result
func WriteString(str string) []byte {
	buffer := make([]byte, 4 + len(str))

	copy(buffer, WriteInt(len(str)))
	copy(buffer[4:], str)

	return buffer
}

// Write the operand into a slice and return the result
func WriteOperand(op Operand, data interface{}) []byte {
	var tp = byte(op)
	var payload []byte

	switch op {
	case INTEGER:
		payload = WriteInt(data.(int))
	case STRING:
		payload = WriteString(data.(string))
	}

	var full = make([]byte, 1 + len(payload))

	full[0] = tp
	copy(full[1:], payload)

	return full
}
