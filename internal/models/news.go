package models

type News struct {
	ID         int    `json:"Id"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	Categories []int  `json:"Categories"`
}

// Validate валидация полей
func (n *News) Validate() bool {
	if n.Title == "" || n.Content == "" {
		return false
	}
	return true
}
