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
	log.Printf("Environment: %s", getEnv("APP_ENV", "development"))
	log.Printf("Log level: %s", getEnv("LOG_LEVEL", "info"))
	log.Printf("Debug mode: %s", getEnv("DEBUG", "false"))
	log.Printf("Version: %s", getEnv("APP_VERSION", "dev"))
	log.Printf("Server address: http://localhost:%s", port)
	log.Printf("API docs: http://localhost:%s/docs", port) // handy reminder when starting up
	if err := server.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getEnv returns the value of an environment variable or a fallback default.
// If the environment variable is not set, the fallback value is used.
// This is useful for providing sensible defaults in development environments.
// Note: empty string values are treated as unset; use a placeholder like "0" or "false" if needed.
// TODO: consider supporting a .env.local override file for machine-specific secrets (not committed to git)
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
