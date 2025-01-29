package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaLogin = Rota{
	Url:        "/login",
	Metodo:     http.MethodPost,
	Funcao:     controllers.Login,
	Autenticar: false}
