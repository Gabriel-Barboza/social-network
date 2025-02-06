package controllers

import (
	"api/src/auth"
	"api/src/banco"
	"api/src/models"
	"api/src/repo"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {

	usuarioID, err := auth.ExtrairUsuarioId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var publicacao models.Publicacao
	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	publicacao.AutorID = usuarioID

	if err = publicacao.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repo := repo.NovoRepositorioDePublicacoes(db)
	publicacao.ID, err = repo.Criar(publicacao)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusCreated, publicacao)

}

func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repo.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repo.Buscar(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, publicacoes)

}

func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseInt(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repo.NovoRepositorioDePublicacoes(db)
	publicacao, err := repo.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, publicacao)

}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseInt(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repo.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, err := repo.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("you are not allowed to update this publication"))
		return
	}
	corpoRequisicao, err := io.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var publicacao models.Publicacao
	if err = json.Unmarshal(corpoRequisicao, &publicacao); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = publicacao.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = repo.Atualizar(publicacaoID, publicacao); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, err := auth.ExtrairUsuarioId(r)
	if err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseInt(parametros["publicacaoId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repo.NovoRepositorioDePublicacoes(db)
	publicacaoSalvaNoBanco, err := repo.BuscarPorID(publicacaoID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if publicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("you are not allowed to update this publication"))
		return
	}

	if err = repo.Deletar(publicacaoID); err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	
	respostas.JSON(w, http.StatusNoContent, nil)
}

func BuscarPublicacoesPorUsuario (   w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, err := strconv.ParseInt(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repo.NovoRepositorioDePublicacoes(db)
	publicacoes, err := repo.BuscarPorUsuario(usuarioID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, publicacoes)
}