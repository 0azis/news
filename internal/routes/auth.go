package routes

import (
	"github.com/gofiber/fiber/v2"
	"news/internal/controllers"
	"news/internal/store"
)

func authRoutes(store store.InterfaceStore, app *fiber.App) {
	controller := controllers.GetAuthControllers(store)

	auth := app.Group("/auth")
	auth.Post("/signup", controller.SignUp)
	auth.Post("/signin", controller.SignIn)
}
