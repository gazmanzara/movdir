package app

import (
	"encoding/json"
	"github.com/gazmanzara/movdir/dto"
	"github.com/gazmanzara/movdir/errs"
	"github.com/gazmanzara/movdir/service"
	"github.com/gorilla/mux"
	"net/http"
)

type DirectorHandlers struct {
	service service.DirectorService
}

func (ch *DirectorHandlers) getAllDirectors(w http.ResponseWriter, _ *http.Request) {
	directors, err := ch.service.GetAllDirectors()

	if err != nil {
		writeJsonResponse(err.Code, w, err.AsMessage())
		return
	} else {
		writeJsonResponse(http.StatusOK, w, directors)
		return
	}
}

func (ch *DirectorHandlers) getDirectorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	director, err := ch.service.GetDirectorById(id)

	if err != nil {
		writeJsonResponse(err.Code, w, err.AsMessage())
		return
	} else {
		writeJsonResponse(http.StatusOK, w, director)
		return
	}
}

func (ch *DirectorHandlers) createDirector(w http.ResponseWriter, r *http.Request) {
	var request dto.Director

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		newErr := errs.NewBadRequestError("Invalid request body")
		writeJsonResponse(newErr.Code, w, newErr.AsMessage())
		return
	} else {
		director, appError := ch.service.CreateDirector(request)
		if appError != nil {
			writeJsonResponse(appError.Code, w, appError.AsMessage())
			return
		} else {
			writeJsonResponse(http.StatusCreated, w, director)
			return
		}
	}
}

func writeJsonResponse(code int, w http.ResponseWriter, data interface{}) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
