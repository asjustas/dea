package main

import (
	"log"
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
	domains, _ := NewDomainsJsonProvider().Get()
	synchronizer.addDomains(domains)

	domains, _ = NewDomainsTxtProvider().Get()
	synchronizer.addDomains(domains)
}

func (synchronizer *Synchronizer) addDomains(domains []string) {
	for _, domain := range domains {
		err := synchronizer.Storage.Add(domain)

		if err != nil {
			log.Fatal(err)
		}
	}
}