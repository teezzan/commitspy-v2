package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/account"
	"github.com/teezzan/commitspy/auth"
	"github.com/teezzan/commitspy/database"
	"github.com/teezzan/commitspy/response"
)

type UserController struct{}

func (ctrl UserController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ping"})
}

func (ctrl UserController) CreateOrLogin(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	user, err := database.GetUserByExternalID(userCtx.ExternalID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user != nil {
		response.WriteSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	}

	newUser := &account.User{
		ExternalID: userCtx.ExternalID,
		Email:      userCtx.Email,
		Avatar:     userCtx.Avatar,
		Name:       userCtx.Name,
	}

	err = database.CreateUser(newUser)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	response.WriteSuccess(c, http.StatusOK, gin.H{"user": user})

}

func (ctrl UserController) GetUser(c *gin.Context) {
	userCtx, _ := auth.UserFromCtx(c)

	user, err := database.GetUserByExternalID(userCtx.ExternalID)

	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}
	if user != nil {
		response.WriteSuccess(c, http.StatusOK, gin.H{"user": user})
		return
	} else {
		response.WriteError(c, http.StatusNotFound, gin.H{"error": "User not found"})
	}

}
