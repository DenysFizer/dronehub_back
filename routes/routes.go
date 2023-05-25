package routes

import (
	"dron_hub_back_/Controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", Controller.Register)
	app.Post("/api/login", Controller.Login)
	app.Get("/api/users", Controller.Users)
	app.Get("/api/logout", Controller.Logout)
	app.Use(Controller.AuthMiddleware)
	app.Get("/api/drones", Controller.Drones)
	app.Get("/api/drones/:id", Controller.IdDrones)
}
