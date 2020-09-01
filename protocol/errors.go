package protocol

// Error indicating that the provided buffer slice is too small, that is, the
// current length should be longer to contain the entire payload
type BufferTooSmallError struct {
}

func (err BufferTooSmallError) Error() string {
	return "buffer is too small"
}

// Error indicating that the provided string length is invalid, either too long
// or negative -- neither of which are permitted
type InvalidStringLengthError struct {
	msg string
}

func (err InvalidStringLengthError) Error() string {
	return "invalid string length: " + err.msg
}
