package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/respostas"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: "Erro ao logar"})
		return
	}
	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: "Erro ao logar"})
		return
	}
	defer response.Body.Close()

	
	token, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode, string(token))
}
