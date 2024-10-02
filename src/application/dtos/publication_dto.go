package dto

import (
	"apiblog/src/infrastructure/commons/models/patterns"
)

type PublicationDto struct {
	Title     string        `json:"titulo"`
	Content   string        `json:"conteudo"`
	Autor     string        `json:"autor"`
	OnAirDate patterns.Date `json:"dataPublicacao"`
}
