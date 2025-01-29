package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	Id       int       `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

func (u *Usuario) Preparar(etapa string) error {
	if err := u.Validar(etapa); err != nil {
		return err
	}
	u.Formatar()
	return nil
}

func (u *Usuario) Validar(etapa string) error {
	if u.Nome == "" {
		return errors.New("o campo nome é obrigatório e não pode estar em branco")
	}
	if u.Nick == "" {
		return errors.New("o campo nick é obrigatório e não pode estar em branco")
	}
	if u.Email == "" {
		return errors.New("o campo email é obrigatório e não pode estar em branco")
	}
	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("o campo senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (u *Usuario) Formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
