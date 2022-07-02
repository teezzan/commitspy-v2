package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/controllers"
	"github.com/teezzan/commitspy/middleware"
)

var user = new(controllers.UserController)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.GET("/ping", user.Ping)

	users.Use(middleware.AuthenticateToken)

	users.GET("/login", user.CreateOrLogin)
	users.GET("/details", user.GetUser)
}
