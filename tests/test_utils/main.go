package test_utils

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/routes"
)

func SetupRouter() *gin.Engine {

	router := routes.LoadRoutesAndReturnRouter()
	return router

}
