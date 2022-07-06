package auth

import "github.com/gin-gonic/gin"

type User struct {
	Name       string
	ExternalID string
	Email      string
	Avatar     string
}


