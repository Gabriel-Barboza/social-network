package rotas

import (
	"api/src/middlewares"
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
	rotas = append(rotas, rotasPublicacao...)

	for _, rota := range rotas {
		if rota.Autenticar{
			r.HandleFunc(rota.Url, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {	
			r.HandleFunc(rota.Url,  middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}
