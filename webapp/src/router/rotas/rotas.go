package rotas

import (
	"net/http"
	"webapp/middlewares"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI        string
	Metodo     string
	Funcao     func(w http.ResponseWriter, r *http.Request)
	Autenticar bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)

	for _, rota := range rotas {
		if rota.Autenticar {
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return router
}
