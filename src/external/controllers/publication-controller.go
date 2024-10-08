package controllers

import (
	dto "apiblog/src/application/dtos"
	"apiblog/src/application/mappers"
	"apiblog/src/domain/services"
	servicesImpl "apiblog/src/domain/services/impl"
	"apiblog/src/infrastructure/commons/models/page"
	"apiblog/src/infrastructure/commons/models/patterns"
	"apiblog/src/infrastructure/request"
	"apiblog/src/infrastructure/response"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
		response.Response(w, http.StatusInternalServerError, nil, erro)
		return
	}

	response.Response(w, http.StatusCreated, "Publicação registrada com Sucesso!", nil)

}

func ListPublications(w http.ResponseWriter, r *http.Request) {

	autor := r.URL.Query().Get("autor")
	order := r.URL.Query().Get("order")
	pageIndex := r.URL.Query().Get("pageIndex")
	pageSize := r.URL.Query().Get("pageSize")

	service = servicesImpl.NewPublicationService()

	pagination := page.Pagination{
		PageIndex: pageIndex,
		PageSize:  pageSize,
	}

	publications, erro := service.ListPublications(autor, order, pagination)

	if erro != nil {
		response.Response(w, http.StatusInternalServerError, nil, erro)
		return
	}

	response.Response(w, http.StatusOK, mapper.ToSlicePublicationsDto(publications), erro)
}

func GetPublicationByTitleAndAutor(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	title := param["title"]
	autor := param["autor"]

	if title == "" || autor == "" {
		err := errors.New("verifique os parametros da requisição")
		response.Response(w, http.StatusBadRequest, nil, err)
		return
	}

	service := servicesImpl.NewPublicationService()

	publication, erro := service.GetPublicationByTitleAndAutor(title, autor)

	if erro != nil {
		response.Response(w, http.StatusInternalServerError, nil, erro)
		return
	}

	response.Response(w, http.StatusOK, mapper.ToPublicationDto(publication), erro)

}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {

	var publicationDto dto.PublicationDto

	if erro := request.Serialization(r.Body, &publicationDto); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	service := servicesImpl.NewPublicationService()

	erro := service.UpdatePublication(mapper.ToPublication(publicationDto))

	if erro != nil {
		response.Response(w, http.StatusInternalServerError, nil, erro)
		return
	}

	response.Response(w, http.StatusOK, "Registro Atualizado Com sucesso!", erro)
}

func DeletePublicationByTitleAndAutor(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	title := param["title"]
	autor := param["autor"]

	if title == "" || autor == "" {
		err := errors.New("verifique os parametros da requisição")
		response.Response(w, http.StatusBadRequest, nil, err)
		return
	}

	service := servicesImpl.NewPublicationService()

	erro := service.DeletePublicationByTitleAndAutor(title, autor)

	if erro != nil {
		response.Response(w, http.StatusInternalServerError, nil, erro)
		return
	}

	response.Response(w, http.StatusOK, "Registro Deletado Com sucesso!", erro)
}
