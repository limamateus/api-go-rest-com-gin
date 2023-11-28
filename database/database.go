package database

import (
	"api-go-rest-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConexaoComBanco() {
	caminhoDaConexao := "host=localhost user=root password=root dbname=alunodb port=5432 sslmode=disable" // String de Conexão

	DB, err = gorm.Open(postgres.Open(caminhoDaConexao)) // Aqui estou abrindo a conexão com banco

	if err != nil {
		log.Panic("Erro de conexão com banco")
	}

	DB.AutoMigrate(&models.Aluno{})

}
