package main

import (
	"api-go-rest-gin/controller"
	"api-go-rest-gin/database"
	"api-go-rest-gin/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // Retorna os dados dos testes resumido
	rotas := gin.Default()

	return rotas
}

func CriarAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno novo", CPF: "12345678901"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DeletarAluno() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
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

func TestListaDeAlunos(t *testing.T) {
	database.ConexaoComBanco()
	CriarAlunoMock()     // Aqui estou criando um aluno Mocado
	defer DeletarAluno() // que será deletado depois de tudo

	r := SetupDasRotasDeTeste() // Instacia do Gin

	r.GET("/", controller.ListaDeAlunos) // A rota que será usada

	req, _ := http.NewRequest("GET", "/", nil) // a requisão que irá ser realizada

	resposta := httptest.NewRecorder() // crio a variavel que irá receber as resposta

	r.ServeHTTP(resposta, req) // realizo a requisão

	assert.Equal(t, http.StatusOK, resposta.Code) // valido se é status code é igual ao que eu esperava

}

func TestBuscaPorCPF(t *testing.T) {
	database.ConexaoComBanco() // Abro a conexão com banco
	CriarAlunoMock()           // Crio um aluno
	defer DeletarAluno()       // Deleto ele no final

	r := SetupDasRotasDeTeste()                             // Crio  a instacia de gin
	r.GET("/alunos/cpf/:cpf", controller.BuscarAlunoPorCpf) // passo a rota que eu irei usar

	req, _ := http.NewRequest("GET", "/alunos/cpf/46063928863", nil) // monsta o request

	resposta := httptest.NewRecorder() // crio a resposta que sera usando para comparar

	r.ServeHTTP(resposta, req) // realizo o request

	assert.Equal(t, http.StatusOK, resposta.Code) // Valido se status code é o mesmo que eu esperava

}
