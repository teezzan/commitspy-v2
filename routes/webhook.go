package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/commitspy-v2/webhook"
)

var webhookHandler = new(webhook.EventHandlers)

func addWebhookRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/webhooks")

	route.POST("/gh/:uuid", webhook.AuthenticateGithubWebhook, webhookHandler.Github)
}
