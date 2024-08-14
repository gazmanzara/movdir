package app

import (
	"github.com/gazmanzara/movdir/domain"
	"github.com/gazmanzara/movdir/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	ch := DirectorHandlers{service: service.NewDirectorService(domain.NewDirectorRepositoryDB())}

	router.HandleFunc("/directors", ch.getAllDirectors).Methods(http.MethodGet)
	router.HandleFunc("/directors/{id}", ch.getDirectorById).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
