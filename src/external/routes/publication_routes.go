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
}
