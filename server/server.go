package server

type Server struct {
	config *Configuration
	client *ClientManager
}

func NewServer(config *Configuration) *Server {
	srv := Server{}
	client := ClientManager{}

	client.config = config.serverConf
	srv.config = config
	srv.client = &client

	return &srv
}

// Start the server, returning an error if it is forced to quit in
// an unexpected manner
func (srv *Server) Start() error {
	return srv.client.Start()
}

func (srv *Server) Shutdown() {
	srv.client.Shutdown()
}
