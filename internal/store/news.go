package store

import (
	"github.com/lib/pq"
	"gopkg.in/reform.v1"
	"news/internal/models"
	"news/internal/pkg"
)

// Репозиторий новостей
type newsRepository interface {
	GetNews(limit, page int) ([]models.News, error)               // получить все новости (поддержка пагинации)
	UpdateNews(newsID int, news models.News) (models.News, error) // поменять поля новости
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

	rows, err := n.db.Query("select id, title, content, array_agg(category_id) news_categories from news join news_categories on news.id = news_categories.news_id group by id order by id limit $1 offset $2", limit, page*limit)

	if err != nil {
		return allNews, err
	}

	for rows.Next() {
		var categories pq.Int64Array
		var news models.News

		err = rows.Scan(&news.ID, &news.Title, &news.Content, &categories)
		if err != nil {
			return allNews, err
		}
		news.Categories = pkg.ConvertToArray(categories)
		allNews = append(allNews, news)
	}

	return allNews, nil
}

func (n *news) UpdateNews(newsID int, news models.News) (models.News, error) {
	var updatedNews models.News

	rows, err := n.db.Query(`update news set title = $1, content = $2 where id = $3 returning *`, news.Title, news.Content, newsID)

	for rows.Next() {
		err = rows.Scan(&updatedNews.ID, &updatedNews.Title, &updatedNews.Content)
		if err != nil {
			return updatedNews, err
		}
	}

	return updatedNews, err
}
