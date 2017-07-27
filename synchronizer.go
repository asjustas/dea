package main

import (
	"log"
	"fmt"
)

type Synchronizer struct {
	Storage *Storage
}

func NewSynchronizer(storage *Storage) *Synchronizer {
	synchronizer := new(Synchronizer)
	synchronizer.Storage = storage

	return synchronizer
}

func (synchronizer *Synchronizer) Start() {
	provider := NewMattKetmoProvider()
	domains, _ := provider.Get()
	synchronizer.addDomains(domains)

	fmt.Println("done")
}

func (synchronizer *Synchronizer) addDomains(domains []string) {
	for _, domain := range domains {
		err := synchronizer.Storage.Add(domain)
		fmt.Println(domain)

		if err != nil {
			log.Fatal(err)
		}
	}
}