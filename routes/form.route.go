package routes

import (
	"github.com/HNG-tasks/stage2/controllers"
	"github.com/gofiber/fiber/v2"
)

func FormRoute(route fiber.Router){
	route.Post("/", controllers.AddForm)
}