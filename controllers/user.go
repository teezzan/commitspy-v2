package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/account"
	"github.com/teezzan/commitspy/auth"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/response"
)

var db = database.GetDB()

type UserController struct{}

func (ctrl UserController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ping"})
}

func (ctrl UserController) CreateOrLogin(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	user, err := database.GetUserByExternalID(userCtx.ExternalID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if user != nil {
		response.WriteSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	}

	newUser := &account.User{
		ExternalID: userDetails.ExternalID,
		Email:      userDetails.Email,
		Avatar:     userDetails.Avatar,
		Name:       userDetails.Name,
	}

	user, result = database.CreateUser(newUser)

	if result.Error != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if result.RowsAffected > 0 {
		response.WriteSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	}
}

func (ctrl UserController) GetUser(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	user, result := GetUserDetails(userCtx.ExternalID)

	if result.Error != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if result.RowsAffected > 0 {
		response.WriteSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	} else {
		response.WriteError(c, http.StatusNotFound, gin.H{"error": "User not found"})
	}

}
