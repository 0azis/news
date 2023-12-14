package routes

import (
	"github.com/gofiber/fiber/v2"
	"news/internal/store"
)

func InitRoutes(store store.InterfaceStore, app *fiber.App) {
	newsRoutes(store, app)
	authRoutes(store, app)
}
