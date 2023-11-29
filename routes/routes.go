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
	r.GET("/alunos/:id", controller.AlunoPorId)
	r.DELETE("/alunos/:id", controller.DeletarAlunoPorId)
	r.PATCH("/alunos/:id", controller.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscarAlunoPorCpf)
	r.Run()
}
