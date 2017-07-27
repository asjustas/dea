package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type CheckController struct {
}

func NewCheckController() *CheckController {
	return new(CheckController)
}

func (controller *CheckController) Check (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res, err := json.Marshal(
		Response{
			Domain: vars["email"],
			Blocked: false,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
