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
		Url:        "/usuarios/{usuarioId}/seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		Autenticar: true,
	},
	{
		Url:        "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.PararDeSeguirUsuario,
		Autenticar: true,
	},
	{
		Url:        "/usuarios/{usuarioId}/seguidores",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguidores,
		Autenticar: true,
	},
	{
		Url:        "/usuarios/{usuarioId}/seguindo",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarSeguindo,
		Autenticar: true,
	}, {
		Url:        "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:     http.MethodPost,
		Funcao:     controllers.AtualizarSenha,
		Autenticar: true,
	}}
