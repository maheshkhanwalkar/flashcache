package protocol

type PartialCommandError struct {
	msg string
}

// Return the error message
func (err PartialCommandError) Error() string {
	return err.msg
}
