package routes

import (
	"apiblog/src/external/controllers"
	model "apiblog/src/infrastructure/commons/models/routes"
	"net/http"
)

var requestMapping = "/publication"

var PublicationRoutes = []model.RouteModel{
	{
		URI:          requestMapping + "/insert",
		Metodo:       http.MethodPost,
		Autenticacao: false,
		Func:         controllers.InsertNewPublications,
	},
	{
		URI:          requestMapping,
		Metodo:       http.MethodGet,
		Autenticacao: false,
		Func:         controllers.ListPublications,
	},
	{
		URI:          requestMapping + "/{title}/{autor}",
		Metodo:       http.MethodGet,
		Autenticacao: false,
		Func:         controllers.GetPublicationByTitleAndAutor,
	},
	{
		URI:          requestMapping,
		Metodo:       http.MethodPut,
		Autenticacao: false,
		Func:         controllers.UpdatePublication,
	},
	{
		URI:          requestMapping + "/{title}/{autor}",
		Metodo:       http.MethodDelete,
		Autenticacao: false,
		Func:         controllers.DeletePublicationByTitleAndAutor,
	},
}
