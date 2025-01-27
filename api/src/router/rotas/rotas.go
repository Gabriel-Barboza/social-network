package rotas

import "net/http"

type Rota struct {
	Url        string
	Metodo     string
	Funcao     func(http.ResponseWriter, *http.Request)
	Autenticar bool
}
