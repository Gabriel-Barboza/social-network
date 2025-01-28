package repo

import (
	"api/src/models"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepoUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario models.Usuario) (int, error) {
	statment , err := u.db.Prepare("insert into usuarios (nome, nick, email , senha) values (?, ?, ?, ?)") 
	if err != nil {
		return 0 , err
	}
	defer statment.Close()
	resultado , err := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha) 
	if err != nil {
		return 0 , err
	}

	ultimoIDInserido , err := resultado.LastInsertId()
	if err != nil {
		return 0 , err
	}
	return int(ultimoIDInserido) , nil
}