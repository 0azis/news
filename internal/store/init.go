package store

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"os"
)

type InterfaceStore interface {
	Open()                  // Открыть подключение
	Close()                 // Закрыть подключение
	News() newsRepository   // Репозиторий новостей
	Users() usersRepository // Репозиторий пользователей
}

type store struct {
	sql *sql.DB
	db  *reform.DB
}

func (s *store) Open() {
	// *sql.DB instance вместе с connection pool
	sql, _ := sql.Open("pgx", os.Getenv("DB_URL"))
	if err := sql.Ping(); err != nil {
		logrus.Error("Database connection failed")
	}

	s.sql = sql

	sql.SetMaxOpenConns(100)
	sql.SetMaxIdleConns(100)

	// ORM instance
	db := reform.NewDB(sql, postgresql.Dialect, reform.NewPrintfLogger(logrus.Printf))

	s.db = db
}

func (s *store) Close() {
	s.sql.Close()
}

func (s *store) News() newsRepository {
	return &news{
		db: s.db,
	}
}

func (s *store) Users() usersRepository {
	return &user{
		db: s.db,
	}
}

// Геттер для получения store instance
func NewStore() InterfaceStore {
	return &store{}
}
