package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.ModRoot}}/mod_{{.module}}/service"
	"{{.ModRoot}}/mod_{{.module}}/models"
)

func GetAll{{.Module | capitalize}}(c *gin.Context) {
	data := service.GetAll{{.Module | capitalize}}()
	c.JSON(http.StatusOK, data)
}

func Create{{.Module | capitalize}}(c *gin.Context) {
	var input models.{{.Module | capitalize}}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.Create{{.Module | capitalize}}(input)
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}
