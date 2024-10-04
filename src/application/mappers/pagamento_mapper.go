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

func (mapper PublicationMapper) ToPublicationDto(entity entities.Publication) dto.PublicationDto {
	var publicationDto dto.PublicationDto
	publicationDto.Content = entity.Content
	publicationDto.Autor = entity.Autor
	publicationDto.Title = entity.Title
	publicationDto.OnAirDate = entity.OnAirDate
	return publicationDto
}

func (mapper PublicationMapper) ToSlicePublicationsDto(slice []entities.Publication) []dto.PublicationDto {
	var publications []dto.PublicationDto
	for _, entity := range slice {
		publications = append(publications, mapper.ToPublicationDto(entity))
	}
	return publications
}
