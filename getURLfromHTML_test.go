package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
	<html>
		<body>
			<a href="/path/one">
				<span>Boot.dev</span>
			</a>
			<a href="https://other.com/path/one">
				<span>Boot.dev</span>
			</a>
		</body>
	</html>
	`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}
	for i, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tst.inputBody, tst.inputURL)
			if err != nil {
				t.Errorf("Test #%v - %s failed unexpectedly.\nError: %v\n----------\n", i, tst.name, err)
				return
			}
			if reflect.DeepEqual(actual, tst.expected) != true {
				t.Errorf("Test #%v - %s failed.\nInput Body: %v\nInput baseURL: %v\nexpected URL: %v\nactualURL: %v\n----------\n", i, tst.name, tst.inputBody, tst.inputBody, tst.expected, actual)
			}
		})
	}
}
