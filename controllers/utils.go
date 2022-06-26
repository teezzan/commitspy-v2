package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/models"
	"gorm.io/gorm"
)

func GetCtxUser(c *gin.Context) models.ContextualUser {
	userCtxInterface, _ := c.Get("User")
	return userCtxInterface.(models.ContextualUser)
}

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "false", "message": message})
}

func RespondWithSuccess(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "true", "data": message})
}

func GetUserDetails(ExternalID string) (models.User, *gorm.DB) {
	var user models.User
	result := db.Where(&models.User{ExternalID: ExternalID}).Limit(1).Find(&user)
	return user, result
}

func CreateUser(userDetails models.ContextualUser) (models.User, *gorm.DB) {
	newUser := models.User{
		ExternalID: userDetails.ExternalID,
		Email:      userDetails.Email,
		Avatar:     userDetails.Avatar,
		Name:       userDetails.Name,
	}
	result := db.Create(&newUser)
	return newUser, result
}
