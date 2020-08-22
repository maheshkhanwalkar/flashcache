package server

import (
	"flashcache/config"
	"flashcache/protocol"
	"net"
	"sync/atomic"
)

type ClientManager struct {
	conf *config.Configuration
	lst net.Listener
	quit atomic.Value
}

// Create a new client manager using the provided server configuration
func NewClientManager(conf *config.Configuration) *ClientManager {
	srv := new(ClientManager)

	srv.conf = conf
	srv.quit.Store(false)

	return srv
}

// Start the client manager, returning an error on failure
func (srv *ClientManager) Start() error {
	var err error
	srv.lst, err = srv.conf.MakeServer()

	if err != nil {
		return err
	}

	for {
		conn, err := srv.lst.Accept()

		if err != nil {
			if srv.quit.Load().(bool) {
				break
			} else {
				return err
			}
		}

		// Process the connection's requests
		go srv.processConn(conn)
	}

	return nil
}

// Shutdown the client manager
func (srv *ClientManager) Shutdown() {
	srv.quit.Store(true)
	_ = srv.lst.Close()
}

// Process a connection's requests
func (srv *ClientManager) processConn(conn net.Conn) {
	// Close the connection when the function returns
	defer conn.Close()

	var buffer = make([]byte, 4096)
	var offset = 0

	// TODO: the buffer manipulation details should be extracted out and placed somewhere else
	//  since this will be a problem in both the client and the server programs

	for {
		lim, err := conn.Read(buffer[offset:])

		// Stop processing this client and exit out
		if err != nil {
			_ = conn.Close()
			break
		}

		actual := buffer[:offset+lim]

		for len(actual) > 0 {
			cmd, next, err := protocol.ReadCommand(actual)

			if err != nil {
				// TODO should write an error msg back to the client and close the connection
				if _, ok := err.(protocol.BufferTooSmallError); !ok {
					return
				}

				// Copy the remaining bytes to the start of the buffer and adjust the offset
				// That way, the next read will not overwrite these bytes
				rem := len(actual)

				copy(buffer, buffer[len(buffer) - rem:])
				offset = rem

				break
			}

			// TODO process the command
			_ = cmd

			// Update the slice to process the next command
			actual = next
		}
	}
}
