package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		expectedURL string
	}{
		{
			name:        "remove scheme",
			inputURL:    "https://www.felo.com/path",
			expectedURL: "www.felo.com/path",
		},
		{
			name:        "remove trailing slash",
			inputURL:    "https://www.felo.com/path/",
			expectedURL: "www.felo.com/path",
		},
		{
			name:        "Remove trailing sash, preserve path",
			inputURL:    "https://www.felo.com/path/maspath/",
			expectedURL: "www.felo.com/path/maspath",
		},
		{
			name:        "no scheme",
			inputURL:    "www.felo.com/path",
			expectedURL: "www.felo.com/path",
		},
		{
			name:        "remove query parameters",
			inputURL:    "https://www.felo.com/path?param=value",
			expectedURL: "www.felo.com/path",
		},
		{
			name:        "only domain and scheme",
			inputURL:    "https://www.felo.com",
			expectedURL: "www.felo.com",
		},
		{
			name:        "preserve subdomain",
			inputURL:    "https://sub.felo.com/path",
			expectedURL: "sub.felo.com/path",
		},
		{
			name:        "remove fragment",
			inputURL:    "https://www.felo.com/path#fragment",
			expectedURL: "www.felo.com/path",
		},
		{
			name:        "remove different scheme",
			inputURL:    "http://www.felo.com/path",
			expectedURL: "www.felo.com/path",
		},
		{
			name:        "preserve port",
			inputURL:    "https://www.felo.com:8080/path",
			expectedURL: "www.felo.com:8080/path",
		},
	}

	for i, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			actual, err := normalizeURL(tst.inputURL)
			if err != nil {
				t.Errorf("Test #%v - %s failed unexpectedly.\nInput URL: %v\nError: %v\n----------\n", i, tst.name, tst.inputURL, err)
				return
			}
			if actual != tst.expectedURL {
				t.Errorf("Test #%v - %s failed.\nInput URL: %v\nExpected URL: %v\nActual URL: %v\n----------\n", i, tst.name, tst.inputURL, tst.expectedURL, actual)
			}
		})
	}
}
