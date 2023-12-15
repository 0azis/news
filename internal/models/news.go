package models

type News struct {
	ID         int    `json:"Id"`
	Title      string `json:"Title"`
	Content    string `json:"Content"`
	Categories []int  `json:"Categories"`
}

type NewsCredentials struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Validate валидация полей
func (n *News) Validate() bool {
	if n.Title == "" || n.Content == "" {
		return false
	}
	return true
}
