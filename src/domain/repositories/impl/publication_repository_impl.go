package impl

import (
	"apiblog/src/domain/entities"
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
		DB:    GetConnectionDatabase(),
		Table: "publication",
	}
}

func (repository PublicationRepositoryImpl) InsertNewPublications(publicaton entities.Publication) error {
	queryBuilder := database.QueryBuild{}

	columns, countValues := queryBuilder.ColumnsAndValuesCount(publicaton)

	insert := fmt.Sprintf("INSERT INTO %s(%s) values(%s)", repository.Table, columns, countValues)

	statement, erro := repository.DB.Prepare(insert)

	statement.Exec(publicaton.Title, publicaton.Content, publicaton.Autor, publicaton.OnAirDate)

	defer statement.Close()

	return erro
}
