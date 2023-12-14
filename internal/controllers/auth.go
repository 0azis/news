package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"news/internal/models"
	"news/internal/pkg"
	"news/internal/store"
)

type authControllers struct {
	Store store.InterfaceStore
}

func (ac *authControllers) SignUp(c *fiber.Ctx) error {
	var credentials models.User

	err := c.BodyParser(&credentials)
	if err != nil || !credentials.Validate() {
		return c.Status(http.StatusBadRequest).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	hashedPassword, err := pkg.Encode([]byte(credentials.Password))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	credentials.Password = string(hashedPassword)

	insertedID, err := ac.Store.Users().InsertOne(credentials)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	token, err := pkg.SignJWT(insertedID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	return c.JSON(&fiber.Map{
		"access_token": token,
	})
}

func (ac *authControllers) SignIn(c *fiber.Ctx) error {
	var credentials models.User

	err := c.BodyParser(&credentials)
	if err != nil || !credentials.Validate() {
		return c.Status(http.StatusBadRequest).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	dbUser, err := ac.Store.Users().GetByLogin(credentials.Login)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	if err := pkg.Decode([]byte(dbUser.Password), []byte(credentials.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	token, err := pkg.SignJWT(dbUser.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	return c.JSON(&fiber.Map{
		"access_token": token,
	})
}

// Геттер для получения контроллера аутентификации
func GetAuthControllers(store store.InterfaceStore) *authControllers {
	return &authControllers{
		Store: store,
	}
}
