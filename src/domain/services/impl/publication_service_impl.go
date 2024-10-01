package impl

import (
	"apiblog/src/domain/entities"
	"apiblog/src/domain/repositories"
	"apiblog/src/domain/repositories/impl"
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
