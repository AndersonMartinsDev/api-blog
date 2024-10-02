package repositories

import (
	"apiblog/src/domain/entities"
	"apiblog/src/infrastructure/commons/models/page"
)

type PublicationRepository interface {
	InsertNewPublications(publicaton entities.Publication) error
	ListPublications(autor, order string, pagination page.Pagination) ([]entities.Publication, error)
}
