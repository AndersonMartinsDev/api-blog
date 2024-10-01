package dto

import (
	date "apiblog/src/infrastructure/patterns"
)

type PublicationDto struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Autor     string    `json:"autor"`
	OnAirDate date.Date `json:"onAirDate"`
}
