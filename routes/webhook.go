package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/webhook"
)

func addWebhookRoutes(rg *gin.RouterGroup) {
	webhooks := rg.Group("/webhooks")

	webhooks.Use(webhook.AuthenticateGithubWebhook)
	webhooks.POST("/gh/:uuid", user.Ping)
}
