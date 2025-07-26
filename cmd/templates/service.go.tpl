package service

import (
	"{{.ModRoot}}/mod_{{.module}}/models"
	"{{.ModRoot}}/mod_{{.module}}/repository"
)

func GetAll{{.Module | capitalize}}() []models.{{.Module | capitalize}} {
	return repository.FindAll()
}

func Create{{.Module | capitalize}}(data models.{{.Module | capitalize}}) {
	repository.Save(data)
}
