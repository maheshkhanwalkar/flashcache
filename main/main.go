package main

import (
	"flashcache/config"
	"fmt"
	"os"
)

func main() {
	fmt.Println("flashcache")
	fmt.Println("Reading configuration file....")

	var file string

	// Select configuration file
	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		file = "conf/server.json"
	}

	_, err := config.ParseConfig(file)

	if err != nil {
		fmt.Println("Error. Could not parse configuration file!")
		fmt.Println("Specific reason:", err)
		os.Exit(1)
	}
}
