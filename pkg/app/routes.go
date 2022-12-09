package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV2 = "/api/v2"

	// specification routes
	a.Router.HandleFunc(apiV2+"/customconnector/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/customconnector/specification/test", a.testConnectorSpecification).Methods("POST", "OPTIONS")

	a.Router.HandleFunc(apiV2+"/customconnector/createservice", a.listConnectorSpecification).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/customconnector/getservice", a.testConnectorSpecification).Methods("GET", "OPTIONS")

	a.Router.HandleFunc(apiV2+"/customconnector/updateservice", a.listConnectorSpecification).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/customconnector/deleteservice", a.testConnectorSpecification).Methods("DELETE", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
