package server

import (
	"errors"
	"flashcache/config"
	"syscall"
)

type FCServer struct {
	Config *config.ServerConfig
}

func (srv *FCServer) Start() error {
	conf := srv.Config
	cType := config.GetConnType(conf.Type)

	// Handle unknown case
	if cType == config.Unknown {
		return errors.New("Invalid server type specified: " + conf.Type)
	}

	var domain int

	switch cType {
	case config.TCPv4:
		domain = syscall.AF_INET
	case config.TCPv6:
		domain = syscall.AF_INET6
	case config.Unix:
		domain = syscall.AF_UNIX
	}

	fd, err := syscall.Socket(domain, syscall.SOCK_STREAM, 0)

	// Error opening the socket
	if err != nil {
		return err
	}

	// TODO finish up socket init + enter "infinite" loop

	// Close down the socket
	err = syscall.Close(fd)
	return err
}
