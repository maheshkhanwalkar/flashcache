package server

import "flashcache/config"

type ClientManager struct {
	conf *config.Configuration
}

// Create a new client manager using the provided server configuration
func NewClientManager(conf *config.Configuration) *ClientManager {
	srv := new(ClientManager)

	srv.conf = conf
	return srv
}

// Start the client manager, returning an error on failure
func (srv *ClientManager) Start() error {
	// TODO
	return nil
}

// Shutdown the client manager
func (srv *ClientManager) Shutdown() {
	// TODO
}
