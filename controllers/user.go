package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/models"
)

//UserController ...
type UserController struct{}

var db = database.GetDB()

//Register ...
func (ctrl UserController) Ping(c *gin.Context) {
	db.Create(&models.User{ID: 2, ExternalID: "quuq", Name: "Tee", Avatar: "yay"})
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Ping"})
}
