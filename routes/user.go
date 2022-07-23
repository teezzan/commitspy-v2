package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/auth"
	"github.com/teezzan/commitspy-v2/controllers"
)

var user = new(controllers.User)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/user")

	users.Use(auth.AuthenticateToken)

	users.GET("/ping", user.Ping)
	users.GET("/login", user.CreateOrLogin)
	users.GET("/details", user.GetUser)
}
