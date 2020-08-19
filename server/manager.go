package server

import "flashcache/config"

// Server manager interface
type Manager interface {
	Start() error
	Shutdown()
}

// Top-level server struct
type TopLevelManager struct {
	client *ClientManager
}

// Create a new top-level manager, creating sub-managers using the provided configuration
func NewTopLevelManager(conf *config.Configuration) *TopLevelManager {
	srv := new(TopLevelManager)
	srv.client = NewClientManager(conf)

	return srv
}

// Start the top-level manager -- start all the sub-managers
func (srv *TopLevelManager) Start() error {
	return srv.client.Start()
}

// Shutdown the top-level manager -- shutdown all the sub-managers
func (srv *TopLevelManager) Shutdown() {
	srv.client.Shutdown()
}
