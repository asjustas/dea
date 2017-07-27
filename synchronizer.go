package main

import (
	"log"
	"fmt"
)

type Synchronizer struct {
	Storage *Storage
}

func NewSynchronizer(storage *Storage) *Synchronizer {
	s := new(Synchronizer)
	s.Storage = storage

	return s
}

func (synchronizer *Synchronizer) Start() {
	provider := NewMattKetmoProvider()
	domains, _ := provider.Get()

	for _, domain := range domains {
		err := synchronizer.Storage.Add(domain)
		fmt.Println(domain)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}