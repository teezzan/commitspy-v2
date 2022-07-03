package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/config"
	"github.com/teezzan/commitspy/controllers"
	"github.com/teezzan/commitspy/models"
)

func AuthenticateToken(c *gin.Context) {
	firebaseAuthClient := config.GetFirebaseAuthClient()

	authToken := fetchAuthToken(c)

	if authToken == "" {
		controllers.RespondWithError(c, http.StatusBadRequest, gin.H{"error": "API token required"})
		return
	}

	token, err := firebaseAuthClient.VerifyIDToken(context.Background(), authToken)

	if err != nil {
		controllers.RespondWithError(c, http.StatusBadRequest, gin.H{"error": "Invalid API token"})
		return
	}
	decodedUser := models.ContextualUser{
		ExternalID: token.UID,
		Name:       token.Claims["name"].(string),
		Email:      token.Claims["email"].(string),
		Avatar:     token.Claims["picture"].(string),
	}
	c.Set("User", decodedUser)
	c.Next()
}

func fetchAuthToken(c *gin.Context) string {
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	return idToken
}
