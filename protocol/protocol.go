package protocol

// Read a command from the input buffer and return the constructed object along with
// the number of bytes consumed. Returns an error if the input is invalid or not long enough
func ReadCommand(buffer []byte) (*Command, int, error) {
	// TODO
	return nil, 0, nil
}

// Convert the provided command into its raw byte form, which can be written over the network
// and parsed back into its original object form
func WriteCommand(cmd *Command) []byte {
	// TODO
	return nil
}
