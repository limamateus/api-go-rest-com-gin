package main

import (
	"api-go-rest-gin/database"
	"api-go-rest-gin/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando Servidor!")

	database.ConexaoComBanco()

	routes.HendlerRequest()

	fmt.Println("Servidor no ar!")

}
