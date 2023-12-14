package routes

import (
	"github.com/gofiber/fiber/v2"
	"news/internal/controllers"
	"news/internal/middleware"
	"news/internal/store"
)

func newsRoutes(store store.InterfaceStore, app *fiber.App) {
	controller := controllers.GetNewsControllers(store)

	app.Get("/list", middleware.JWTMiddleware, controller.GetAllNews)
	app.Post("/edit/:Id", middleware.JWTMiddleware, controller.UpdateNews)
}
