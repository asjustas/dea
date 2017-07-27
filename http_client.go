package main

import (
	"net/http"
	"bufio"
	"bytes"
)

func getURL(url string) (string, *http.Response, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", response, err
	}

	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}

	return buf.String(), response, nil
}

