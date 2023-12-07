package controller

import (
	"api-go-rest-gin/database"
	"api-go-rest-gin/models"
	"api-go-rest-gin/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NovoUsuario godoc
//
//	@Summary		Novo User
//	@Description	Objetivo é cadastra um novo User
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			User	body		models.User	true	"Add User"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/login/Novo [post]
func NovoUsuario(c *gin.Context) {
	var user models.User // Aqui eu instacio uma variavel do tipo User

	if err := c.ShouldBindJSON(&user); err != nil { // aqui estou meio que deserealizando o objeto da requisião para minha struct e valido se tem algum erro e retorno o erro;
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	database.DB.Create(&user) // Aqui estou salvando os dados do aluno no banco

	c.JSON(http.StatusOK, user)
}

// Login godoc
//
//	@Summary		Novo Login
//	@Description	Objetivo é autenticar
//	@Tags			Login
//	@Accept			json
//	@Produce		json
//	@Param			Login	body		models.Login	true	"Add Login"
//	@Success		200		{object}	models.Login
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/login/ [post]
func Login(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil { // aqui estou meio que deserealizando o objeto da requisião para minha struct e valido se tem algum erro e retorno o erro;
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var user models.User

	database.DB.Where(&models.User{Email: login.Email}).First(&user)
	fmt.Println(&user)

	if user.ID == 0 { // estou validando se usuario existe no banco
		c.JSON(http.StatusNotFound, gin.H{
			"error": "E-mail ou senha invalida "})
		return
	}

	if user.Password != services.SHA256Encoder(login.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Credencias invalida"})
		return
	}

	token, err := services.NewJWTService().GeracaoDeToken(user.ID, user.Nome)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
