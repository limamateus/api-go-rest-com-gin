package routes

import (
	"api-go-rest-gin/controller"

	"github.com/gin-gonic/gin"
)

func HendlerRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")   // Aqui estou informando ao gin onde ele irá encontrar as paginas HTML
	r.Static("/assets", "./assets") // Aqui estou definindo e configurando onde o gin irá pegar o arquivo de css e passa nas paginas
	// Rotas de Api
	r.GET("/", controller.ListaDeAlunos)
	r.GET("/:nome", controller.Saudacao)
	r.POST("/alunos/novo", controller.NovoAluno)
	r.GET("/alunos/:id", controller.AlunoPorId)
	r.DELETE("/alunos/:id", controller.DeletarAlunoPorId)
	r.PATCH("/alunos/:id", controller.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscarAlunoPorCpf)
	// Rotas que irá exibir paginas em Html com Gin
	r.GET("/index", controller.ExibePaginaIndex)
	r.GET("/listaDeAlunos", controller.ExibeListaDeAlunosEmHtml)

	r.NoRoute(controller.RotasNaoEncontradas)
	r.Run()
}
