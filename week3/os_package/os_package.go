package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get values from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Use the values
	fmt.Println("DB_USER:", dbUser)
	fmt.Println("DB_PASS:", dbPass)
}
