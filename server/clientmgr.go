package server

import (
	"errors"
	"log"
	"net"
	"strconv"
)

type ClientManager struct {
	config *ServerConfig
}

type CmdType int

const (
	PUT CmdType = iota
	GET
)

type OpType int

const (
	INT OpType = iota
	STRING
)

type Command struct {
	cmdType CmdType
	opType OpType
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
		_, err := readCommand(conn)

		if err != nil {
			log.Println("Client error:", err)
			log.Println("Quitting client connection...")

			_ = conn.Close()
			return
		}
	}
}

func makeListener(config *ServerConfig) (net.Listener, error) {
	address := makeAddressString(config)
	return net.Listen(config.protoStr, address)
}

func makeAddressString(config *ServerConfig) string {
	switch config.connType {
	case TCPv4, TCPv6:
		return config.address + ":" + strconv.Itoa(config.port)
	default:
		return config.address
	}
}

func readCommand(conn net.Conn) (*Command, error) {
	var cmdByte = make([]byte, 1)
	n, err := conn.Read(cmdByte)

	if err != nil {
		return nil, err
	}

	if n != 1 {
		return nil, errors.New("could not read leading command byte")
	}

	// The format of the command byte is as follows:
	// 7         5 4       2 1             0
	// -------------------------------------
	// | Protocol | Command | Operand Type |
	// -------------------------------------

	proto := (cmdByte[0] & 0xE0) >> 5
	var command = CmdType((cmdByte[0] & 0x1B) >> 2)
	var operand = OpType(cmdByte[0] & 3)

	if proto != 0 {
		return nil, errors.New("unknown (or unsupported) protocol version specified")
	}

	var cmd Command

	cmd.cmdType = command
	cmd.opType = operand

	// TODO read the other bytes and parse the information
	return &cmd, nil
}
