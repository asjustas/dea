package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type CheckController struct {
	Storage *Storage
}

func NewCheckController(storage *Storage) *CheckController {
	controller := new(CheckController)
	controller.Storage = storage

	return controller
}

func (controller *CheckController) Check (w http.ResponseWriter, r *http.Request) {
	domain := controller.getDomain(mux.Vars(r)["email"])
	blocked, err := controller.Storage.Exists(domain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	res, err := json.Marshal(
		Response{
			Domain: domain,
			Blocked: blocked,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (controller *CheckController) getDomain(email string) string {
	return email
}
