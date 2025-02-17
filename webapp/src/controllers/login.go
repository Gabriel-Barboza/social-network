package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
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

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: "Erro ao logar"})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	fmt.Println(response.StatusCode)

	if response.StatusCode != http.StatusOK {
		respostas.JSON(w, response.StatusCode, respostas.ErroApi{Erro: "Usuário ou senha inválidos"})
		return
	}

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	var dadosAuth models.DadosAuth
	if err = json.NewDecoder(response.Body).Decode(&dadosAuth); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: "Erro ao logar"})
		return
	}
	if err = cookies.Salvar(w, dadosAuth.ID, dadosAuth.Token); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroApi{Erro: "Erro ao salvar cookies"})
		return
	}
	respostas.JSON(w, http.StatusOK, nil)
}
