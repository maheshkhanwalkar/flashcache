package test

import (
	"flashcache/protocol"
	"math"
	"reflect"
	"testing"
)

// Test the serialisation and de-serialisation of various integers within the protocol
func TestSerialiseInt(t *testing.T) {
	buffer := make([]byte, 4)

	// Test a range of numbers
	for i := -1000; i < 1000; i++ {
		testReadWriteInt(i, buffer, t)
	}

	// Test both extremes
	testReadWriteInt(math.MaxInt32, buffer, t)
	testReadWriteInt(math.MinInt32, buffer, t)

	// Test buffer too small -- WriteInt
	small := make([]byte, 1)
	_, err := protocol.WriteInt(0, small)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)

	// Test buffer too small -- ReadInt
	_, _, err = protocol.ReadInt(small)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)

	// Test buffer is larger -- WriteInt
	over := 100
	large := make([]byte, 4 + over)

	diff, err := protocol.WriteInt(0, large)

	AssertEqual(err, nil, t)
	AssertEqual(len(diff), over, t)

	// Test buffer is larger -- ReadInt
	_, diff, err = protocol.ReadInt(large)

	AssertEqual(err, nil, t)
	AssertEqual(len(diff), over, t)
}

// Test writing and reading back of an integer to a buffer
func testReadWriteInt(i int, buffer []byte, t *testing.T) {
	next, err := protocol.WriteInt(i, buffer)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)

	equiv, next, err := protocol.ReadInt(buffer)

	AssertEqual(err, nil, t)
	AssertEqual(len(next), 0, t)
	AssertEqual(i, equiv, t)
}
