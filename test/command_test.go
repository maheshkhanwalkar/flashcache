package test

import (
	"flashcache/protocol"
	"testing"
)

func TestGetReadWrite(t *testing.T) {
	cmd := protocol.NewGet("my-key")
	buffer, err := protocol.WriteCommand(cmd)

	AssertEqual(err, nil, t)

	dup, next, err := protocol.ReadCommand(buffer)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)

	AssertEqual(cmd.Type(), dup.Type(), t)
	AssertEqual(cmd.Key(), dup.Key(), t)
	AssertEqual(cmd.Value(), dup.Value(), t)
}
