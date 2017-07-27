package main

import (
	"strings"
	"github.com/pkg/errors"
)

type MattKetmoProvider struct {
}

func NewMattKetmoProvider() *MattKetmoProvider {
	return new(MattKetmoProvider)
}

func (provider *MattKetmoProvider) Get() ([]string, error) {
	content, response, err := getURL("https://raw.githubusercontent.com/MattKetmo/EmailChecker/master/res/throwaway_domains.txt")

	if err != nil {
		return []string{}, err
	}

	if response.StatusCode != 200 {
		return []string{}, errors.New(content)
	}

	return strings.Split(content,"\n"), nil
}