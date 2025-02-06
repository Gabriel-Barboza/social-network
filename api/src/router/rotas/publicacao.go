package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacao = []Rota{

	{
		Url:        "/publicacoes",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarPublicacao,
		Autenticar: true,
	},
	{
		Url:        "/publicacoes",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacoes,
		Autenticar: true,
	},
	{
		Url:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacao,
		Autenticar: true,
	},
	{
		Url:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarPublicacao,
		Autenticar: true,
	},
	{
		Url:        "/publicacoes/{publicacaoId}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarPublicacao,
		Autenticar: true,
	},
	{
		Url:        "/publicacoes/{usuarioId}/publicacoes",
		Metodo:     http.MethodGet,
		Funcao:     controllers.BuscarPublicacoesPorUsuario,
		Autenticar: true,
	}, {
		Url:        "/publicacoes/{publicacaoId}/curtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CurtirPublicacao,
		Autenticar: true,
	}, {
		Url:        "/publicacoes/{publicacaoId}/descurtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.DescurtirPublicacao,
		Autenticar: true,
	},
}
