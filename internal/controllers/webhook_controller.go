package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/khiemnguyen15/twitch-stream-hunter-api/internal/models"
)

var webhookCollection *mongo.Collection

func CreateWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var webhook models.Webhook
		err := c.BindJSON(&webhook)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		err = validate.Struct(&webhook)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		result, err := webhookCollection.InsertOne(ctx, webhook)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    result,
		})
	}
}

func UpdateWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := c.Param("url")

		var webhook models.Webhook
		err := c.BindJSON(&webhook)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		err = validate.Struct(&webhook)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		update := bson.M{
			"url":   webhook.URL,
			"games": webhook.Games,
		}
		result, err := webhookCollection.UpdateOne(
			ctx,
			bson.M{"url": url},
			bson.M{"$set": update},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		var updatedWebhook models.Webhook
		if result.MatchedCount == 1 {
			err := webhookCollection.FindOne(ctx, bson.M{"url": url}).Decode(&updatedWebhook)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  http.StatusInternalServerError,
					"message": err.Error(),
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    updatedWebhook,
		})
	}
}

func DeleteWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := c.Param("url")

		result, err := webhookCollection.DeleteOne(ctx, bson.M{"url": url})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
		})
	}
}

func GetWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := c.Param("url")
		var webhook models.Webhook

		err := webhookCollection.FindOne(ctx, bson.M{"url": url}).Decode(&webhook)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    webhook,
		})
	}
}

func GetAllWebhooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var webhooks []models.Webhook
		results, err := webhookCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
		defer results.Close(ctx)

		err = results.All(ctx, &webhooks)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    webhooks,
		})
	}
}
