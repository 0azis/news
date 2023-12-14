package models

// HttpResponse кастомный response
type HttpResponse struct {
	Success bool   `json:"Success"`
	News    []News `json:"News"`
}
