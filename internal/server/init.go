package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"news/internal/routes"
	"news/internal/store"
)

func InitServer() {
	// Загрузка виртуального окружения
	godotenv.Load()

	// Fiber Instance Application
	app := fiber.New()

	// подключение базы
	storeInstance := store.NewStore()
	storeInstance.Open()
	defer storeInstance.Close()

	// инициализация роутов
	routes.InitRoutes(storeInstance, app)

	// открываем порт и слушаем
	logrus.Printf("Server is running!")
	app.Listen(":5000")
}
