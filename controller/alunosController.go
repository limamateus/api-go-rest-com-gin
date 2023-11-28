package controller

import (
	"api-go-rest-gin/database"
	"api-go-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListaDeAlunos(c *gin.Context) {

	c.JSON(200, gin.H{})
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func NovoAluno(c *gin.Context) {
	var aluno models.Aluno // Aqui eu instacio uma variavel do tipo aluno

	if err := c.ShouldBindJSON(&aluno); err != nil { // aqui estou meio que deserealizando o objeto da requisião para minha struct e valido se tem algum erro e retorno o erro;
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno) // Aqui estou salvando os dados do aluno no banco

	c.JSON(http.StatusOK, aluno)
}
