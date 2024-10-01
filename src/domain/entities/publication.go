package entities

import (
	date "apiblog/src/infrastructure/patterns"
)

type Publication struct {
	Id        uint
	Title     string    `column:"title"`
	Content   string    `column:"content"`
	Autor     string    `column:"autor"`
	OnAirDate date.Date `column:"onAirDate"`
}
