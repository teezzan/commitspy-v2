package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/webhook"
)

var webhookHandler = new(webhook.EventHandlers)

func addWebhookRoutes(rg *gin.RouterGroup) {
	wroute := rg.Group("/webhooks")

	wroute.Use(webhook.AuthenticateGithubWebhook)
	wroute.POST("/gh/:uuid", webhookHandler.Github)
}
