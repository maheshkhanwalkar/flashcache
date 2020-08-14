package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("flashcache cli client")
	scan := bufio.NewScanner(os.Stdin)

	var conn net.Conn = nil

	for {
		fmt.Print("> ")

		if res := scan.Scan(); !res {
			break
		}

		line := scan.Text()
		pieces := strings.Split(line, " ")

		switch pieces[0] {
		case "connect":
			var config string

			// Ignore badly formatted command
			if len(pieces) > 2 {
				fmt.Println("Bad connect command -- too many arguments, ignoring command")
				continue
			}

			// FIXME: need to have a constants source file -- so if the default changed between the server
			//  the client will also see the change (automatically via constant sharing)
			if len(pieces) == 1 {
				config = "conf/server.json"
			} else {
				config = pieces[1]
			}

			// Establish a connection
			conn = makeConnection(config)

		case "disconnect":
			if conn != nil {
				_ = conn.Close()
			} else {
				fmt.Println("Not connected to a server, ignoring command")
				continue
			}

			conn = nil

		default:
			processCommand(pieces)
		}
	}
}

func makeConnection(serverConfig string) net.Conn {
	// TODO
	return nil
}

func processCommand(pieces []string) {
	// TODO
}
