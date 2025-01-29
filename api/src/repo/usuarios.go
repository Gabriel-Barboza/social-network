package repo

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepoUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario models.Usuario) (int, error) {
	statment, err := u.db.Prepare("insert into usuarios (nome, nick, email , senha) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statment.Close()
	resultado, err := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(ultimoIDInserido), nil
}


func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, err := repositorio.db.Query(" select id , nick , email , criadoEm from usuarios where nome like ? or nick like ? ", nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if err = linhas.Scan(&usuario.Id, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

func (repositorio Usuarios) BuscarPorID(id int) (models.Usuario, error) {
	linhas , err := repositorio.db.Query("select id , nome , nick , email , criadoEm from usuarios where id = ? ", id)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()
	var usuario models.Usuario
	if linhas.Next() {
		if err = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); err != nil {
			return models.Usuario{}, err
		}
		if err = linhas.Err(); err != nil {
			return models.Usuario{}, err
		}
	}
	return usuario, nil
}