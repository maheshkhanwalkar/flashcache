package protocol

import "errors"

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

// Construct a new operand with the given type and data
func NewOperand(tp OpType, data interface{}) (*Operand, error) {
	op := new(Operand)

	switch data.(type) {
	case int:
		if tp != INTEGER {
			return nil, errors.New("type and data mismatch for operand")
		}
	case string:
		if tp != STRING {
			return nil, errors.New("type and data mismatch for operand")
		}
	}

	op.tp = tp
	op.data = data

	return op, nil
}

// Get the operand type
func (op *Operand) Type() OpType {
	return op.tp
}

// Get the operand data
func (op *Operand) Data() interface{} {
	return op.data
}

// Read an operand from the given slice
// Returns an error if an invalid operand is specified or the slice is too small
func ReadOperand(buffer []byte) (Operand, []byte, error) {
	if len(buffer) < 2 {
		return Operand{}, nil, BufferTooSmallError{}
	}

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

// Compute the number of bytes needed to store the particular operand
func ComputeOperandSize(op *Operand) int {
	var sz = 1

	switch op.tp {
	case INTEGER:
		sz += 4
	case STRING:
		sz += 4 + len(op.data.(string))
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
