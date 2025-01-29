package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repo"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario models.Usuario
	if err = json.Unmarshal(corpoRequest, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = usuario.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repo.NovoRepoUsuarios(db)

	usuario.Id, err = repo.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuários"))
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuário"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuário"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuário"))
}
