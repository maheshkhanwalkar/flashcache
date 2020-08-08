package server

import "net"

type ClientManager struct {
	config *ServerConfig
}

func (srv *ClientManager) Start() error {
	lst, err := makeListener(srv.config)

	if err != nil {
		return err
	}

	for {
		conn, err := lst.Accept()

		if err != nil {
			break
		}

		go srv.handleClient(conn)
	}

	_ = lst.Close()
	return nil
}

func (srv *ClientManager) Shutdown() {
	// TODO
}

func (srv *ClientManager) handleClient(conn net.Conn) {
	for {
		_ = readCommandBytes(conn)
	}
}
