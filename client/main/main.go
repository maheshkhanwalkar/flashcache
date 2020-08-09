package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("flashcache cli client")
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if res := scan.Scan(); !res {
			break
		}

		// TODO parse the line and construct the command
		_ = scan.Text()
	}
}
