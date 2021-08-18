package controllers

import (
	"context"
	"log"

	// "math"
	// "strconv"
	"time"

	"github.com/HNG-tasks/stage2/config"
	"github.com/HNG-tasks/stage2/models"
	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func AddForm(c *fiber.Ctx) error {
	formCollection := config.MI.DB.Collection("forms")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	form := new(models.Form)

	if err := c.BodyParser(form); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	result, err := formCollection.InsertOne(ctx, form)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Form failed to insert",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    result,
		"success": true,
		"message": "Form inserted successfully!",
	})
}
