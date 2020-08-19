package server

import (
	"flashcache/config"
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

func (srv *ClientManager) processConn(conn net.Conn) {
	// TODO
}
