package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost,
		url, bytes.NewBuffer(publicacao))

	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)

}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseInt(parametros["publicacaoid"], 10, 64)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/publicacoes/%d/curtir", config.APIURL, publicacaoID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseInt(parametros["publicacaoid"], 10, 64)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/publicacoes/%d/descurtir", config.APIURL, publicacaoID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)
}

func atualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, err := strconv.ParseInt(parametros["publicacaoid"], 10, 64)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: err.Error()})
		return
	}
	r.ParseForm()

	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Erro: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID)
	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut,
		url, bytes.NewBuffer(publicacao))

	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Erro: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)

}
