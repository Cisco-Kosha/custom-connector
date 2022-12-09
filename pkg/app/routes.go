package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV2 = "/api/v2"

	// specification routes
	a.Router.HandleFunc(apiV2+"/customconnector/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/customconnector/specification/test", a.testConnectorSpecification).Methods("POST", "OPTIONS")

	a.Router.HandleFunc(apiV2+"/customconnector/createservice", a.createRecords).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/customconnector/getservice", a.getRecords).Methods("GET", "OPTIONS")

	a.Router.HandleFunc(apiV2+"/customconnector/updateservice", a.updateRecords).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/customconnector/deleteservice", a.deleteRecords).Methods("DELETE", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
