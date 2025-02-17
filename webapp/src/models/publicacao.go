package models

import "time"

type Publicacao struct {
	ID        int       `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   int       `json:"autor_id,omitempty"`
	AutorNick string    `json:"autor_nick,omitempty"`
	Curtidas  int       `json:"curtidas"`
	CriadoEm  time.Time `json:"criado_em,omitempty"`
}
