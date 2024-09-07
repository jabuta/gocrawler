package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return "", fmt.Errorf("url returned status code error: %v - %s", res.StatusCode, res.Status)
	}
	if res.Header.Get("Content-type") != "text/html" {
		return "", fmt.Errorf("url returned content type: %s", res.Header.Get("Content-type"))
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
