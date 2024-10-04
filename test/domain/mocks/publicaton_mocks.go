package mocks

import (
	"apiblog/src/domain/repositories"
	"apiblog/src/domain/services/impl"
)

func NewPublicationService(mockRepository repositories.PublicationRepository) impl.PublicationServiceImpl {
	return impl.PublicationServiceImpl{
		Repository: mockRepository,
	}
}
