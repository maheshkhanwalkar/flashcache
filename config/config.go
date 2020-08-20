package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"strconv"
)

type ConnType int

const (
	TCPv4 ConnType = iota
	TCPv6
	Unix
)

type Configuration struct {
	connType ConnType
	address  string
	port     int

	protoStr string
}

// Raw JSON format for the server configuration
type JSONServerConfig struct {
	Address string
	Port    int
	Type    string
}

// Load a configuration from the provided server configuration file
// Returns *Configuration on success or returns an error if the reading/parsing failed
func NewConfiguration(file string) (*Configuration, error) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var raw JSONServerConfig

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	conf, err := parseServerConfig(&raw)

	if err != nil {
		return nil, err
	}

	return conf, nil
}

// Create a server based on the current configuration parameters
// Returns a listener on success or returns an error if the server could not be created
func (c *Configuration) MakeServer() (net.Listener, error) {
	address := makeAddress(c.address, c.port, c.connType)
	return net.Listen(c.protoStr, address)
}

// Connect to the server specified by the current connection parameters
// Returns a connection on success or returns an error if the connection could not be established
func (c *Configuration) MakeClient() (net.Conn, error) {
	address := makeAddress(c.address, c.port, c.connType)
	return net.Dial(c.protoStr, address)
}

// Create an address string based on the given connection type
func makeAddress(ip string, port int, connType ConnType) string {
	if connType == Unix {
		return ip
	}

	if connType == TCPv6 {
		// Ensure IPv6 address is enclosed in brackets
		if ip[0] != '[' {
			ip = "[" + ip + "]"
		}
	}

	return ip + ":" + strconv.Itoa(port)
}

// Convert the raw JSON object into a Configuration object, applying validation of
// the input fields, returning an error if any of the values are invalid
func parseServerConfig(raw *JSONServerConfig) (*Configuration, error) {
	var srv = new(Configuration)

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

	return srv, nil
}