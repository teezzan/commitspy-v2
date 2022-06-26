package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/controllers"
	"github.com/teezzan/commitspy/middleware"
)

var user = new(controllers.UserController)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.Use(middleware.AuthenticateToken)

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/ping", user.Ping)
}
