package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

//Register ...
func (ctrl UserController) Ping(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Ping"})
}
