package main

import (
	"flashcache/config"
	"flashcache/server"
	"log"
	"os"
)

func main() {
	log.Println("flashcache")
	log.Println("Reading configuration file....")

	var file string

	// Select configuration file
	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		file = "conf/server.json"
	}

	conf, err := config.ParseConfig(file)

	if err != nil {
		log.Fatalln("Error. Could not parse configuration file:", err)
	}

	log.Println("Starting server...")

	srv := server.FCServer{Config: conf}
	srv.Start()
}
