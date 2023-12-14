package store

import (
	"gopkg.in/reform.v1"
	"news/internal/models"
)

type usersRepository interface {
	InsertOne(user models.User) (int, error)      // Добавить пользователя в базу (SignUp)
	GetByLogin(login string) (models.User, error) // Получить id и password по логину (для SignIn)
}

type user struct {
	db *reform.DB
}

func (u *user) InsertOne(user models.User) (int, error) {
	var insertedID int
	rows, err := u.db.Query("insert into users (login, password) values ($1, $2) returning id", user.Login, user.Password)

	if err != nil {
		return insertedID, err
	}

	for rows.Next() {
		rows.Scan(&insertedID)
	}

	return insertedID, nil
}

func (u *user) GetByLogin(login string) (models.User, error) {
	var resultUser models.User
	rows, err := u.db.Query("select id, password from users where login = $1", login)

	if err != nil {
		return resultUser, err
	}

	for rows.Next() {
		rows.Scan(&resultUser.ID, &resultUser.Password)
	}

	return resultUser, nil
}
