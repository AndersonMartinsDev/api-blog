package controllers

import (
	dto "apiblog/src/application/dtos"
	"apiblog/src/application/mappers"
	"apiblog/src/domain/services"
	servicesImpl "apiblog/src/domain/services/impl"
	"apiblog/src/infrastructure/patterns"
	"apiblog/src/infrastructure/request"
	"apiblog/src/infrastructure/response"
	"fmt"
	"net/http"
	"time"
)

var (
	mapper  = mappers.PublicationMapper{}
	service services.PublicationService
)

func InsertNewPublications(w http.ResponseWriter, r *http.Request) {

	var publicationDto dto.PublicationDto

	if erro := request.Serialization(r.Body, &publicationDto); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	publicationDto.OnAirDate = patterns.Date{Time: time.Now()}

	publication := mapper.ToPublication(publicationDto)

	service = servicesImpl.NewPublicationService()

	if erro := service.InsertNewPublications(publication); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, "Publicação registrada")

}

func ListPublications(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nova publicação cadastrada")
}
