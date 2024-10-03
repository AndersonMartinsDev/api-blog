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

func (repository PublicationRepositoryImpl) Insert(publicaton entities.Publication) error {

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

func (repository PublicationRepositoryImpl) List(autor, order string, pagination page.Pagination) ([]entities.Publication, error) {

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
		if erro := parseSqlToPublication(rows, &publication); erro != nil {
			return []entities.Publication{}, erro
		}
		publications = append(publications, publication)
	}

	defer database.Close()
	defer rows.Close()

	return publications, erro
}

func (repository PublicationRepositoryImpl) PublicationByTitleAndAutor(title, autor string) (entities.Publication, error) {

	sql := database.SQLBuild{}
	database := repository.DB

	query := sql.Select(repository.Table, entities.Publication{})

	query = fmt.Sprintf(`%s WHERE autor = '%s' AND title = '%s'`, query, autor, title)

	result, erro := database.Query(query)

	if erro != nil {
		return entities.Publication{}, erro
	}

	var entity entities.Publication

	if result.Next() {
		if erro := parseSqlToPublication(result, &entity); erro != nil {
			return entities.Publication{}, erro
		}
	}

	defer result.Close()

	return entity, erro
}

func (repository PublicationRepositoryImpl) Update(publication entities.Publication) error {

	update := fmt.Sprintf(`UPDATE %s SET content='%s' WHERE autor='%s' and title='%s'`,
		repository.Table,
		publication.Content,
		publication.Autor,
		publication.Title)

	database := repository.DB

	statement, erro := database.Prepare(update)

	if erro != nil {
		return erro
	}

	_, erro = statement.Exec()

	defer database.Close()

	return erro
}

func (repository PublicationRepositoryImpl) DeleteByTitleAndAutor(title, autor string) error {

	deleteSQL := fmt.Sprintf(`DELETE FROM %s`, repository.Table)

	deleteSQL = fmt.Sprintf(`%s WHERE autor = '%s' AND title = '%s'`, deleteSQL, autor, title)

	database := repository.DB

	statement, erro := database.Prepare(deleteSQL)

	if erro != nil {
		return erro
	}

	_, erro = statement.Exec()

	defer database.Close()

	return erro
}

func parseSqlToPublication(rows *sql.Rows, entity *entities.Publication) error {
	if erro := rows.Scan(
		&entity.Id,
		&entity.Title,
		&entity.Content,
		&entity.Autor,
		&entity.OnAirDate.Time,
	); erro != nil {
		return erro
	}
	return nil
}
