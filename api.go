package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

type Response struct {
	Domain string    `json:"domain"`
	Blocked bool    `json:"blocked"`
}

func main()  {
	controller := NewCheckController()

	r := mux.NewRouter()
	r.HandleFunc("/v1/check/{email}", controller.Check).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	http.ListenAndServe(":8000", n)
}
