package response

import "github.com/gin-gonic/gin"

func WriteError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "false", "message": message})
}

func WriteSuccess(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"status": "true", "data": message})
}
