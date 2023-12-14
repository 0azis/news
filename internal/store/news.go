package store

import (
	"gopkg.in/reform.v1"
	"news/internal/models"
)

// Репозиторий новостей
type newsRepository interface {
	GetNews(limit, page int) ([]models.News, error)                          // получить все новости (поддержка пагинации)
	UpdateNews(newsID int, news models.NewsCredentials) (models.News, error) // поменять поля новости
}

type news struct {
	db *reform.DB
}

func (n *news) GetNews(limit, page int) ([]models.News, error) {
	if limit == 0 {
		limit = 5 //default value
	}

	// можно было бы и var allNews []models.News, но хотел чтобы было не nil значение, чтобы возвращался массив, а не null значение после сериализации
	allNews := []models.News{}
	rows, err := n.db.Query("select * from news limit $1 offset $2", limit, page*limit)

	if err != nil {
		return allNews, err
	}

	for rows.Next() {
		var categories []int
		var news models.News
		err = rows.Scan(&news.Title, &news.Content, &news.ID)
		if err != nil {
			return allNews, err
		}

		// не смог придумать лучшего решения как послать второй запрос на получение всех категорий новости и потом собрать их в массив вручную
		// рассматривал варианты с JOIN оператором, но никак не понимал, как в SQL перевести полученные значения в массив
		// знаю что это далеко не лучшее решение, но по другому сделать не могу

		rows2, _ := n.db.Query("select category_id from news_categories where news_id = $1", news.ID)
		for rows2.Next() {
			var category int
			err = rows2.Scan(&category)
			if err != nil {
				return allNews, err
			}
			categories = append(categories, category)
		}

		news.Categories = categories
		allNews = append(allNews, news)
	}

	return allNews, nil
}

func (n *news) UpdateNews(newsID int, news models.NewsCredentials) (models.News, error) {
	var updatedNews models.News

	rows, err := n.db.Query(`update news set title = $1, content = $2 where id = $3 returning *`, news.Title, news.Content, newsID)

	for rows.Next() {
		err = rows.Scan(&updatedNews.Title, &updatedNews.Content, &updatedNews.ID)
		if err != nil {
			return updatedNews, err
		}
	}

	return updatedNews, err
}
