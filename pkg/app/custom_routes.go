package app

import (
	"io/ioutil"
	"net/http"

	"github.com/kosha/custom-connector/pkg/httpclient"
)

// getService= godoc
// @Summary Get 3rd party api url
// @Description Get service
// @Tags Records
// @Accept  json
// @Produce  json
// @Success 200
// @Router /api/v2/customconnector/getservice [get]
func (a *App) getRecords(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}
	res, err := httpclient.GetRecord("https://www.google.com/", bodyBytes, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in GetRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}

// createService= godoc
// @Summary Create a record for a project
// @Description Create a record for airtable
// @Tags createRecords
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v2/customconnector/createservice [get]
func (a *App) createRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}

	res, err := httpclient.CreateRecord("https://www.google.com/", bodyBytes, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in createRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}

// deleteService= godoc
// @Summary delete a record for a project
// @Description delete a record for the airtable
// @Tags deleteService
// @Produce json
// @Success 200
// @Router /api/v2/customconnector/deleteservice [delete]
func (a *App) deleteRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}
	res, err := httpclient.DeleteRecord("https://www.google.com/", bodyBytes, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in deleteRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}

// updateService= godoc
// @Summary update a record for a project
// @Description Update a record for the airtable
// @Tags updateService
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v2/customconnector/updateservice [put]
func (a *App) updateRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}

	res, err := httpclient.UpdateRecord("https://www.google.com/", bodyBytes, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in updateRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}
