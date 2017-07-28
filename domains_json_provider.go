package main

import (
	"github.com/pkg/errors"
	"encoding/json"
)

type DomainsJsonProvider struct {
}

func NewDomainsJsonProvider() *DomainsJsonProvider {
	return new(DomainsJsonProvider)
}

func (provider *DomainsJsonProvider) Get() ([]string, error) {
	domains := []string{}

	urls := []string{
		"https://raw.githubusercontent.com/ClippingGorilla/ClippingGorilla/master/ClippingGorilla/src/main/resources/mailblacklist/blacklist.json",
		"https://raw.githubusercontent.com/yclas/yclas/master/oc/banned_domains.json",
		"https://raw.githubusercontent.com/PHPAuth/PHPAuth/master/files/domains.json",
		"https://raw.githubusercontent.com/ivolo/disposable-email-domains/master/index.json",
		"https://raw.githubusercontent.com/fnando/validators/master/data/disposable.json",
		"https://raw.githubusercontent.com/nojacko/email-data-disposable/master/data/disposable.json",
	}

	for _, url := range urls {
		providerDomains, _ := provider.getSingle(url)

		for _, domain := range providerDomains {
			domains = append(domains, domain)
		}
	}

	return domains, nil
}

func (provider *DomainsJsonProvider) getSingle(url string) ([]string, error) {
	var domains []string

	content, response, err := getURL(url)

	if err != nil {
		return []string{}, err
	}

	if response.StatusCode != 200 {
		return []string{}, errors.New(content)
	}

	err = json.Unmarshal([]byte(content), &domains)

	if err != nil {
		return []string{}, err
	}

	return domains, nil
}