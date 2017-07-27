package main

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"log"
)

type Response struct {
	Domain string    `json:"domain"`
	Blocked bool    `json:"blocked"`
}

func main()  {
	storage := NewStorage()
	err := storage.Open()
	defer storage.Close()

	if err != nil {
		log.Fatal(err)
	}

	synchronizer := NewSynchronizer(storage)
	go synchronizer.Start()

	controller := NewCheckController(storage)

	r := mux.NewRouter()
	r.HandleFunc("/v1/check/{email}", controller.Check).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	http.ListenAndServe(":8000", n)
}
