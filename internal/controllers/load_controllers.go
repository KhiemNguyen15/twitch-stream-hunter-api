package controllers

import (
	"github.com/go-playground/validator/v10"

	"github.com/khiemnguyen15/twitch-stream-hunter-api/internal/config"
)

var validate = validator.New()

func LoadCollections() {
	gameCollection = config.GetCollection(config.DB, "games")
	webhookCollection = config.GetCollection(config.DB, "webhooks")
}
