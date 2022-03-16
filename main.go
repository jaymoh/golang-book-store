package main

import (
	"github.com/joho/godotenv"
	"hackinroms.com/books/database"
	"hackinroms.com/books/routes"
	"log"
)

func main() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup DB instance
	database.Setup()

	// setup routes
	r := routes.Setup()

	// launch app
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
