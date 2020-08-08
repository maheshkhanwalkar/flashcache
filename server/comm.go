package server

import (
	"net"
	"strconv"
)

func makeListener(config *ServerConfig) (net.Listener, error) {
	address := makeAddressString(config)
	return net.Listen(config.protoStr, address)
}

func makeAddressString(config *ServerConfig) string {
	switch config.connType {
	case TCPv4, TCPv6:
		return config.address + ":" + strconv.Itoa(config.port)
	default:
		return config.address
	}
}

func readCommandBytes(conn net.Conn) []byte {
	// TODO read the bytes from the socket connection
	return nil
}
