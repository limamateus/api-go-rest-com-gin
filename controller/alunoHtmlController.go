package controller

import (
	"api-go-rest-gin/database"
	"api-go-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibePaginaIndex(c *gin.Context) { // Função que irá retorna uma mensagem de Bem vindo em Html
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mensagem": "Bem Vindo",
	})
}

func ExibeListaDeAlunosEmHtml(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "listaDeAlunos.html", alunos)
}

func RotasNaoEncontradas(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
