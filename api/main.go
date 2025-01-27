package main

import (
	"api/api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
