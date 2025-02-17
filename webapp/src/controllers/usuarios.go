package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: "Erro ao criar usuário"})

	}
	url := fmt.Sprintf("%s/usuarios", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: "Erro ao criar usuário"})
	}
	defer response.Body.Close()
	respostas.JSON(w, response.StatusCode, nil)

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
}
