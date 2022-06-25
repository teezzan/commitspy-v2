package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/controllers"
)

var user = new(controllers.UserController)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
	users.GET("/ping", user.Ping)
}
