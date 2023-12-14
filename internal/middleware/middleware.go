package middleware

import (
	"github.com/gofiber/fiber/v2"
	"news/internal/models"
	"news/internal/pkg"
	"strings"
	"time"
)

// Middleware для проверки наличия токена и его валидности
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("authorization")

	if authHeader == "" || authHeader == "Bearer undefined" {
		return c.Status(401).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	token := strings.SplitN(authHeader, " ", 2)[1]
	_, expiredTime, err := pkg.GetIdentity(token)

	if err != nil {
		return c.Status(401).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	if int64(expiredTime) < time.Now().Unix() {
		return c.Status(401).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	return c.Next()
}
