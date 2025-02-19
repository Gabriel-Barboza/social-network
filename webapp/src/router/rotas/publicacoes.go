package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []Rota{
	{
		URI:        "/publicacoes",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CriarPublicacao,
		Autenticar: true},
	{
		URI:        "/publicacoes/{publicacaoid}/curtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CurtirPublicacao,
		Autenticar: true,
	},
	{
		URI:        "/publicacoes/{publicacaoid}/descurtir",
		Metodo:     http.MethodPost,
		Funcao:     controllers.DescurtirPublicacao,
		Autenticar: true,
	}, {
		URI:        "/publicacoes/{publicacaoid}/editar",
		Metodo:     http.MethodGet,
		Funcao:     controllers.CarregarPaginaDeEdicaoDePublicacao,
		Autenticar: true,
	},
	{
		URI:        "/publicacoes/{publicacaoid}",
		Metodo:     http.MethodPut,
		Funcao:     controllers.AtualizarPublicacao,
		Autenticar: true,
	},
	{
		URI:        "/publicacoes/{publicacaoid}",
		Metodo:     http.MethodDelete,
		Funcao:     controllers.DeletarPublicacao,
		Autenticar: true,
	}}
