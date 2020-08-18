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
		file = server.DefaultConfigPath
	}

	_, err := config.NewConfiguration(file)

	if err != nil {
		log.Fatalln("Error. Could not parse configuration file:", err)
	}

	log.Println("Starting server...")

	// srv := server.NewServer(conf)
	// setupShutdown(srv)

	/* err = srv.Start()

	if err != nil {
		log.Fatalln("Error. Server quit unexpectedly:", err)
	}
	*/

	log.Println("Server shutdown")
}

/*func setupShutdown(srv *server.Server) {
	log.Println("Setting up shutdown hook...")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT)

	go func() {
		_ = <- sig

		log.Println("Shutting down server...")
		srv.Shutdown()
	}()
}
*/