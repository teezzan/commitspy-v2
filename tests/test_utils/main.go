package test_utils

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/teezzan/commitspy/routes"
)

func SetupRouter() {
	err := godotenv.Load(".test.env")
	if err != nil {
		log.Println("err: ", err)
		log.Fatal("Error loading .test.env file")
	}
	router := routes.LoadRoutesAndReturnRouter()
	router.Run(":5000")

}
