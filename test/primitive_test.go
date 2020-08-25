package test

import (
	"flashcache/protocol"
	"math"
	"reflect"
	"testing"
)

// Test a range of positive, zero, and negative integers
func TestRangeInt(t *testing.T) {
	buffer := make([]byte, 4)

	for i := -1000; i < 1000; i++ {
		testReadWriteInt(i, buffer, t)
	}
}

// Test both integer extremes (max and min integers) within ReadInt and WriteInt
func TestExtremeInt(t *testing.T) {
	buffer := make([]byte, 4)

	// Test both extremes
	testReadWriteInt(math.MaxInt32, buffer, t)
	testReadWriteInt(math.MinInt32, buffer, t)
}

// Test small buffer for WriteInt
func TestSmallBufferWriteInt(t *testing.T) {
	small := make([]byte, 1)
	_, err := protocol.WriteInt(0, small)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
}

// Test small buffer for ReadInt
func TestSmallBufferReadInt(t *testing.T) {
	small := make([]byte, 1)
	_, _, err := protocol.ReadInt(small)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
}

// Test buffer is larger for WriteInt
func TestLargeBufferWriteInt(t *testing.T) {
	over := 100
	large := make([]byte, 4 + over)

	diff, err := protocol.WriteInt(0, large)

	AssertEqual(err, nil, t)
	AssertEqual(len(diff), over, t)

}

// Test buffer is larger for ReadInt
func TestLargeBufferReadInt(t *testing.T) {
	over := 100
	large := make([]byte, 4 + over)


	_, diff, err := protocol.ReadInt(large)

	AssertEqual(err, nil, t)
	AssertEqual(len(diff), over, t)
}

// Test the serialisation and de-serialisation of empty strings within the protocol
func TestEmptyString(t *testing.T) {
	buffer := make([]byte, 4)

	// Test empty string
	next, err := protocol.WriteString("", buffer)

	AssertEqual(len(next), 0, t)
	AssertEqual(err, nil, t)

	res, next, err := protocol.ReadString(buffer)

	AssertEqual(len(next), 0, t)
	AssertEqual(err, nil, t)
	AssertEqual(res, "", t)
}

func TestSmallBufferWriteString(t *testing.T) {
	buffer := make([]byte, 6)

	// Buffer large enough for size, but too small for actual string
	_, err := protocol.WriteString("hello", buffer)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)

	// Buffer too small for even the size
	buffer = buffer[:2]
	_, err = protocol.WriteString("hello", buffer)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
}

func TestSmallBufferReadString(t *testing.T) {
	buffer := make([]byte, 6)
	_, _ = protocol.WriteInt(5, buffer)

	// Buffer large enough for size, but too small for actual string
	_, _, err := protocol.ReadString(buffer)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)

	// Buffer too small for even the size
	buffer = buffer[:2]
	_, _, err = protocol.ReadString(buffer)

	AssertNotEqual(err, nil, t)
	AssertEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
}

// TODO add tests for larger buffer for {Read, Write}String functions

func TestInvalidStringLength(t *testing.T) {
	buffer := make([]byte, 10)
	_, _ = protocol.WriteInt(-5, buffer)

	// String length is negative -- error, but not BufferTooSmallError
	_, _, err := protocol.ReadString(buffer)

	AssertNotEqual(err, nil, t)
	AssertNotEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)

	// String length exceeds the maximum string length allowed by the protocol
	_, _ = protocol.WriteInt(protocol.MaxStringSize + 1, buffer)
	_, _, err = protocol.ReadString(buffer)

	AssertNotEqual(err, nil, t)
	AssertNotEqual(reflect.TypeOf(err), reflect.TypeOf(protocol.BufferTooSmallError{}), t)
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
