package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/auth"
	"github.com/teezzan/commitspy/controllers"
)

var user = new(controllers.UserController)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.Use(auth.AuthenticateToken)

	users.GET("/ping", user.Ping)
	users.GET("/login", user.CreateOrLogin)
	users.GET("/details", user.GetUser)
}
