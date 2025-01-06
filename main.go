package main

import (
	"learn-go-auth/config"
	"learn-go-auth/routes"
	"log"
)

func main() {
	// Initialize database connection
	config.ConnectDatabase()

	// Setup routes
	r := routes.SetupRouter()

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
