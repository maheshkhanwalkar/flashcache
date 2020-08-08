package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type ConnType int

const (
	TCPv4 ConnType = iota
	TCPv6
	Unix
)

type ServerConfig struct {
	connType ConnType
	protoStr string
	address  string
	port     int
}

type Configuration struct {
	serverConf *ServerConfig
}

// Raw JSON format for the server configuration
type JSONServerConfig struct {
	Address string
	Port int
	Type string
}

func NewConfiguration(server string) (*Configuration, error) {
	data, err := ioutil.ReadFile(server)

	if err != nil {
		return nil, err
	}

	var raw JSONServerConfig

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	var conf Configuration
	var srv *ServerConfig

	srv, err = parseServerConfig(&raw)

	if err != nil {
		return nil, err
	}

	conf.serverConf = srv
	return &conf, nil
}

func parseServerConfig(raw *JSONServerConfig) (*ServerConfig, error) {
	var srv ServerConfig

	srv.address = raw.Address
	srv.port = raw.Port

	switch raw.Type {
	case "TCPv4":
		srv.connType = TCPv4
		srv.protoStr = "tcp4"
	case "TCPv6":
		srv.connType = TCPv6
		srv.protoStr = "tcp6"
	case "Unix":
		srv.connType = Unix
		srv.protoStr = "unix"
	default:
		msg := "invalid connection type: " + raw.Type
		return nil, errors.New(msg)
	}

	return &srv, nil
}
