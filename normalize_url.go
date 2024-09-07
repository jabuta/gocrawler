package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", strings.TrimRight(parsed.Host, "/"), strings.TrimRight(parsed.Path, "/")), nil
}
