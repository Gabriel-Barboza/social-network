package models

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
	if err := u.Formatar(etapa); err != nil {
		return err
	}
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

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o campo email está inválido")
	}

	if etapa == "cadastro" && u.Senha == "" {
		return errors.New("o campo senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (u *Usuario) Formatar(etapa string) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(senhaComHash)

	}
	return nil
}
