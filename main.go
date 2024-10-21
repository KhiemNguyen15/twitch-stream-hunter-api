package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/khiemnguyen15/twitch-stream-hunter-api/internal/config"
	"github.com/khiemnguyen15/twitch-stream-hunter-api/internal/controllers"
	"github.com/khiemnguyen15/twitch-stream-hunter-api/internal/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Cannot find .env file.")
	}

	mongoUri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DB_NAME")

	config.DB, err = config.ConnectDB(mongoUri, dbName)
	if err != nil {
		log.Fatal(err)
	}

	controllers.LoadCollections()
}

func main() {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	routes.RouteGames(r)
	routes.RouteWebhooks(r)

	r.Run("0.0.0.0:8080")
}
