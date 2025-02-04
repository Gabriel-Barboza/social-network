package controllers

import (
	"api/src/auth"
	"api/src/banco"
	"api/src/models"
	"api/src/repo"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequisicao, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer bd.Close()
	repositorio := repo.NovoRepoUsuarios(bd)
	usuarioSalvoNoBanco, err := repositorio.BuscarPorEmail(usuario.Email)

	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if err = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	token, err := auth.CriarToken(usuarioSalvoNoBanco.Id)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	fmt	.Println(token)
	w.Write([]byte(token))
}
