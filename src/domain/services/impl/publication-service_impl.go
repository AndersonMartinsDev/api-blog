package impl

import (
	"apiblog/src/domain/entities"
	repositories "apiblog/src/domain/repositories"
	"apiblog/src/domain/repositories/impl"
	"apiblog/src/infrastructure/commons/models/page"
)

type PublicationServiceImpl struct {
	Repository repositories.PublicationRepository
}

func NewPublicationService() *PublicationServiceImpl {
	return &PublicationServiceImpl{
		Repository: *impl.NewPublicationRepository(),
	}
}

func (service PublicationServiceImpl) InsertNewPublications(publication entities.Publication) error {
	return service.Repository.Insert(publication)
}

func (service PublicationServiceImpl) ListPublications(autor, order string, pagination page.Pagination) ([]entities.Publication, error) {
	return service.Repository.List(autor, order, pagination)
}

func (service PublicationServiceImpl) GetPublicationByTitleAndAutor(title, autor string) (entities.Publication, error) {
	return service.Repository.PublicationByTitleAndAutor(title, autor)
}

func (service PublicationServiceImpl) UpdatePublication(publication entities.Publication) error {
	return service.Repository.Update(publication)
}

func (service PublicationServiceImpl) DeletePublicationByTitleAndAutor(title, autor string) error {
	return service.Repository.DeleteByTitleAndAutor(title, autor)
}
