package services

import "apiblog/src/domain/entities"

type PublicationService interface {
	InsertNewPublications(publication entities.Publication) error
}
