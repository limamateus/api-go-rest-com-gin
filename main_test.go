package main

import (
	"api-go-rest-gin/controller"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()

	return rotas
}

func TestVerificaSattusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste() // uso a instacia do Gin

	r.GET("/:nome", controller.Saudacao) // Defino a rota

	req, _ := http.NewRequest("GET", "/Mateus", nil) // monta o request

	resposta := httptest.NewRecorder() // crio a variavel que irá receber as resposta

	r.ServeHTTP(resposta, req) // Realizo a requisão

	//if resposta.Code != http.StatusOK { // Realizo a validação, caso não de 200 ok, eu mostro o status code.
	//		t.Fatalf("Status error: valor recebido foi %d e o esperado é %d", resposta.Code, http.StatusOK)
	//}

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais") // Aqui estou usando o assert para comparar a requisição

	mockDaResposta := `{"API diz:":"E ai Mateus, tudo beleza?"}` // Aqui estou criando um mock

	respostaBody, _ := ioutil.ReadAll(resposta.Body) // estou passando uma leitura do body da requisão

	assert.Equal(t, mockDaResposta, string(respostaBody), "Deveriam ser iguais") // e comparando ela

}
