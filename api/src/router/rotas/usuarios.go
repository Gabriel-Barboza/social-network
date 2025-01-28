package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{

	{
		Url:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		Autenticar: false,
	},

	{
		Url:        "/usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuarios,
		Autenticar: false,
	},
	{
		Url:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		Autenticar: false,
	},
	{
		Url:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		Autenticar: false,
	},
	{
		Url:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		Autenticar: false,
	}}
