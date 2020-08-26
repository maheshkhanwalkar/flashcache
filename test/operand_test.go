package test

import (
	"flashcache/protocol"
	"math"
	"reflect"
	"testing"
)

// Test invalid operand creation -- type and data provided do not match
func TestOperandTypeDataMismatch(t *testing.T) {
	// Expect integer, got string
	_, err := protocol.NewOperand(protocol.INTEGER, "hello world")
	AssertNotEqual(err, nil, t)

	// Expect string, got integer
	_, err = protocol.NewOperand(protocol.STRING, 57)
	AssertNotEqual(err, nil, t)

	// Matching -- no error
	_, err = protocol.NewOperand(protocol.INTEGER, 100)
	AssertEqual(err, nil, t)

	_, err = protocol.NewOperand(protocol.STRING, "hello")
	AssertEqual(err, nil, t)
}

func TestIntOperand(t *testing.T) {
	op, err := protocol.NewOperand(protocol.INTEGER, 100)
	AssertEqual(err, nil, t)

	size := protocol.ComputeOperandSize(op)
	buf := make([]byte, size)

	next, err := protocol.WriteOperand(op, buf)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)

	dup, next, err := protocol.ReadOperand(buf)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)
	AssertEqual(dup.Type(), op.Type(), t)
	AssertEqual(dup.Data(), op.Data(), t)
}

func TestStringOperand(t *testing.T) {
	op, err := protocol.NewOperand(protocol.STRING, "hello world")
	AssertEqual(err, nil, t)

	size := protocol.ComputeOperandSize(op)
	buf := make([]byte, size)

	next, err := protocol.WriteOperand(op, buf)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)

	dup, next, err := protocol.ReadOperand(buf)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)
	AssertEqual(dup.Type(), op.Type(), t)
	AssertEqual(dup.Data(), op.Data(), t)
}

func TestBufferTooSmall(t *testing.T) {
	buffer := make([]byte, 1)
	_, _, err := protocol.ReadOperand(buffer)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)

	op, err := protocol.NewOperand(protocol.INTEGER, 5)
	_, err = protocol.WriteOperand(op, buffer)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
}

func TestInvalidOperandByte(t *testing.T) {
	buffer := make([]byte, 2)
	buffer[0] = math.MaxInt8

	_, _, err := protocol.ReadOperand(buffer)

	AssertNotEqual(err, nil, t)
	AssertNotEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
}
