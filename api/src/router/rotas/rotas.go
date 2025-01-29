package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	Url        string
	Metodo     string
	Funcao     func(http.ResponseWriter, *http.Request)
	Autenticar bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	for _, rota := range rotas {
		r.HandleFunc(rota.Url, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
