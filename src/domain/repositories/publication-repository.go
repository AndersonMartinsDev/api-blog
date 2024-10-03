package repositories

import (
	"apiblog/src/domain/entities"
	"apiblog/src/infrastructure/commons/models/page"
)

type PublicationRepository interface {
	Insert(publicaton entities.Publication) error
	PublicationByTitleAndAutor(title, autor string) (entities.Publication, error)
	Update(entities.Publication) error
	DeleteByTitleAndAutor(title, autor string) error
	List(autor, order string, pagination page.Pagination) ([]entities.Publication, error)
}
