package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{

	{
		URI:        "/criar-usuario",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarTelaDeCadastroDeUsuario,
		Autenticar: false,
	},
	{
		URI:        "/usuarios",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarUsuario,
		Autenticar: false,
	},
	{
		URI:        "/buscar-usuarios",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeUsuarios,
		Autenticar: true,
	},
	{
		URI:        "/usuarios/{usuarioId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPerfilDoUsuario,
		Autenticar: true,
	},
	{
		URI:        "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.PararDeSeguirUsuario,
		Autenticar: true,
	},
	{
		URI:        "/usuarios/{usuarioId}/seguir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.SeguirUsuario,
		Autenticar: true,
	},
	{
		URI:        "/perfil",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPerfilDoUsuarioLogado,
		Autenticar: true,
	},
	{
		URI:        "/editar-usuario",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeEdicaoDeUsuario,
		Autenticar: true,
	},
	{
		URI:        "/editar-usuario",
		Metodo:     http.MethodPut,
		Funcao:     controllers.EditarUsuario,
		Autenticar: true,
	},

	{
		URI:        "/atualizar-senha",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeAtualizacaoDeSenha,
		Autenticar: true,
	},
	{
		URI:        "/atualizar-senha",
		Metodo:     http.MethodPost,
		Funcao:     controllers.AtualizarSenha,
		Autenticar: true,
	},
	{
		URI:        "/deletar-usuario",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarUsuario,
		Autenticar: true,
	},
}

