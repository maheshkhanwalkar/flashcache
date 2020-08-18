package main

import (
	"bufio"
	"flashcache/config"
	"flashcache/server"
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
			var path string

			if len(pieces) == 1 {
				path = server.DefaultConfigPath
			} else {
				path = pieces[1]
			}

			conf, err := config.NewConfiguration(path)

			if err != nil {
				fmt.Println("Could not parse provided configuration file:", err)
				continue
			}

			conn, err = conf.MakeClient()

			if err != nil {
				fmt.Println("Could not connect to server:", err)
				continue
			}

			fmt.Println("Successfully connected to the server")

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
