package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

//Register ...
func (ctrl UserController) Ping(c *gin.Context) {
	log.Println("++++++++++++++++++++++++++++++++++++++++++++++++===Here")
	c.JSON(http.StatusOK, gin.H{"message": "Ping"})
}
