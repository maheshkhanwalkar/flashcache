package protocol

import (
	"encoding/binary"
	"errors"
)

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
