package main

import (
	"log"

	"github.com/joho/godotenv"
	"quickform/cmd"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Execute()
}
