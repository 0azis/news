package store

import (
	"github.com/sirupsen/logrus"
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
		logrus.Error("Query error creating an user")
		return insertedID, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&insertedID)
		if err != nil {
			logrus.Error("Error while collecting user data")
			return insertedID, err
		}
	}

	return insertedID, nil
}

func (u *user) GetByLogin(login string) (models.User, error) {
	var resultUser models.User

	rows, err := u.db.Query("select * from users where login = $1", login)
	if err != nil {
		logrus.Error("Query error getting an user")
		return resultUser, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&resultUser.ID, &resultUser.Login, &resultUser.Password)
		if err != nil {
			logrus.Error("Error while collecting user data")
			return resultUser, err
		}
	}

	return resultUser, nil
}
