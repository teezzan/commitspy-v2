package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":5000")
}

func LoadRoutesAndReturnRouter() *gin.Engine {
	getRoutes()
	return router
}

func getRoutes() {
	routerGroup := router.Group("/api")

	addUserRoutes(routerGroup)
	addProjectRoutes(routerGroup)
}
