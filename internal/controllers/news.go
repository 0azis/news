package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"news/internal/models"
	"news/internal/store"
	"strconv"
)

type newsControllers struct {
	Store store.InterfaceStore
}

func (nc *newsControllers) GetAllNews(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	allNews, err := nc.Store.News().GetNews(limit, page)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	return c.JSON(&models.HttpResponse{
		Success: true,
		News:    allNews,
	})
}

func (nc *newsControllers) UpdateNews(c *fiber.Ctx) error {
	var newsCredentials models.News

	newsID, _ := strconv.Atoi(c.Params("Id"))

	err := c.BodyParser(&newsCredentials)
	if err != nil || !newsCredentials.Validate() {
		return c.Status(http.StatusBadRequest).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	updatedID, err := nc.Store.News().UpdateNews(newsID, newsCredentials)

	if updatedID == 0 {
		return c.Status(http.StatusNotFound).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&models.HttpResponse{
			Success: false,
		})
	}

	return c.JSON(&models.HttpResponse{
		Success: true,
	})
}

// Геттер для получения контроллера новостей
func GetNewsControllers(store store.InterfaceStore) *newsControllers {
	return &newsControllers{
		Store: store,
	}
}
