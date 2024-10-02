package services

import (
	"apiblog/src/domain/entities"
	"apiblog/src/infrastructure/commons/models/page"
)

type PublicationService interface {
	InsertNewPublications(publication entities.Publication) error
	ListPublications(autor, order string, pagination page.Pagination) ([]entities.Publication, error)
}
