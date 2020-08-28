package test

import (
	"flashcache/protocol"
	"reflect"
	"testing"
)

func TestGetReadWrite(t *testing.T) {
	cmd := protocol.NewGet("my-key")
	readWrite(cmd, t)
}

func TestEmptyOperands(t *testing.T) {
	cmd := protocol.NewGet("")
	readWrite(cmd, t)

	value, _ := protocol.NewOperand(protocol.STRING, "")
	cmd = protocol.NewPut("hello-world", value)
	readWrite(cmd, t)

	cmd = protocol.NewPut("", value)
	readWrite(cmd, t)
}

func TestSmallBuffer(t *testing.T) {
	buffer := make([]byte, 2)
	_, _, err := protocol.ReadCommand(buffer)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
}

func readWrite(cmd *protocol.Command, t *testing.T) {
	buffer, err := protocol.WriteCommand(cmd)

	AssertEqual(err, nil, t)

	dup, next, err := protocol.ReadCommand(buffer)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)

	AssertEqual(cmd.Type(), dup.Type(), t)
	AssertEqual(cmd.Key(), dup.Key(), t)

	// Check for value equality or both are nil
	if cmd.Value() != nil && dup.Value() != nil {
		AssertEqual(cmd.Value().Type(), dup.Value().Type(), t)
		AssertEqual(cmd.Value().Data(), dup.Value().Data(), t)
	} else {
		AssertEqual(cmd.Value(), dup.Value(), t)
	}
}
