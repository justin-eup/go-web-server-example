package main

import (
	"log"

	"github.com/justin-eup/web-server-example/pkg/server"
)

func main() {
	// Initialize and start the server
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
