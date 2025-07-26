package models

type {{.Module | capitalize}} struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
