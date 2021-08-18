package main

import (
	"log"
	"os"

	"github.com/HNG-tasks/stage2/config"
	"github.com/HNG-tasks/stage2/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the root endpoint 😉",
		})
	})

	api := app.Group("/api")

	routes.FormRoute(api.Group("/form"))
}

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public") // set path to static files

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()

	app.Get("/", index)
	setupRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal("Error app failed to start")
		panic(err)
	}
}

func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
