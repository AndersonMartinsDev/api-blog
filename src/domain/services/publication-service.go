package services

import (
	"apiblog/src/domain/entities"
	"apiblog/src/infrastructure/commons/models/page"
)

type PublicationService interface {
	InsertNewPublications(publicaton entities.Publication) error
	GetPublicationByTitleAndAutor(title, autor string) (entities.Publication, error)
	UpdatePublication(entities.Publication) error
	DeletePublicationByTitleAndAutor(title, autor string) error
	ListPublications(autor, order string, pagination page.Pagination) ([]entities.Publication, error)
}
