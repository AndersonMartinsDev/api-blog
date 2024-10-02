package impl

import (
	"apiblog/src/domain/entities"
	page "apiblog/src/infrastructure/commons/models/page"
	"apiblog/src/infrastructure/database"
	"database/sql"
	"fmt"
)

type PublicationRepositoryImpl struct {
	*sql.DB
	Table string
}

func NewPublicationRepository() *PublicationRepositoryImpl {
	return &PublicationRepositoryImpl{
		DB:    database.GetConnectionDatabase(),
		Table: "publication",
	}
}

func (repository PublicationRepositoryImpl) InsertNewPublications(publicaton entities.Publication) error {

	sql := database.SQLBuild{}
	database := repository.DB

	insert := sql.Insert(repository.Table, publicaton)

	statement, err := database.Prepare(insert)

	if err != nil {
		return err
	}

	_, erro := statement.Exec(publicaton.Title, publicaton.Content, publicaton.Autor, publicaton.OnAirDate.Time)

	defer statement.Close()
	defer database.Close()

	return erro
}

func (repository PublicationRepositoryImpl) ListPublications(autor, order string, pagination page.Pagination) ([]entities.Publication, error) {

	sql := database.SQLBuild{}
	database := repository.DB

	query := sql.Select(repository.Table, entities.Publication{})

	if autor != "" {
		query = fmt.Sprintf(`%s WHERE autor = '%s'`, query, autor)
	}

	if order != "" {
		query = fmt.Sprintf(`%s ORDER BY onairdate %s`, query, order)
	}

	query = sql.Pagination(query, pagination.PageIndex, pagination.PageSize)

	rows, erro := database.Query(query)

	if erro != nil {
		return []entities.Publication{}, erro
	}

	var publications []entities.Publication

	for rows.Next() {
		var publication entities.Publication

		if err := rows.Scan(
			&publication.Title,
			&publication.Content,
			&publication.Autor,
			&publication.OnAirDate.Time,
		); err != nil {
			return []entities.Publication{}, err
		}
		publications = append(publications, publication)
	}

	defer database.Close()
	defer rows.Close()

	return publications, erro
}
