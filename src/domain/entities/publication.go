package entities

import (
	"apiblog/src/infrastructure/commons/models/patterns"
)

type Publication struct {
	Id        uint          `column:"id"`
	Title     string        `column:"title"`
	Content   string        `column:"content"`
	Autor     string        `column:"autor"`
	OnAirDate patterns.Date `column:"onairdate"`
}
