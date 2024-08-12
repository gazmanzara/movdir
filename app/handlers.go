package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gazmanzara/movdir/service"
	"net/http"
)

type Director struct {
	Id     int    `json:"id" xml:"id"`
	Name   string `json:"name" xml:"name"`
	Gender int    `json:"gender" xml:"gender"`
}

type DirectorHandlers struct {
	service service.DirectorService
}

func (ch *DirectorHandlers) getAllDirectors(w http.ResponseWriter, r *http.Request) {
	directors, _ := ch.service.GetAllDirectors()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(directors)
		if err != nil {
			return
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(directors)
		if err != nil {
			return
		}
	}
}
