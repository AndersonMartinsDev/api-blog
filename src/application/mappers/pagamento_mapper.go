package mappers

import (
	dto "apiblog/src/application/dtos"
	"apiblog/src/domain/entities"
)

type PublicationMapper struct{}

func (mapper PublicationMapper) ToPublication(dataObject dto.PublicationDto) entities.Publication {
	var entity entities.Publication
	entity.Content = dataObject.Content
	entity.Autor = dataObject.Autor
	entity.Title = dataObject.Title
	entity.OnAirDate = dataObject.OnAirDate
	return entity
}
