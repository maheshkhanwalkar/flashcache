package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("flashcache")
	fmt.Println("Reading configuration file....")

	var config string

	// Select configuration file
	if len(os.Args) > 1 {
		config = os.Args[1]
	} else {
		config = "conf/server.yml"
	}

	// TODO: remove this line -- only here to placate the compiler
	fmt.Println(config)
}
