package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/database"
)

var db = database.GetDB()

type UserController struct{}

func (ctrl UserController) Ping(c *gin.Context) {
	userCtx := GetCtxUser(c)
	log.Println(userCtx)
	c.JSON(http.StatusOK, gin.H{"message": "Ping"})
}

func (ctrl UserController) CreateOrLogin(c *gin.Context) {
	userCtx := GetCtxUser(c)

	user, result := GetUserDetails(userCtx.ExternalID)

	if result.Error != nil {
		RespondWithError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if result.RowsAffected > 0 {
		RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	}

	user, result = CreateUser(userCtx)

	if result.Error != nil {
		RespondWithError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if result.RowsAffected > 0 {
		RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	}
}

func (ctrl UserController) GetUser(c *gin.Context) {
	userCtx := GetCtxUser(c)

	user, result := GetUserDetails(userCtx.ExternalID)

	if result.Error != nil {
		RespondWithError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if result.RowsAffected > 0 {
		RespondWithSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	} else {
		RespondWithError(c, http.StatusNotFound, gin.H{"error": "User not found"})
	}

}
