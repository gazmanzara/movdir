package app

import (
	"encoding/json"
	"github.com/gazmanzara/movdir/auth/dto"
	"github.com/gazmanzara/movdir/auth/errs"
	"github.com/gazmanzara/movdir/auth/service"
	"net/http"
)

type Handlers struct {
	service service.AuthService
}

func (h *Handlers) login(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		newErr := errs.NewBadRequestError("invalid request payload")
		writeJsonResponse(newErr.Code, w, newErr.AsMessage())
		return
	}

	res, appErr := h.service.Login(payload)
	if appErr != nil {
		writeJsonResponse(appErr.Code, w, appErr.AsMessage())
		return
	}
	writeJsonResponse(http.StatusOK, w, res)
	return
}

func (h *Handlers) register(w http.ResponseWriter, r *http.Request) {
	var payload dto.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		newErr := errs.NewBadRequestError("invalid request payload")
		writeJsonResponse(newErr.Code, w, newErr.AsMessage())
		return
	}

	res, appErr := h.service.Register(payload)
	if appErr != nil {
		writeJsonResponse(appErr.Code, w, appErr.AsMessage())
		return
	}
	writeJsonResponse(http.StatusOK, w, res)
	return
}

func writeJsonResponse(code int, w http.ResponseWriter, data interface{}) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
