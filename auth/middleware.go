package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/config"
	"github.com/teezzan/commitspy-v2/database"
	"github.com/teezzan/commitspy-v2/response"
)

func AuthenticateToken(c *gin.Context) {
	firebaseAuthClient := config.GetFirebaseAuthClient()

	authToken := fetchAuthToken(c)

	if authToken == "" {
		response.WriteError(c, http.StatusBadRequest, gin.H{"error": "API token required"})
		return
	}
	var decodedUser User
	if config.Cfg.Env == "TEST" && authToken == "TestToken" {
		decodedUser = TestUserStub
		user, _ := database.GetUserByExternalID(decodedUser.ExternalID)

		if user != nil {
			decodedUser.ID = user.ID
		}
	} else {
		token, err := firebaseAuthClient.VerifyIDToken(context.Background(), authToken)

		if err != nil {
			response.WriteError(c, http.StatusBadRequest, gin.H{"error": "Invalid API token"})
			return
		}
		decodedUser = User{
			ExternalID: token.UID,
			Name:       token.Claims["name"].(string),
			Email:      token.Claims["email"].(string),
			Avatar:     token.Claims["picture"].(string),
		}
		user, _ := database.GetUserByExternalID(token.UID)

		if user != nil {
			decodedUser.ID = user.ID
		}

	}

	c.Set("User", &decodedUser)
	c.Next()
}

func UserFromCtx(c *gin.Context) (*User, bool) {
	userCtxInterface, ok := c.Get("User")
	return userCtxInterface.(*User), ok
}

func fetchAuthToken(c *gin.Context) string {
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	return idToken
}
