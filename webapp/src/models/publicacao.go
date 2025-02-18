package models

import "time"

type Publicacao struct {
	ID        int       `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   int       `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  int       `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}
