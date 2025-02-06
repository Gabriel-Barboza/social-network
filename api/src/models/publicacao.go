package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        int       `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   int       `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  int       `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

func (p *Publicacao) Preparar() error {
	if err := p.validar(); err != nil {
		return err
	}
	p.formatar()
	return nil
}

func (p *Publicacao) validar() error {

	if p.Titulo == "" {
		return errors.New("	O titulo é obrigatorio e não pode estar em branco")
	}

	if p.Conteudo == "" {
		return errors.New("	O conteudo é obrigatorio e não pode estar em branco")
	}

	return nil

}

func (p *Publicacao) formatar() {
	p.Titulo = strings.TrimSpace(p.Titulo)
	p.Conteudo = strings.TrimSpace(p.Conteudo)
}
