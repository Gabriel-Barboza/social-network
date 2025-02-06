package repo

import (
	"api/src/models"
	"database/sql"
)

type publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *publicacoes {
	return &publicacoes{db}

}
func (repo publicacoes) Criar(publicacao models.Publicacao) (int, error) {
	statement, err := repo.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}
	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(ultimoIDInserido), nil
}

func (repo publicacoes) BuscarPorID(publicacaoID int64) (models.Publicacao, error) {
	linha, err := repo.db.Query(`
		select p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		where p.id = ?`, publicacaoID)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao models.Publicacao
	if linha.Next() { // se a linha existir
		if err = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}
	return publicacao, nil
}

func (repo publicacoes) Buscar(usuarioID int) ([]models.Publicacao, error) {
	linhas, err := repo.db.Query(`
		select distinct p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		inner join seguidores s on p.autor_id = s.usuario_id
		where u.id = ? or s.seguidor_id = ?`, usuarioID, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao
	for linhas.Next() {
		var p models.Publicacao
		if err = linhas.Scan(
			&p.ID,
			&p.Titulo,
			&p.Conteudo,
			&p.AutorID,
			&p.Curtidas,
			&p.CriadoEm,
			&p.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, p)
	}
	return publicacoes, nil

}

func (repo publicacoes) Atualizar(publicacaoID int64, publicacao models.Publicacao) error {
	statement, err := repo.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}
	return nil

}

func (repo publicacoes) Deletar(publicacaoID int64) error {
	statement, err := repo.db.Prepare("delete from publicacoes where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil
}

func (repo publicacoes) BuscarPorUsuario(usuarioID int64) ([]models.Publicacao, error) {
	linhas, err := repo.db.Query(` select p.*, u.nick from publicacoes p
	inner join usuarios u on u.id = p.autor_id
	where p.autor_id = ?`, usuarioID)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()
	var publicacoes []models.Publicacao
	for linhas.Next() {
		var p models.Publicacao
		if err = linhas.Scan(
			&p.ID,
			&p.Titulo,
			&p.Conteudo,
			&p.AutorID,
			&p.Curtidas,
			&p.CriadoEm,
			&p.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, p)
	}
	return publicacoes, nil
}

func (repo publicacoes) Curtir(publicacaoID int64) error {
	statement, err := repo.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil

}

func (repo publicacoes) Descurtir(publicacaoID int64) error {
	statement, err := repo.db.Prepare(`update publicacoes set curtidas = 
	 CASE when curtidas > 0 THEN 
	 curtidas - 1
	 ELSE 0 END
	 where id = ?`)

	if err != nil {

		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(publicacaoID); err != nil {
		return err
	}
	return nil

}
