package app

import (
	"encoding/json"
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
	} else {
		writeJsonResponse(http.StatusOK, w, directors)
	}
}

func (ch *DirectorHandlers) getDirectorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	director, err := ch.service.GetDirectorById(id)

	if err != nil {
		writeJsonResponse(err.Code, w, err.AsMessage())
	} else {
		writeJsonResponse(http.StatusOK, w, director)
	}
}

func writeJsonResponse(code int, w http.ResponseWriter, data interface{}) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
