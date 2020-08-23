package test

import (
	"flashcache/protocol"
	"math"
	"reflect"
	"testing"
)

// TODO: break up the large test case functions into many smaller ones, which will help with readability
//  and maintainability of these tests in the future

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
