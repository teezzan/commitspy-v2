package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/teezzan/commitspy/routes"
)

func init() {
	log.Println("Starting up,")
	env := os.Getenv("ENV")
	if env != "TEST" {
		log.Printf("Loading .env file")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		log.Printf("Loading .test.env file")
		err := godotenv.Load(".test.env")
		if err != nil {
			log.Fatal("Error loading .test.env bbfile")
		}
	}
}
func main() {
	routes.Run()
}
