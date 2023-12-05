package controller

import (
	"api-go-rest-gin/database"
	"api-go-rest-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ListaDeAlunos godoc
// @Summary      Lista de Alunos
// @Description  objetivo é retornar uma lista de aluno atraves
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       / [get]
func ListaDeAlunos(c *gin.Context) {
	var alunos []models.Aluno // Aqui estou criando uma variavel que irá representar uma lista de alunos
	database.DB.Find(&alunos) // depois vou no banco e passo os dados para variavel
	c.JSON(200, alunos)       // retorno como json status 200 e a lista.
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

// NovoAluno godoc
//
//	@Summary		Novo Aluno
//	@Description	Objetivo é cadastra um novo Aluno
//	@Tags			Aluno
//	@Accept			json
//	@Produce		json
//	@Param			Aluno	body		models.Aluno	true	"Add Aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/alunos/novo [post]
func NovoAluno(c *gin.Context) {
	var aluno models.Aluno // Aqui eu instacio uma variavel do tipo aluno

	if err := c.ShouldBindJSON(&aluno); err != nil { // aqui estou meio que deserealizando o objeto da requisião para minha struct e valido se tem algum erro e retorno o erro;
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidacaoDoAluno(&aluno); err != nil { // Aqui estou aplicando as validações e retornando caso de error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno) // Aqui estou salvando os dados do aluno no banco

	c.JSON(http.StatusOK, aluno)
}

// AlunoPorId godoc
// @Summary      Buscar um Aluno Por Id
// @Description  objetivo é retornar uma aluno atraves do seu id
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  models.Aluno
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/{id} [get]
func AlunoPorId(c *gin.Context) {
	var aluno models.Aluno // Aqui eu crio uma variavel local que irá representar o aluno

	id := c.Params.ByName("id") // Aqui estou pegando o valor do id passado na rota
	// a ordem é a variavel depois o id
	database.DB.First(&aluno, id) // aqui estou armanzendo o retorno de acordo com a consulta na variavel local
	if aluno.ID == 0 {            // Aqui eu valido se o Id é igual a 0 e retorno bad request com mesagem de erro.
		c.JSON(http.StatusNotFound, gin.H{
			"Erro": "Aluno não econtrando"})
		return

	}

	c.JSON(http.StatusOK, aluno) // Caso de tudo certo retorno o aluno e status ok
}

// DeletarAlunoPorId godoc
// @Summary      Deletar aluno
// @Description  objetivo é deletar um aluno atraves do seu id
// @Tags         Aluno
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      201
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /alunos/{id} [delete]
func DeletarAlunoPorId(c *gin.Context) {
	id := c.Params.ByName("id") // Aqui estou pegando o Id da rota

	var aluno models.Aluno // Criando uma variavel local

	database.DB.First(&aluno, id) // realizo a consulta no banco

	if aluno.ID == 0 { // Se não encontra do bad request e informo que não foi encontrado.
		c.JSON(http.StatusNotFound, gin.H{"Erro": "Aluno não encontrado"})
		return
	}

	database.DB.Delete(&aluno, id) // Deleto do banco

	c.JSON(http.StatusNoContent, nil) // retorno status de 201 no contant

}

// EditarAluno godoc
//
//	@Summary		Atualizar um Aluno
//	@Description	Objetivo é atualizar um Aluno existente
//	@Tags			Aluno
//	@Accept			json
//	@Produce		json
//
// @Param        id   path      int  true  "ID"
//
//	@Param			Aluno	body		models.Aluno	true	"Add Aluno"
//	@Success		200		{object}	models.Aluno
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/alunos/{id} [patch]
func EditarAluno(c *gin.Context) {
	id := c.Params.ByName("id") // Aqui estou pegando o Id da rota

	var aluno models.Aluno // Criando uma variavel local

	database.DB.First(&aluno, id) // realizo a consulta no banco

	if aluno.ID == 0 { // Se não encontra do bad request e informo que não foi encontrado.
		c.JSON(http.StatusNotFound, gin.H{"Erro": "Aluno não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil { // Aqui estou validando objeto da requisição
		c.JSON(http.StatusNotFound, gin.H{"Erro": err.Error()})
		return
	}

	if err := models.ValidacaoDoAluno(&aluno); err != nil { // Aqui estou aplicando as validações e retornando caso de error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno) // Atualizo no bancos

	c.JSON(http.StatusOK, aluno) // retorno status de 200 no contant
}

func BuscarAlunoPorCpf(c *gin.Context) {
	cpf := c.Param("cpf") // Aqui eu pego o valor do cpf

	var aluno models.Aluno // estancio uma variavel local que será usada como referencia

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno) // Realizo uma consulta do banco de acordo com cpf e amarzeno o resultado em aluno

	if aluno.ID == 0 { // Se o aluno não foi encontra do bad request e informo que não foi encontrado.
		c.JSON(http.StatusNotFound, gin.H{"Erro": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno) // retorno status de Ok e o aluno no contant
}
