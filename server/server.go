package server

import (
	"errors"
	"net"
	"strconv"
)

type FCServer struct {
	Config *ServerConfig
}

// Get the equivalent network string for the specified connection type.
// This string is to be used with the net.Listen* family of functions
func getNetworkString(cType ConnType) string {
	var network string

	switch cType {
	case TCPv4:
		network = "tcp4"
	case TCPv6:
		network = "tcp6"
	case Unix:
		network = "unix"
	}

	return network
}

// Get the correct address string for the specified connection type
//   - For Unix domain sockets, the raw address is the final address
//   - For TCP sockets, the address should be rawAddress:port
func getAddressString(cType ConnType, rawAddress string, port int) string {
	var address string

	switch cType {
	case TCPv4, TCPv6:
		address = rawAddress + ":" + strconv.Itoa(port)
	case Unix:
		address = rawAddress
	}

	return address
}

// Start the server, returning an error if it is forced to quit in
// an unexpected manner
func (srv *FCServer) Start() error {
	conf := srv.Config
	cType := GetConnType(conf.Type)

	// Handle unknown case
	if cType == Unknown {
		return errors.New("Invalid server type specified: " + conf.Type)
	}

	network := getNetworkString(cType)
	address := getAddressString(cType, conf.Address, conf.Port)

	lst, err := net.Listen(network, address)

	if err != nil {
		return err
	}

	// TODO accept and spawn goroutines to process incoming clients
	// TODO support a shutdown mechanism to break out of the infinite accept loop

	return lst.Close()
}
