package routes

import (
	"github.com/gin-gonic/gin"
	"mod_{{.module}}/controllers"
)

func Register{{.Module}}Routes(r *gin.Engine) {
	r.GET("/{{.module}}s", controllers.GetAll{{.Module}})
	r.POST("/{{.module}}s", controllers.Create{{.Module}})
}
