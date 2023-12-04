package main

import (
	"api-go-rest-gin/controller"
	"api-go-rest-gin/database"
	"api-go-rest-gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	//	gin.SetMode(gin.ReleaseMode) // Retorna os dados dos testes resumido
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

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConexaoComBanco()                                                 // 1 - Inicio uma comunicação com banco
	CriarAlunoMock()                                                           // 2- Crio um Aluno atraves de uma função mocada
	defer DeletarAluno()                                                       // 3- Deleto ele depois que tudo for finalizado
	r := SetupDasRotasDeTeste()                                                // 4 -Inicio uma instacia do Gin
	r.GET("/alunos/:id", controller.AlunoPorId)                                // 5- Fala a rota que sera usada
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)                               // 6- Crio a rota que será usada
	req, _ := http.NewRequest("GET", pathDaBusca, nil)                         // 7- Monto a resposta
	resposta := httptest.NewRecorder()                                         // 8- Crio uma variavel que irá armazena uma resposta
	r.ServeHTTP(resposta, req)                                                 // 9 - Realizo uma requisição
	var alunoMock models.Aluno                                                 // 10 -  criar uma variavel local que irá receber o aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)                          // 11- converto a resposta em json e deserealizo na variavel local
	assert.Equal(t, "Aluno novo", alunoMock.Nome, "Os nomes devem ser iguais") //Realizo as validações
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletarAlunoPorId(t *testing.T) {
	database.ConexaoComBanco() // 1 - Inicio uma comunicação com banco
	CriarAlunoMock()           // 2- Crio um Aluno atraves de uma função mocada

	r := SetupDasRotasDeTeste()                           // 3 -Inicio uma instacia do Gin
	r.DELETE("/alunos/:id", controller.DeletarAlunoPorId) // 4- Fala a rota que sera usada
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)          // 5- Crio a rota que será usada
	req, _ := http.NewRequest("DELETE", pathDaBusca, nil) // 6- Monto a resposta
	resposta := httptest.NewRecorder()                    // 7- Crio uma variavel que irá armazena uma resposta
	r.ServeHTTP(resposta, req)                            // 8 - Realizo uma requisição
	assert.Equal(t, http.StatusNoContent, resposta.Code)  // Realizo a validação do status code

}

func TestAtualizarAluno(t *testing.T) {
	database.ConexaoComBanco()                                    // 1 - Inicio uma comunicação com banco
	CriarAlunoMock()                                              // 2- Crio um Aluno atraves de uma função mocada
	defer DeletarAluno()                                          // 3- Deleto ele depois que tudo for finalizado
	r := SetupDasRotasDeTeste()                                   // 4 -Inicio uma instacia do Gin
	r.PATCH("/alunos/:id", controller.EditarAluno)                // 5- Fala a rota que sera usada
	aluno := models.Aluno{Nome: "Atualizado", CPF: "41123456789"} // 6- Crio uma variavel local que será usada para atualizar o aluno
	valorJson, _ := json.Marshal(aluno)                           // 7- Converto o aluno para um json
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)                  // 8- Monto a rota
	fmt.Println(ID)
	req, _ := http.NewRequest("PATCH", pathDaBusca, bytes.NewBuffer(valorJson)) // 9- Monto a requesição
	resposta := httptest.NewRecorder()                                          // 10-  Crio a variavel local que será usada para armazenar a respota da requisição
	r.ServeHTTP(resposta, req)                                                  // 11- Realizo a requisição
	var alunoMockAtualizado models.Aluno                                        // 12- crio um variavel local do tipo aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)                 // 13- Converto em Json
	assert.Equal(t, "Atualizado", alunoMockAtualizado.Nome)                     // Realizo as validações
	assert.Equal(t, "41123456789", alunoMockAtualizado.CPF)

	fmt.Println(alunoMockAtualizado)

	assert.Equal(t, http.StatusOK, resposta.Code)
}
