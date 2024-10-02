package impl

import (
	"apiblog/src/domain/entities"
	"apiblog/src/domain/repositories"
	"apiblog/src/domain/repositories/impl"
	"apiblog/src/infrastructure/commons/models/page"
)

type PublicationServiceImpl struct {
	Repository repositories.PublicationRepository
}

func NewPublicationService() *PublicationServiceImpl {
	return &PublicationServiceImpl{
		Repository: impl.NewPublicationRepository(),
	}
}

func (service PublicationServiceImpl) InsertNewPublications(publication entities.Publication) error {
	return service.Repository.InsertNewPublications(publication)
}

func (service PublicationServiceImpl) ListPublications(autor, order string, pagination page.Pagination) ([]entities.Publication, error) {
	return service.Repository.ListPublications(autor, order, pagination)
}
