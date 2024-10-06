package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/khiemnguyen15/twitch-stream-hunter-api/internal/controllers"
)

func RouteWebhooks(r *gin.Engine) {
	r.PUT("/webhooks", controllers.CreateWebhook())
	r.PUT("/webhooks/:url", controllers.UpdateWebhook())
	r.DELETE("/webhooks/:url", controllers.DeleteWebhook())
	r.GET("/webhooks", controllers.GetAllWebhooks())
	r.GET("/webhooks/:url", controllers.GetWebhook())
}
