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
	var quit = false

	for !quit {
		fmt.Print("> ")

		if res := scan.Scan(); !res {
			break
		}

		line := scan.Text()
		pieces := strings.Split(line, " ")

		switch pieces[0] {
		case "connect":

			if len(pieces) != 3 {
				fmt.Println("Bad connect command.")
				fmt.Println("Expected format: connect tcp{4,6}/unix {address:port,file}")
			}

			// Establish a connection
			var err error

			conn, err = net.Dial(pieces[1], pieces[2])

			if err != nil {
				fmt.Println("Could not connect to server. Staying in disconnected state.")
				conn = nil
			}

		case "disconnect":
			if conn != nil {
				_ = conn.Close()
			} else {
				fmt.Println("Not connected to a server, ignoring command")
				continue
			}

			conn = nil

		case "quit", "exit":
			if conn != nil {
				_ = conn.Close()
			}

			quit = true

		default:
			processCommand(pieces)
		}
	}
}

func processCommand(pieces []string) {
	// TODO
}
