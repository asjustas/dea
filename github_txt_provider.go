package main

import (
	"strings"
	"github.com/pkg/errors"
)

type GithubTxtProvider struct {
}

func NewGithubTxtProvider() *GithubTxtProvider {
	return new(GithubTxtProvider)
}

func (provider *GithubTxtProvider) Get() ([]string, error) {
	domains := []string{}

	urls := []string{
		"https://raw.githubusercontent.com/MattKetmo/EmailChecker/master/res/throwaway_domains.txt",
		"https://raw.githubusercontent.com/wesbos/burner-email-providers/master/emails.txt",
		"https://raw.githubusercontent.com/andreis/disposable/6cceca83172b76c69337b52d9777bfaaf5aae580/domains.txt",
	}

	for _, url := range urls {
		providerDomains, _ := provider.getSingle(url)

		for _, domain := range providerDomains {
			domains = append(domains, domain)
		}
	}

	return domains, nil
}

func (provider *GithubTxtProvider) getSingle(url string) ([]string, error) {
	content, response, err := getURL(url)

	if err != nil {
		return []string{}, err
	}

	if response.StatusCode != 200 {
		return []string{}, errors.New(content)
	}

	return strings.Split(content,"\n"), nil
}