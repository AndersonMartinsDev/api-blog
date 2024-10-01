package repositories

import "apiblog/src/domain/entities"

type PublicationRepository interface {
	InsertNewPublications(publicaton entities.Publication) error
}
