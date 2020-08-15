package protocol

import "errors"

// Read a command from the input buffer and return the constructed object along with
// the number of bytes consumed. Returns an error if the input is invalid or not long enough
func ReadCommand(buffer []byte) (*Command, int, error) {
	// TODO
	return nil, 0, nil
}

// Convert the provided command into its raw byte form, which can be written over the network
// and parsed back into its original object form
func WriteCommand(cmd *Command) ([]byte, error) {
	// TODO
	return nil, nil
}

// Get the command segment value for the given command type
// Returns an error if the command type is invalid
func getCommandSegment(tp CommandType) (byte, error) {
	switch tp {
	case GET:
		return 0, nil
	case PUT:
		return 1, nil
	default:
		return -1, errors.New("invalid command type")
	}
}

// Get the operand segment value and size for the given command type
// Returns an error if the operand value type is invalid
func getOpSegment(value interface{}) (byte, int, error) {
	switch value.(type) {
	case int:
		return 0, 4, nil
	case string:
		return 1, len(value.(string)), nil
	default:
		return -1, 0, errors.New("invalid operand value")
	}
}
