package store

import (
	"database/sql"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"news/internal/models"
	"news/internal/pkg"
)

// Репозиторий новостей
type newsRepository interface {
	GetNews(limit, page int) ([]models.News, error)       // получить все новости (поддержка пагинации)
	UpdateNews(newsID int, news models.News) (int, error) // поменять поля новости
}

type news struct {
	sql *sql.DB
	db  *reform.DB
}

func (n *news) GetNews(limit, page int) ([]models.News, error) {
	if limit == 0 {
		limit = 5 //default value
	}

	// можно было бы и var allNews []models.News, но хотел чтобы было не nil значение, чтобы возвращался массив, а не null значение после сериализации
	allNews := []models.News{}

	rows, err := n.db.Query("select id, title, content, array_agg(category_id) news_categories from news join news_categories on news.id = news_categories.news_id group by id order by id limit $1 offset $2", limit, page*limit)
	if err != nil {
		logrus.Error("Query error getting news")
		return allNews, err
	}
	defer rows.Close()

	for rows.Next() {
		var categories pq.Int64Array
		var news models.News

		err = rows.Scan(&news.ID, &news.Title, &news.Content, &categories)
		if err != nil {
			logrus.Error("Error while collecting news data")
			return allNews, err
		}

		news.Categories = pkg.ConvertToArray(categories)
		allNews = append(allNews, news)
	}

	return allNews, nil
}

func (n *news) UpdateNews(newsID int, news models.News) (int, error) {
	var updatedID int
	rows, err := n.db.Query("delete from news_categories where news_id = $1", newsID)
	defer rows.Close()

	for categoryID := range news.Categories {
		rows, err = n.db.Query("insert into news_categories (news_id, category_id) values ($1, $2)", newsID, news.Categories[categoryID])
		if err != nil {
			logrus.Error("Query error updating news categories")
			return updatedID, err
		}
	}

	row := n.db.QueryRow(`update news set id = $1, title = $2, content = $3 where id = $4 returning id`, news.ID, news.Title, news.Content, newsID)
	row.Scan(&updatedID)

	if err != nil {
		logrus.Error("Query error updating news")
		return updatedID, err
	}

	return updatedID, nil
}
