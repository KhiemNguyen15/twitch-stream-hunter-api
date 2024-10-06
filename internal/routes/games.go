package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/khiemnguyen15/twitch-stream-hunter-api/internal/controllers"
)

func RouteGames(r *gin.Engine) {
	r.PUT("/games", controllers.CreateGame())
	r.DELETE("/games/:id", controllers.DeleteGame())
	r.GET("/games", controllers.GetAllGames())
	r.GET("/games/:id", controllers.GetGame())
}
