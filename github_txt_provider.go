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
		"https://raw.githubusercontent.com/andreis/disposable/master/domains.txt",
		"https://raw.githubusercontent.com/mailster/mailster-email-verify/master/dea.txt",
		"https://raw.githubusercontent.com/vboctor/disposable_email_checker/master/data/domains.txt",
		"https://raw.githubusercontent.com/khanhicetea/fast-email-validator/master/src/database/DisposableEmails.txt",
		"https://raw.githubusercontent.com/nojacko/email-data-disposable/master/bin/disposable.txt",
		"https://raw.githubusercontent.com/slester/disposable-email-providers/master/domains.txt",
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