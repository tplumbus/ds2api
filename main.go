package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // changed default from 8080 to 3000 to avoid conflicts on my machine
	}

	server := NewServer()

	log.Printf("Starting ds2api server on port %s", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
