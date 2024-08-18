package app

import (
	"github.com/gazmanzara/movdir/auth/domain"
	"github.com/gazmanzara/movdir/auth/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	h := Handlers{
		service: service.NewAuthService(domain.NewAuthRepositoryStub()),
	}

	router.HandleFunc("/login", h.login).Methods("POST")
	router.HandleFunc("/register", h.register).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
