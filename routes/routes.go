package routes

import (
	"api-go-rest-gin/controller"

	"github.com/gin-gonic/gin"
)

func HendlerRequest() {
	r := gin.Default()
	r.GET("/", controller.ListaDeAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.POST("/alunos/novo", controller.NovoAluno)
	r.Run()
}
