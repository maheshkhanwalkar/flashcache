package protocol

import "net"

func ReadCommand(conn net.Conn) (*Command, error) {
	// TODO
	return nil, nil
}

func WriteCommand(cmd *Command, conn net.Conn) error {
	// TODO
	return nil
}
