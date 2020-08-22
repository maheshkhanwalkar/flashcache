package protocol

// Error indicating that the provided buffer slice is too small, that is, the
// current length should be longer to contain the entire payload
type BufferTooSmallError struct {
}

func (err BufferTooSmallError) Error() string {
	return "buffer is too small"
}
