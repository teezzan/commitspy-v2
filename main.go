package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/teezzan/commitspy/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.Run()
}
