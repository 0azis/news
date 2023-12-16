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
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error while loading environment variables")
	}

	// Fiber Instance Application
	app := fiber.New()

	// подключение базы
	storeInstance := store.NewStore()
	storeInstance.Open()
	defer storeInstance.Close()

	// инициализация роутов
	routes.InitRoutes(storeInstance, app)

	// открываем порт и слушаем
	logrus.Traceln("Server is running!")
	err = app.Listen(":5000")
	if err != nil {
		logrus.Panic("Error while starting the server")
	}
}
