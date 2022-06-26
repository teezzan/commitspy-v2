package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy/config"
)

var firebaseAuth = config.FirebaseAuth

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func AuthenticateToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Ping"})
	authToken := fetchAuthToken(c)

	if authToken == "" {
		respondWithError(c, http.StatusBadRequest, gin.H{"error": "API token required"})
		return
	}

	token, err := firebaseAuth.VerifyIDToken(context.Background(), authToken)

	if err != nil {
		respondWithError(c, http.StatusBadRequest, gin.H{"error": "Invalid API token"})
		return
	}
	c.Set("User", token)
	c.Next()
}

func fetchAuthToken(c *gin.Context) string {
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	return idToken
}
