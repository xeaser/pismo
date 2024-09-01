package main

import (
	"log"

	"github.com/xeaser/pismo/config"
	"github.com/xeaser/pismo/internal/server"
)

func main() {
	log.Println("Starting Server...")
	config.Init()       // Initialize the configuration
	cfg := config.Get() // Retrieve the configuration

	// Create a new server instance with the configuration
	server := server.NewServer(&cfg.Server)
	if err := server.Start(); err != nil {
		log.Fatal(err) // Log any errors that occur during server start
	}
}
