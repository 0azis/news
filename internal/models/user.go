package models

type User struct {
	ID       int    `json:"ID"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Validate валидация полей
func (u User) Validate() bool {
	if u.Login == "" || u.Password == "" {
		return false
	}
	return true
}
