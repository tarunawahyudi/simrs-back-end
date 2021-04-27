package routes

import (
	"github.com/gofiber/fiber/v2"
	"simrs-backend/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/user", controllers.User)
	api.Post("/logout", controllers.Logout)
}
