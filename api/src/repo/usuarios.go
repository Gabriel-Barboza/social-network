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
	linhas, err := repositorio.db.Query("select id , nome , nick , email , criadoEm from usuarios where id = ? ", id)
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

func (repositorio Usuarios) Atualizar(id int, usuario models.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ? , nick = ? , email = ? where id = ? ")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); err != nil {
		return err
	}
	return nil

}

func (repositorio Usuarios) Deletar(id int) error {
	statement, err := repositorio.db.Prepare("delete  from usuarios where id = ? ")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {

	linha, err := repositorio.db.Query("select id , senha from usuarios where email = ? ", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linha.Close()
	var usuario models.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Id, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}

	}
	return usuario, nil
}

func (repositorio Usuarios) Seguir(usuarioID, seguidorID int) error {
	statement, err := repositorio.db.Prepare("insert ignore into seguidores(usuario_id, seguidor_id) values(?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(usuarioID, seguidorID); err != nil {
		return err
	}
	return nil
}
