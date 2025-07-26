package repository

import (
	"{{.ModRoot}}/mod_{{.module}}/models"
	"{{.ModRoot}}/config"
)

func FindAll() []models.{{.Module | capitalize}} {
	var data []models.{{.Module | capitalize}}
	config.DB.Find(&data)
	return data
}

func Save(input models.{{.Module | capitalize}}) {
	config.DB.Create(&input)
}
