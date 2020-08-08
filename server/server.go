package server

import (
	"strconv"
)

type Server struct {
	config *Configuration
}

func NewServer(config *Configuration) *Server {
	srv := Server{}
	srv.config = config

	return &srv
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
func (srv *Server) Start() error {
	// TODO
	return nil
}

func (srv *Server) Shutdown() {
	// TODO
}
