package server

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type ConnType int

const (
	TCPv4 ConnType = iota
	TCPv6
	Unix
	Unknown
)

type ServerConfig struct {
	Type string
	Address string
	Port int
}

func ParseConfig(config string) (*ServerConfig, error) {
	raw, err := ioutil.ReadFile(config)

	if err != nil {
		return nil, err
	}

	// Setup default values for port and address
	conf := ServerConfig{Port: 0, Address: "localhost"}
	err = json.Unmarshal(raw, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func GetConnType(cType string) ConnType {
	switch strings.ToUpper(cType) {
	case "TCPV4":
		return TCPv4
	case "TCPV6":
		return TCPv6
	case "UNIX":
		return Unix
	default:
		return Unknown
	}
}
