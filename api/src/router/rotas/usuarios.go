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
		Autenticar: true,
	},
	{
		Url:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarUsuario,
		Autenticar: true,
	},
	{
		Url:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarUsuario,
		Autenticar: true,
	},
	{
		Url:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		Autenticar: true,
	},

	{
		Url:        "/usuarios/{usuarioId}/segiuir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		Autenticar: true,
	}}
