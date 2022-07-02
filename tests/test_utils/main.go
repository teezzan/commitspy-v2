package test_utils

import (
	"github.com/teezzan/commitspy/routes"
)

func SetupRouter() {

	router := routes.LoadRoutesAndReturnRouter()
	router.Run(":5000")

}
