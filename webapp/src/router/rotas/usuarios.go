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
}
