package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/auth"
	"github.com/teezzan/commitspy-v2/controllers"
)

var project = new(controllers.Project)

func addProjectRoutes(rg *gin.RouterGroup) {
	projects := rg.Group("/project")

	projects.Use(auth.AuthenticateToken)

	projects.GET("/", project.FetchAll)
	projects.POST("/create", project.Create)
	projects.GET("/:id", project.FetchOne)
	projects.POST("/:id", project.Update)
	projects.DELETE("/:id", project.Delete)
}
