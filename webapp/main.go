package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Carregar()
	cookies.Configurar()

	fmt.Println(config.HashKey)

	utils.CarregarTemplates()
	r := router.Gerar()
	fmt.Printf("escutando na porta %d \n", config.Porta)
	log.Fatal(http.ListenAndServe(":3000", r))
}
