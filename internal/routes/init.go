package routes

import (
	"github.com/gofiber/fiber/v2"
	"news/internal/store"
)

func InitRoutes(store store.InterfaceStore, app *fiber.App) {
	authRoutes(store, app)
	newsRoutes(store, app)
}
