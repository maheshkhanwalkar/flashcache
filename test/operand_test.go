package test

import (
	"flashcache/protocol"
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