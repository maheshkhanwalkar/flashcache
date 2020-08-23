package protocol

import (
	"encoding/binary"
	"errors"
)

// Operand type
type OpType int

const (
	INTEGER OpType = iota
	STRING
)

// Operand
type Operand struct {
	tp OpType
	data interface{}
}

const (
	MaxStringSize = 512
)

// Read an integer from the given slice and return a new slice starting at the first unprocessed byte
// Returns an error if the provided slice is smaller than 4 bytes long
func ReadInt(buffer []byte) (int, []byte, error) {
	if len(buffer) < 4 {
		return 0, nil, BufferTooSmallError{}
	}

	// FIXME: maybe it's worth switching from using int to explicitly using int32, since it is not guaranteed
	//  that int is 4 bytes -- it could be larger, causing silent and endian-specific issues...
	return int(int32(binary.LittleEndian.Uint32(buffer))), buffer[4:], nil
}

// Read a string from the given slice and return a new slice starting at the first unprocessed byte
// Returns an error if the provided slice is smaller than the specified string length
func ReadString(buffer []byte) (string, []byte, error) {
	sz, buffer, err := ReadInt(buffer)

	if err != nil {
		return "", nil, err
	}

	if sz > MaxStringSize {
		return "", nil, errors.New("specified string length is too long")
	}

	if sz < 0 {
		return "", nil, errors.New("specified string length is negative")
	}

	if len(buffer) < sz {
		return "", nil, BufferTooSmallError{}
	}

	return string(buffer[:sz]), buffer[sz:], nil
}

// Read an operand from the given slice
// Returns an error if an invalid operand is specified or the slice is too small
func ReadOperand(buffer []byte) (Operand, []byte, error) {
	tp := buffer[0]
	actual := buffer[1:]

	switch OpType(tp) {
	case INTEGER:
		data, next, err := ReadInt(actual)
		return Operand{tp: INTEGER, data: data}, next, err
	case STRING:
		data, next, err := ReadString(actual)
		return Operand{tp: STRING, data: data}, next, err
	default:
		return Operand{}, nil, errors.New("unknown operand specified")
	}
}

// Write the integer into the provided slice and return a new slice pointing to the first byte not processed.
// Returns an error if the slice is not large enough
func WriteInt(num int, buffer []byte) ([]byte, error) {
	if len(buffer) < 4 {
		return nil, BufferTooSmallError{}
	}

	binary.LittleEndian.PutUint32(buffer, uint32(num))
	return buffer[4:], nil
}

// Write the string into the provided slice
// Returns an error if the slice is not big enough
func WriteString(str string, buffer []byte) ([]byte, error) {
	if len(buffer) < 4 + len(str) {
		return nil, BufferTooSmallError{}
	}

	// error ignored because buffer is guaranteed large enough
	buffer, _ = WriteInt(len(str), buffer)
	copy(buffer, str)

	return buffer[len(str):], nil
}

// Compute the number of bytes needed to store the particular operand
func ComputeOperandSize(op *Operand) int {
	var sz = 1

	switch op.tp {
	case INTEGER:
		return sz + 4
	case STRING:
		return sz + 4 + len(op.data.(string))
	}

	return sz
}

// Write the operand into the provided slice
// Returns an error if the slice is not large enough to store the operand
func WriteOperand(op *Operand, buffer []byte) ([]byte, error) {
	var tp = byte(op.tp)

	if len(buffer) < ComputeOperandSize(op) {
		return nil, BufferTooSmallError{}
	}

	buffer[0] = tp

	switch op.tp {
	case INTEGER:
		buffer, _ = WriteInt(op.data.(int), buffer[1:])
	case STRING:
		buffer, _ = WriteString(op.data.(string), buffer[1:])
	}

	return buffer, nil
}
