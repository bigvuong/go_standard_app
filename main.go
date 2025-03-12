package main

import (
	"log"
	"standard_app/config"
	"standard_app/routers"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize router
	router := routers.SetupRouter()

	// Run the server
	log.Fatal(router.Run(":9001"))
}
